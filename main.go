package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"slimlink/services/api"
	"slimlink/services/data"
	"slimlink/services/routers"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

//go:embed all:web/out
var embeddedFS embed.FS

func main() {
	godotenv.Load()
	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {
		bindAddress = "127.0.0.1"
	}
	bindPort := os.Getenv("BIND_PORT")
	if bindPort == "" {
		bindPort = "44558"
	}
	address := fmt.Sprintf("%s:%s", bindAddress, bindPort)
	fmt.Printf("Slimlink\nURL: http://%s\n", address)
	options, err := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data.InitRedisDb(options)
	uiOutput, err := fs.Sub(embeddedFS, "web/out")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	api.InitContentApi(http.FS(uiOutput))
	length, err := strconv.Atoi(os.Getenv("LINK_ID_LENGTH"))
	if err != nil {
		length = 5
	}
	api.InitLinkApi(length)
	http.HandleFunc("/", routers.RootRouter)
	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("Received interrupt signal")
	fmt.Println("Exiting")
}
