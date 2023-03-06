package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"slimlink/config"
	"slimlink/core/ports"
	"slimlink/core/services"
	"slimlink/infrastructure/data"
	"slimlink/infrastructure/logging"
	"slimlink/infrastructure/repos"
	"slimlink/interface/controllers"
	"slimlink/interface/routers"

	"github.com/redis/go-redis/v9"
)

//go:embed all:web/out
var embeddedWebUIFileSystem embed.FS
var consoleLogger ports.Logger

func main() {
	config.Load()
	consoleLogger = logging.NewConsoleLogger()
	address := fmt.Sprintf("%s:%s", config.BindAddress, config.BindPort)
	consoleLogger.Log("Slimlink v%s", config.Version)
	consoleLogger.Log("Bind Address: http://%s", address)
	options, err := redis.ParseURL(config.RedisConnectionString)
	if err != nil {
		exitWithError(err, "failed to parse Redis connection string")
	}
	redisDB, err := data.NewRedisDB(options)
	if err != nil {
		exitWithError(err, "failed to initialise Redis")
	}
	webUIFileSystem, err := fs.Sub(embeddedWebUIFileSystem, "web/out")
	if err != nil {
		exitWithError(err, "failed to read embedded web UI filesystem")
	}
	linkRedisRepo := repos.NewLinkRedisRepo(redisDB)
	httpFileSystem := data.NewHttpFileSystem(http.FS(webUIFileSystem))
	linkService := services.NewLinkService(linkRedisRepo, config.LinkIDLength)
	webUIController := controllers.NewWebUIController(consoleLogger, httpFileSystem)
	linkController := controllers.NewLinkController(consoleLogger, linkService)
	webUIRouter := routers.NewWebUIRouter(webUIController)
	apiRouter := routers.NewApiRouter(linkController)
	rootRouter := routers.NewRootRouter(webUIRouter, apiRouter)
	http.HandleFunc("/", rootRouter.Route)
	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			exitWithError(err, "failed to run HTTP server")
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	consoleLogger.Log("Received interrupt signal")
	consoleLogger.Log("Exiting")
}

func exitWithError(err error, message string) {
	consoleLogger.LogError(fmt.Errorf("%s: %w", message, err), "main")
	os.Exit(1)
}
