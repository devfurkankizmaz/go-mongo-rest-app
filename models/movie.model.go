package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty" validate:"required"`
	Director string             `bson:"director,omitempty" validate:"required"`
	IMDB     float64            `bson:"imdb,omitempty" validate:"required"`
	Release  time.Time          `bson:"release,omitempty" validate:"required"`
	Counts   []int              `bson:"counts,omitempty" validate:"required"`
}
