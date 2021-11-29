package config

// ServiceConfig contains all configuration parameters for the service
type ServiceConfig struct {
	MinPasswordLength int `json:"min_password_length" yaml:"min_password_length"`
}
