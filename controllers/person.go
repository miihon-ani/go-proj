package controllers

import (
	"app/forms"
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
	var form forms.PersonInsertForm
	if err := ctx.Bind(&form); err != nil {
		// 実際は自ページに戻しつつwarning吐くのが正解だと思うがひとまず割愛
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest (Failed Form Validate)"})
		return
	}
	personModel.Insert(form)
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
	var form forms.PersonInsertForm
	if err := ctx.Bind(&form); err != nil {
		// 実際は自ページに戻しつつwarning吐くのが正解だと思うがひとまず割愛
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest (Failed Form Validate)"})
		return
	}
	personModel.Update(id, form)
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
