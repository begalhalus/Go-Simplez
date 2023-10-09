package apps_config

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var PORT = "5005"
var DB_DRIVER = "pgsql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "5432"
var DB_NAME = ""
var DB_USER = ""
var DB_PASSWORD = ""

var STATIC_ROUTE = "/public"
var STATIC_DIR = "/utils/public/files"

var defaultLogFilePath = "./utils/public/files/logs/error.log"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")

	if portEnv != "" {
		PORT = portEnv
	}

	staticRouteEnv := os.Getenv("STATIC_ROUTE")

	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	staticDirEnv := os.Getenv("STATIC_DIR")

	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}
}

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")

	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")

	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")

	if portEnv != "" {
		DB_PORT = portEnv
	}

	nameEnv := os.Getenv("DB_NAME")

	if nameEnv != "" {
		DB_NAME = nameEnv
	}

	userEnv := os.Getenv("DB_USER")

	if userEnv != "" {
		DB_USER = userEnv
	}

	passwordEnv := os.Getenv("DB_PASSWORD")

	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}

func InitCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	return cors.New(config)
}

func createLogFolderNotExist(path string) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Create", dir, "directory")
		err := os.MkdirAll(dir, 0644)

		if err != nil {
			log.Println("Fail directory logs")
		} else {
			log.Println(dir, "directory created")
		}
	}
}

func OpenOrCreateLogFile(path string) (*os.File, error) {

	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		var errCreateFile error
		logFile, errCreateFile = os.Create(path)

		if errCreateFile != nil {
			log.Println("Can't create log file", errCreateFile)
		}

	}

	return logFile, nil

}

func DefaultLogging() {

	gin.DisableConsoleColor()

	createLogFolderNotExist(defaultLogFilePath)

	f, _ := OpenOrCreateLogFile(defaultLogFilePath)

	gin.DefaultWriter = io.MultiWriter(f)

	log.SetOutput(gin.DefaultWriter)

}
