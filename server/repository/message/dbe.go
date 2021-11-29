package message

type MessageDbe struct {
	// ID is the record id in the database
	ID string `bson:"_id,omitempty"`
	// Name is the message name
	Name string `bson:"name"`
	// Message is the message string
	Message string `bson:"message"`
}
