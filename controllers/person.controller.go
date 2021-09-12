package controllers

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"
	. "rest-server-test/models"
	"rest-server-test/services"
	"strconv"
)

func GetPersons(c *gin.Context) {

	idStr, hasId := c.GetQuery("id")

	if hasId {
		id, err := strconv.Atoi(idStr)

		if err == nil {
			var person = services.FindPersonByPk(id)
			c.IndentedJSON(200, person)
		} else {
			fmt.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {

		var persons []Person
		services.FindAllPersons(&persons)
		c.IndentedJSON(200, persons)
	}

}

func CreatePerson(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		return
	}
	person = services.CreatePerson(person)
	c.IndentedJSON(http.StatusCreated, person)
}

func UpdatePerson(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		return
	}
	person = services.UpdatePerson(person)
	c.IndentedJSON(http.StatusCreated, person)
}

func DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		services.RemovePersonById(id)
	} else {
		c.IndentedJSON(http.StatusNotFound, "person not found")
	}
}
