package server

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/*")

	var person = new(controllers.PersonController)

	router.GET("/", person.List)
	router.POST("/new", person.Insert)
	router.GET("/detail/:id", person.Detail)
	router.POST("/update/:id", person.Update)
	router.GET("/delete_check/:id", person.DeleteCheck)
	router.POST("/delete/:id", person.Delete)

	return router
}
