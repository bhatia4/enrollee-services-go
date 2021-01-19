package controllers

import "github.com/gin-gonic/gin"

func RouterSetup(router *gin.Engine) {
	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/enrollees", addEnrollee)
		v1.PUT("/enrollees/:enrolleeid", updateEnrollee)
		v1.DELETE("/enrollees/:enrolleeid", deleteEnrollee)
		v1.POST("/enrollees/:enrolleeid/dependents", addDependent)
		v1.PUT("/enrollees/:enrolleeid/dependents/:dependentid", updateDependent)
		v1.DELETE("/enrollees/:enrolleeid/dependents/:dependentid", deleteDependent)
	}
	
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}