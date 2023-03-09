package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Records map[string]interface{} `json:"records"`
}

func ErrorRes(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    http.StatusInternalServerError,
		Message: "error",
		Records: map[string]interface{}{"data": err.Error()}})
}

func StatusCreated(c *gin.Context, result interface{}) {
	c.JSON(http.StatusCreated, StatusResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Records: map[string]interface{}{"data": result}})
}
