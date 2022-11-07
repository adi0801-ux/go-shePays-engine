package main

import (
	"log"
	"shepays/api"
	"shepays/db"
	"shepays/models"
	"shepays/repositories"
	"shepays/services"
	"shepays/utils"
)

func main() {
	err := utils.InitialiseLogger()
	if err != nil {
		log.Fatalln(err)
	}
	utils.Log.Info("logger initialized")

	//loading config from env file
	//**live loading for config added **
	utils.Log.Info("config loading...")
	config, err := utils.LoadConfig(".")
	if err != nil {
		utils.Log.Fatal(err)
		return
	}

	utils.Log.Info("config loaded")

	utils.Log.Info("database connecting....")

	//Make DB Connection
	store, err := makeDBConnection(config)
	if err != nil {
		utils.Log.Fatal(err)
		return
	}
	// Make Migrations
	err = store.RunMigrations()

	if err != nil {
		utils.Log.Fatal("error creating migrations")
	}

	utils.Log.Info("database connected")

	nsdlRep := repositories.NSDLClient{Client: repositories.CreateHttpClient(),
		LogRep:    &repositories.ApiLogsRepository{Db: store},
		BaseUrl:   config.NSDLUrl,
		ChannelId: config.ChannelId,
		AppDtls: &models.AppDtls{
			Appid:       config.AppId,
			ApplVersion: config.ApplVersion,
			AppRegFlg:   config.AppRegFlg,
		},
	}

	//create service reference
	proxySrv := services.CreateAllRepositoryReferences(store)
	proxySrv.NSDLClient = &nsdlRep

	//create cron Reference

	//creating a config

	utils.Log.Info("api server initializing")
	//Create HTTP Server
	server := api.GetNewServer(config, &proxySrv)

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		utils.Log.Fatal("cannot start server: ", err)
	}

	utils.Log.Info("api server initialized")

	// Results:
	// Name: addUser, Method: GET
	// Name: destroyUser, Method: DELETE

}

func makeDBConnection(config *utils.Config) (*db.Database, error) {
	dbConfig := &db.ConnectionConfig{
		DSN: config.DSN,
	}

	database, err := db.ConnectToDB(dbConfig)
	return database, err
}
