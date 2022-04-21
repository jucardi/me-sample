package config

import (
	"github.com/jucardi/go-titan/configx"
	"github.com/jucardi/go-titan/logx"
)

const (
	configKey  = "service"
	configName = "{{.service_name}}-cfg"
)

var (
	singleton = &ServiceConfig{}
)

func init() {
	configx.AddOnReloadCallback(func(cfg configx.IConfig) {
		config := &ServiceConfig{}

		logx.WithObj(
			cfg.MapToObj(configKey, config),
		).Fatal("unable to map service configuration")

		singleton = config
	}, configName)
}

// Service returns rest configuration
func Service() *ServiceConfig {
	return singleton
}

// Base returns the base configuration from configx
func Base() configx.IConfig {
	return configx.Get()
}
