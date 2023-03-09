package main

import (
	"embed"
	"errors"
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
	dbConnectionStrings := []string{config.RedisConnectionString, config.MySqlConnectionString}
	totalCount := len(dbConnectionStrings)
	emptyCount := countEmpty(dbConnectionStrings)
	if emptyCount == totalCount {
		exitWithError(nil, "no database connection string is set")
	}
	if emptyCount < totalCount-1 {
		exitWithError(nil, "multiple database connection strings are set")
	}
	var linkRepo ports.LinkRepo
	if config.RedisConnectionString != "" {
		options, err := redis.ParseURL(config.RedisConnectionString)
		if err != nil {
			exitWithError(err, "failed to parse Redis connection string")
		}
		db, err := data.NewRedisDB(options)
		if err != nil {
			exitWithError(err, "failed to initialise Redis connection")
		}
		consoleLogger.Log("Connected to Redis")
		linkRepo = repos.NewLinkRedisRepo(db)
	} else {
		db, err := data.NewMySqlDB(config.MySqlConnectionString)
		if err != nil {
			exitWithError(err, "failed to initialise MySQL connection")
		}
		err = db.EnsureCreated()
		if err != nil {
			exitWithError(err, "failed to initialise MySQL database")
		}
		consoleLogger.Log("Connected to MySQL")
		linkRepo = repos.NewLinkMySqlRepo(db)
	}
	if config.LinkIDLength == 0 {
		exitWithError(nil, "invalid link ID length")
	}
	webUIFileSystem, err := fs.Sub(embeddedWebUIFileSystem, "web/out")
	if err != nil {
		exitWithError(err, "failed to read embedded web UI filesystem")
	}
	httpFileSystem := data.NewHttpFileSystem(http.FS(webUIFileSystem))
	linkService := services.NewLinkService(linkRepo, config.LinkIDLength)
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
	var errToLog error
	if err == nil {
		errToLog = errors.New(message)
	} else {
		errToLog = fmt.Errorf("%s: %w", message, err)
	}
	consoleLogger.LogError(errToLog, "main")
	os.Exit(1)
}

func countEmpty(strings []string) int {
	emptyCount := 0
	for _, str := range strings {
		if str == "" {
			emptyCount++
		}
	}
	return emptyCount
}
