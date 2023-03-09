package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Director struct {
	Id     primitive.ObjectID     `bson:"_id,omitempty"`
	Name   string                 `bson:"name,omitempty" binding:"required"`
	DOB    time.Time              `bson:"bod,omitempty" binding:"required"`
	Movies map[string]interface{} `bson:"movies,omitempty"`
}
