package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty" binding:"required"`
	Director string             `bson:"director,omitempty" binding:"required"`
	IMDB     float64            `bson:"imdb,omitempty" binding:"required"`
	Release  time.Time          `bson:"release,omitempty" binding:"required"`
	Counts   []int              `bson:"counts,omitempty" binding:"required"`
}
