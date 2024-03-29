package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var BindAddress string
var BindPort string
var LinkIDLength int
var RedisConnectionString string
var MySqlConnectionString string
var PageFooterText string

func LoadFromEnv() {
	godotenv.Load()
	BindAddress = os.Getenv("BIND_ADDRESS")
	if BindAddress == "" {
		BindAddress = "0.0.0.0"
	}
	BindPort = os.Getenv("BIND_PORT")
	if BindPort == "" {
		BindPort = "44558"
	}
	linkIDLength := os.Getenv("LINK_ID_LENGTH")
	if linkIDLength == "" {
		LinkIDLength = 5
	} else {
		var err error
		LinkIDLength, err = strconv.Atoi(linkIDLength)
		if err != nil || LinkIDLength < 1 || LinkIDLength > 64 {
			LinkIDLength = 0
		}
	}
	RedisConnectionString = os.Getenv("REDIS_CONNECTION_STRING")
	MySqlConnectionString = os.Getenv("MYSQL_CONNECTION_STRING")
	PageFooterText = os.Getenv("PAGE_FOOTER_TEXT")
}
