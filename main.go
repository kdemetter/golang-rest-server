package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

import "github.com/gin-gonic/gin"

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Person{})

	r := gin.Default()

	r.GET("/persons", func(c *gin.Context) {
		if c.QueryMap("id") != nil {
			var person []Person
			id := c.Query("id")
			db.First(&person, id)
			c.IndentedJSON(200, person)
		} else {
			var persons []Person
			db.Find(&persons)
			c.IndentedJSON(200, persons)
		}

	})

	r.POST("/persons", func(c *gin.Context) {
		var person Person
		if err := c.BindJSON(&person); err != nil {
			return
		}
		db.Create(&person)
		c.IndentedJSON(http.StatusCreated, person)
	})

	r.PUT("/persons", func(c *gin.Context) {
		var person Person
		if err := c.BindJSON(&person); err != nil {
			return
		}
		db.Save(&person)
		c.IndentedJSON(http.StatusCreated, person)
	})


	r.DELETE("/persons", func(c *gin.Context) {
		var person Person
		id := c.Query("id")
		db.First(&person, id)

		db.Delete(&person, id)

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
