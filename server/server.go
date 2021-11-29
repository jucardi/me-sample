package server

import (
	"github.com/jucardi/go-beans/beans"
	"github.com/jucardi/go-titan/utils/shutdown"
	"github.com/jucardi/ms-sample/server/api"
	"github.com/jucardi/ms-sample/server/repository"
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
