package services

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "rest-server-test/models"
)

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func init() {
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Person{})
}

func FindPersonByPk(id int) Person {
	var person Person
	db.First(&person, id)
	fmt.Println(person)
	return person
}

func FindAllPersons(persons *[]Person) {
	db.Find(persons)
}

func CreatePerson(person Person) Person {
	db.Create(&person)
	return person
}

func UpdatePerson(person Person) Person {
	db.Save(&person)
	return person
}

func RemovePerson(person Person) {
	db.Delete(person)
}

func RemovePersonById(id int) {
	db.Delete(&Person{}, id)
}
