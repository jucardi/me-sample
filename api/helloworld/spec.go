package helloworld

import "context"

const (
	ApiVersion         = 1
	ApiRoute           = "/hello"
	RouteMessages      = "/message"
	RouteMessagesParam = "/message/:name"
)

// IApi defines the api contract for `helloworld`
type IApi interface {
	// WithCtx passes a context to this instance for all subsequent operations to be executed with said context
	WithCtx(ctx context.Context) IApi

	// GetHelloWorld returns HelloWorld text
	GetHelloWorld() (*Message, error)

	// Create creates a new message entry
	Create(req *Message) error

	// Get attempts to retrieve an existing message entry
	Get(name string) (*Message, error)

	// Update allows to update a message entry by the given name
	Update(req *Message) error
}
