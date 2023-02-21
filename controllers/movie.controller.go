package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/go-mongo-rest-app/configs"
	"github.com/devfurkankizmaz/go-mongo-rest-app/models"
	"github.com/devfurkankizmaz/go-mongo-rest-app/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movies")
var validate = validator.New()

type EntryRecord struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	MaxIMDB   float64 `json:"maxIMDB"`
	MinIMDB   float64 `json:"minIMDB"`
}

func GetFilteredData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var body EntryRecord
		defer cancel()

		if err := c.Bind(&body); err != nil { //validate request body
			c.JSON(http.StatusBadRequest,
				responses.MovieResponse{
					Code:    http.StatusBadRequest,
					Message: "error",
					Records: map[string]interface{}{"data": err},
				})
			return
		}

		if validationErr := validate.Struct(&body); validationErr != nil { //validate required fields by using validator lib
			c.JSON(http.StatusBadRequest, responses.MovieResponse{
				Code:    http.StatusBadRequest,
				Message: "error",
				Records: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		sd, _ := time.Parse("2006-01-02", body.StartDate)
		ed, _ := time.Parse("2006-01-02", body.EndDate)

		fmt.Println(body)

		filterCursor, err := movieCollection.Find(ctx,
			bson.M{
				"release": bson.M{"$gt": sd, "$lt": ed},
				"imdb":    bson.M{"$gt": body.MinIMDB, "$lt": body.MaxIMDB},
			})

		var result []models.Movie

		if err = filterCursor.All(ctx, &result); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.MovieResponse{
					Code:    http.StatusBadRequest,
					Message: "error",
					Records: map[string]interface{}{"data": err},
				})
			return
		}

		c.JSON(http.StatusOK, responses.MovieResponse{
			Code:    http.StatusOK,
			Message: "success",
			Records: map[string]interface{}{"data": result}},
		)
	}
}
