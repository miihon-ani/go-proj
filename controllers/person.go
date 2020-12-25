package controllers

import (
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct{}

var personModel = new(models.Person)

func (p PersonController) List(ctx *gin.Context) {
	people := personModel.GetAll()
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
	})
}

func (p PersonController) Insert(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	personModel.Insert(name, age)
	ctx.Redirect(http.StatusFound, "/")
}

func (p PersonController) Detail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	people := personModel.GetOne(id)
	ctx.HTML(http.StatusOK, "detail.tmpl", gin.H{"people": people})
}

func (p PersonController) Update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	name := ctx.PostForm("name")
	age_str := ctx.PostForm("age")
	age, err := strconv.Atoi(age_str)
	if err != nil {
		panic("ERROR")
	}
	personModel.Update(id, name, age)
	ctx.Redirect(http.StatusFound, "/")
}

func (p PersonController) DeleteCheck(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	people := personModel.GetOne(id)
	ctx.HTML(http.StatusOK, "delete.tmpl", gin.H{"people": people})
}

func (p PersonController) Delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	personModel.Delete(id)
	ctx.Redirect(http.StatusFound, "/")
}
