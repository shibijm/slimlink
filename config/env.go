package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var BindAddress string
var BindPort string
var RedisConnectionString string
var LinkIDLength int

func Load() {
	godotenv.Load()
	BindAddress = os.Getenv("BIND_ADDRESS")
	if BindAddress == "" {
		BindAddress = "127.0.0.1"
	}
	BindPort = os.Getenv("BIND_PORT")
	if BindPort == "" {
		BindPort = "44558"
	}
	RedisConnectionString = os.Getenv("REDIS_CONNECTION_STRING")
	var err error
	LinkIDLength, err = strconv.Atoi(os.Getenv("LINK_ID_LENGTH"))
	if LinkIDLength < 1 || err != nil {
		LinkIDLength = 5
	}
}
