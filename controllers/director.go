package controllers

import (
	"context"
	"time"

	"github.com/devfurkankizmaz/go-mongo-rest-app/configs"
	"github.com/devfurkankizmaz/go-mongo-rest-app/models"
	"github.com/devfurkankizmaz/go-mongo-rest-app/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var directorCollection *mongo.Collection = configs.GetCollection(configs.DB, "directors")

func CreateDirector() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var director models.Director
		var movie []models.Movie

		defer cancel()

		//validate the required fields
		if err := c.BindJSON(&director); err != nil {
			responses.ErrorRes(c, err)
			return
		}

		results, err := movieCollection.Find(ctx, bson.M{"director": director.Name})

		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleMovie models.Movie
			if err = results.Decode(&singleMovie); err != nil {
				responses.ErrorRes(c, err)
				return
			}

			movie = append(movie, singleMovie)

		}

		if err != nil {
			responses.ErrorRes(c, err)
			return
		}

		m := make([]string, len(movie))

		for i := range movie {
			m[i] = movie[i].Title
		}

		newDirector := models.Director{
			Id:     primitive.NewObjectID(),
			Name:   director.Name,
			DOB:    director.DOB,
			Movies: map[string]interface{}{"data": m},
		}

		if result, err := directorCollection.InsertOne(ctx, newDirector); err != nil {
			responses.ErrorRes(c, err)
			return
		} else {
			responses.StatusCreated(c, result)
		}
	}
}
func GetAllDirectors() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var director []models.Director
		defer cancel()

		results, err := directorCollection.Find(ctx, bson.M{})

		if err != nil {
			responses.ErrorRes(c, err)
			return
		}

		//reading from the db
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleDirector models.Director
			if err = results.Decode(&singleDirector); err != nil {
				responses.ErrorRes(c, err)
				return
			}

			director = append(director, singleDirector)
		}

		responses.StatusCreated(c, director)
	}
}
