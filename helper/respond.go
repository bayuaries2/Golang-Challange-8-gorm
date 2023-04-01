package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = "record not found"
)

// 200
func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func OkWithMassage(c *gin.Context, massage string) {
	c.JSON(http.StatusOK, gin.H{"massage": massage})
}

// 201
func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// not ok

func BadRequest(c *gin.Context, massage string, data ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "massage": massage}
	if len(data) > 0 {
		obj["data"] = data[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, massage string) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "massage": massage})
}

func InternalServiceError(c *gin.Context, massage string) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "massage": massage})
}
