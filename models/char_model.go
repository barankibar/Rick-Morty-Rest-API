package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name" validate:"required"`
	Status   string             `json:"status"`
	Species  string             `json:"species"`
	Gender   string             `json:"gender"`
	Origin   map[string]string  `json:"origin"`
	Location map[string]string  `json:"location"`
	Image    string             `json:"image"`
	Episode  []string           `json:"episode"`
	URL      string             `json:"url"`
	Created  string             `json:"created"`
}
