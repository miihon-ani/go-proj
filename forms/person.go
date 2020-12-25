package forms

type PersonInsertForm struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}
