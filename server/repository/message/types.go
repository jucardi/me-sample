package message

import "context"

// IRepository defines the contract for a clients repository
type IRepository interface {
	// WithCtx passes a context to this instance for all subsequent operations to be executed with said context
	WithCtx(ctx context.Context) IRepository

	// Create creates a new message entry record in the database.
	Create(info *MessageDbe) error

	// Update updates the values of an existing `Message` record. Fails if the record does not exist.
	Update(info *MessageDbe) error

	// Delete deletes a `Message` record form the database.
	Delete(name string) error

	// First finds a `Message` record with the given name. Returns `record not found` error if not found.
	First(name string) (*MessageDbe, error)
}
