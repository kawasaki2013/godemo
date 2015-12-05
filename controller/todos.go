package controller

import (
	"github.com/gin-gonic/gin"
	"godemo/database"
	"godemo/model"
	"net/http"
)

type todos struct{}

var Todos todos = todos{}

func (u *todos) List(c *gin.Context) {
	var todos []model.Todo

	db := database.GetDB()
	user := c.MustGet("user").(model.User)

	q := db.Where("user_id = ?", user.ID)
	if err := q.Find(&todos).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func (u *todos) Create(c *gin.Context) {
	var todo model.Todo

	if c.BindJSON(&todo) != nil {
		return
	}

	db := database.GetDB()
	user := c.MustGet("user").(model.User)

	todo.UserID = user.ID
	if err := db.Save(&todo).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}
