package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model untuk resep masakan
type Recipe struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title" validate:"required"`
    Ingredients []string           `json:"ingredients" bson:"ingredients" validate:"required"`
    Steps       []string           `json:"steps" bson:"steps" validate:"required"`
    Duration    int                `json:"duration" bson:"duration" validate:"required"` // dalam menit
    Difficulty string             `json:"difficulty" bson:"difficulty" validate:"required,oneof=easy medium hard"`
}