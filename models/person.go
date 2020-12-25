package models

import (
	"app/db"
	"app/forms"

	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func (h Person) Insert(form forms.PersonInsertForm) {
	db := db.GetDB()
	db.Create(&Person{Name: form.Name, Age: form.Age})
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

func (h Person) Update(id int, form forms.PersonInsertForm) {
	db := db.GetDB()
	var people Person
	db.First(&people, id)
	people.Name = form.Name
	people.Age = form.Age
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
