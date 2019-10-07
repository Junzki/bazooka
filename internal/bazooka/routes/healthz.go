package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
