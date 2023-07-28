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
var logger ports.Logger

func main() {
	config.LoadFromEnv()
	logger = logging.NewConsoleLogger()
	logger.Log("Slimlink v%s", config.Version)
	address := fmt.Sprintf("%s:%s", config.BindAddress, config.BindPort)
	logger.Log("Bind Address: http://%s", address)
	if config.LinkIDLength == 0 {
		exitWithError(nil, "invalid link ID length")
	}
	dbConnectionStrings := []string{config.RedisConnectionString, config.MySqlConnectionString}
	emptyCount := countEmpty(dbConnectionStrings)
	totalCount := len(dbConnectionStrings)
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
		logger.Log("Connected to Redis")
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
		logger.Log("Connected to MySQL")
		linkRepo = repos.NewLinkMySqlRepo(db)
	}
	webUIFileSystem, err := fs.Sub(embeddedWebUIFileSystem, "web/out")
	if err != nil {
		exitWithError(err, "failed to read embedded web UI filesystem")
	}
	infoService := services.NewInfoService(config.PageFooterText)
	linkService := services.NewLinkService(linkRepo, config.LinkIDLength)
	webUIController := controllers.NewWebUIController(logger, http.FS(webUIFileSystem))
	infoController := controllers.NewInfoController(infoService)
	linkController := controllers.NewLinkController(logger, linkService)
	webUIRouter := routers.NewWebUIRouter(webUIController)
	apiRouter := routers.NewApiRouter(infoController, linkController)
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
	logger.Log("Received interrupt signal")
	logger.Log("Exiting")
}

func exitWithError(err error, message string) {
	var errToLog error
	if err == nil {
		errToLog = errors.New(message)
	} else {
		errToLog = fmt.Errorf("%s: %w", message, err)
	}
	logger.LogError(errToLog, "main")
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
