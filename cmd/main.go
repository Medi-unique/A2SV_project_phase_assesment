package main

import (
	"assessment1/config"
	"assessment1/config/db"
	"assessment1/delivery/routers"
)

func main() {
	config.InitiEnvConfigs()
	db.ConnectDB(config.EnvConfigs.MongoURI)

	router := routers.SetupRouter()

	router.Run(config.EnvConfigs.LocalServerPort)
}
