package main

import controllers "github.com/bhatia4/enrollee-services-go/controllers"
import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	controllers.RouterSetup(router)
}