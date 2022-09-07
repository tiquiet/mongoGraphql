package custom_models

import (
	"time"
)

type Book struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Author      string    `json:"author" bson:"author"`
	LastUpdate  time.Time `json:"lastUpdate" bson:"lastUpdate"`
}
type NewBook struct {
	Title       string    `json:"title" bson:"title"`
	Description *string   `json:"description" bson:"description"`
	Author      string    `json:"author" bson:"author"`
	LastUpdate  time.Time `json:"lastUpdate" bson:"lastUpdate"`
}

type UpdateBook struct {
	Title       *string   `json:"title,omitempty" bson:"title,omitempty"`
	Description *string   `json:"description,omitempty" bson:"description,omitempty"`
	Author      *string   `json:"author,omitempty" bson:"author,omitempty"`
	LastUpdate  time.Time `json:"lastUpdate,omitempty" bson:"lastUpdate,omitempty"`
}
