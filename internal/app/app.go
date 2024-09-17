package app

import (
	"Building/config"
	"Building/internal/controller"
)

func App() {
	config.ReadConfig()
	db := config.InitializeStorePostgres()
	defer db.PostgresDB.Close()

	controller.Registry()
}
