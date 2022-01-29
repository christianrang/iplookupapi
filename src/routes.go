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
	searchRoutes := r.Group("/search")
	{
		searchRoutes.GET("/ip/:ip", SearchIP)
		searchRoutes.GET("/domain/:domain", SearchDomain)
		searchRoutes.GET("/file_hash/:file_hash", SearchFileHash)
	}
}
