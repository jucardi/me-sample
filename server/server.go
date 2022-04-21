package server

import (
	"github.com/jucardi/go-beans/beans"
	"github.com/jucardi/go-titan/utils/shutdown"
	"{{.golang.module_path}}/{{.service_name}}/server/api"
	"{{.golang.module_path}}/{{.service_name}}/server/repository"
)

func Run() {
	// Setup Shutdown Listener
	shutdown.ListenForSignals()

	// Initializes the database connection
	repository.InitRepositories()

	// Initializes all registered singleton beans
	beans.InitComponents()

	// Start listening for requests
	api.Start()
}
