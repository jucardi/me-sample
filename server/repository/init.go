package repository

import (
	"github.com/jucardi/go-titan/components/mongo"
	"github.com/jucardi/go-titan/logx"
)

// InitRepositories opens any required connections to be used by the repositories in the service.
func InitRepositories() {
	_, err := mongo.Dial()
	logx.WithObj(err).Fatal("failed to connect to mongo instance")
}
