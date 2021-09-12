package main

import (
	"rest-server-test/controllers"
)
import "github.com/gin-gonic/gin"

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/persons", controllers.GetPersons)
	r.POST("/persons", controllers.CreatePerson)
	r.PUT("/persons", controllers.UpdatePerson)
	r.DELETE("/persons", controllers.DeletePerson)

	r.Run() // listen and serve on 0.0.0.0:8080

}
