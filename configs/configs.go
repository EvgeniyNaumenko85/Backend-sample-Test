package configs

import "BST/models"

var AppSettings models.Settings

func PutAdditionalSettings() {
	AppSettings.AppParams.LogDebug = "./logs/debug.log"
	AppSettings.AppParams.LogInfo = "./logs/info.log"
	AppSettings.AppParams.LogWarning = "./logs/warning.log"
	AppSettings.AppParams.LogError = "./logs/error.log"

	AppSettings.AppParams.LogCompress = true
	AppSettings.AppParams.LogMaxSize = 10
	AppSettings.AppParams.LogMaxAge = 100
	AppSettings.AppParams.LogMaxBackups = 100
	AppSettings.AppParams.AppVersion = "1.0"
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
