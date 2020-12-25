package models

import (
	"github.com/jinzhu/gorm"
)

type PersonInsert struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
