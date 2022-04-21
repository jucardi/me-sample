package config

// ServiceConfig contains all configuration parameters for the service
type ServiceConfig struct {
	SomeConfigField string `json:"some_service_specificy_config_field" yaml:"some_service_specificy_config_field"`
}
