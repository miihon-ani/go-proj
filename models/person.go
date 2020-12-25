package models

import (
	"app/db"

	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age`
}

func (h Person) Insert(name string, age int) {
	db := db.GetDB()
	db.Create(&Person{Name: name, Age: age})
	defer db.Close()
}

func (h Person) GetAll() []Person {
	db := db.GetDB()
	var people []Person
	db.Find(&people)
	defer db.Close()
	return people
}

func (h Person) GetOne(id int) Person {
	db := db.GetDB()
	var people Person
	db.First(&people, id)
	defer db.Close()
	return people
}

func (h Person) Update(id int, name string, age int) {
	db := db.GetDB()
	var people Person
	db.First(&people, id)
	people.Name = name
	people.Age = age
	db.Save(&people)
	defer db.Close()
}

func (h Person) Delete(id int) {
	db := db.GetDB()
	var people Person
	db.First(&people, id)
	db.Delete(&people)
	defer db.Close()
}
