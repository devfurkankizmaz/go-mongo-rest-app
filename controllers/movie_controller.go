package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/go-mongo-rest-app/configs"
	"github.com/devfurkankizmaz/go-mongo-rest-app/models"
	"github.com/devfurkankizmaz/go-mongo-rest-app/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movies")

type EntryRecord struct {
	StartDate string  `json:"startDate" binding:"required"`
	EndDate   string  `json:"endDate" binding:"required"`
	MaxIMDB   float64 `json:"maxIMDB" binding:"required"`
	MinIMDB   float64 `json:"minIMDB" binding:"required"`
}

func ErrorRes(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, responses.MovieResponse{
		Code:    http.StatusInternalServerError,
		Message: "error",
		Records: map[string]interface{}{"data": err.Error()}})
}

func StatusCreated(c *gin.Context, result interface{}) {
	c.JSON(http.StatusCreated, responses.MovieResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Records: map[string]interface{}{"data": result}})
}

func GetFilmByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		movieId := c.Param("movieId")
		var movie models.Movie
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(movieId)

		err := movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&movie)
		if err != nil {
			ErrorRes(c, err)
			return
		}

		StatusCreated(c, movie)
	}
}

func CreateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie models.Movie
		defer cancel()

		//validate the required fields
		if err := c.ShouldBindJSON(&movie); err != nil {
			ErrorRes(c, err)
			return
		}

		newMovie := models.Movie{
			Id:       primitive.NewObjectID(),
			Title:    movie.Title,
			Director: movie.Director,
			IMDB:     movie.IMDB,
			Release:  movie.Release,
			Counts:   movie.Counts,
		}

		if result, err := movieCollection.InsertOne(ctx, newMovie); err != nil {
			ErrorRes(c, err)
			return
		} else {
			StatusCreated(c, result)
		}
	}
}

func GetAllMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie []models.Movie
		defer cancel()

		results, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			ErrorRes(c, err)
			return
		}

		//reading from the db
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleMovie models.Movie
			if err = results.Decode(&singleMovie); err != nil {
				ErrorRes(c, err)
				return
			}

			movie = append(movie, singleMovie)
		}

		StatusCreated(c, movie)
	}
}

func GetFilteredData() gin.HandlerFunc { //this handler allows to filter data with request body
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var reqbody EntryRecord
		var result []models.Movie

		defer cancel()

		if err := c.ShouldBindJSON(&reqbody); err != nil { //validate request body
			ErrorRes(c, err)
			return
		}

		sd, _ := time.Parse("2006-01-02", reqbody.StartDate)
		ed, _ := time.Parse("2006-01-02", reqbody.EndDate)

		if filterCursor, err := movieCollection.Find(ctx,
			bson.M{
				"release": bson.M{"$gt": sd, "$lt": ed},
				"imdb":    bson.M{"$gt": reqbody.MinIMDB, "$lt": reqbody.MaxIMDB},
			}); err != nil {
			ErrorRes(c, err)
			return
		} else {
			if err = filterCursor.All(ctx, &result); err != nil {
				ErrorRes(c, err)
				return
			}
		}

		StatusCreated(c, result)
	}
}
func EditFilmByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		movieId := c.Param("movieId")
		var movie models.Movie
		var updatedMovie models.Movie
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(movieId)

		//validate the request body
		if err := c.ShouldBindJSON(&movie); err != nil {
			ErrorRes(c, err)
			return
		}

		update := bson.M{"title": movie.Title, "director": movie.Director, "imdb": movie.IMDB, "release": movie.Release, "counts": movie.Counts}
		result, err := movieCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			ErrorRes(c, err)
			return
		}

		//get updated movie details
		if result.MatchedCount == 1 {
			if err := movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedMovie); err != nil {
				ErrorRes(c, err)
				return
			}
		}
		StatusCreated(c, result)
	}
}
func DeleteFilmByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := movieCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			ErrorRes(c, err)
			if result.DeletedCount < 1 {
				ErrorRes(c, err)
				return
			}
			return
		} else {
			StatusCreated(c, result)
		}
	}
}
