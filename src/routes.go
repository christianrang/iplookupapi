package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})
	r.GET("/ip/:ip", func(context *gin.Context) {
		SearchIP(context)
	})
	r.GET("/domain/:domain", func(context *gin.Context) {
		SearchDomain(context)
	})
}
