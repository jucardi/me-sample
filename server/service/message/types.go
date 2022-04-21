package message

import (
	"context"
	"{{.golang.module_path}}/{{.service_name}}/api/helloworld"
)

// IService defines the contract for a clients service
type IService interface {
	// WithCtx passes a context to this instance for all subsequent operations to be executed with said context
	WithCtx(ctx context.Context) IService

	// Create creates a new message entry
	Create(name, message string) error

	// Get attempts to retrieve an existing message entry
	Get(name string) (*helloworld.Message, error)

	// Update allows to update a message entry by the given name
	Update(name, message string) error

	// Exists indicates if the specified message entry exists.
	Exists(name string) bool
}
