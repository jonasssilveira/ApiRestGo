package controllers

import (
	"booksAPiProject/database"
	"booksAPiProject/goi18n"
	"booksAPiProject/models"
	"github.com/gin-gonic/gin"
)

var db = database.StardDB()

func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var book models.Book

		err := db.First(&book, "id =?", id).Error
		if err != nil {
			ApiError(c, 404, "book with id "+id+" not found")
			return
		}
		c.JSON(200, book)
	}
}
func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		book := models.NewBook()
		err := c.ShouldBindJSON(&book)
		if err != nil {
			ApiError(c, 400, "nao foi possivel realizar o bind "+err.Error())
			return
		}
		err = db.Create(&book).Error

		if err != nil {
			ApiError(c, 40, "nao foi possivel salvar "+err.Error())
			return
		}
		c.JSON(201, book)
	}
}
func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		book := models.Book{}
		err := c.ShouldBindJSON(&book)
		if err != nil {
			ApiError(c, 400, "nao foi possivel realizar o bind "+err.Error())
			return
		}
		err = db.Save(&book).Error

		if err != nil {
			ApiError(c, 400, "nao foi possivel atualizar "+err.Error())
			return
		}
		c.JSON(200, book)
	}
}

func ApiError(c *gin.Context, status int, errorMessage string) {
	goi18n.InitMessages()
	message := goi18n.GetMessageError(errorMessage)
	c.JSON(status, gin.H{
		"error": message,
	})
}

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		err := db.Find(&books).Error
		if err != nil {
			ApiError(c, 400, err.Error()) //"Cannot list books "+err.Error())
			return
		}
		c.JSON(200, books)
	}
}
func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book

		err := db.First(&book, "id =?", id).Error
		if err != nil {
			ApiError(c, 400, "book with id "+id+" not found")
			return
		}
		err = db.Where("id=?", id).Delete(&book).Error

		if err != nil {
			ApiError(c, 400, "Não foi possivel apagar o book"+err.Error())
			return
		}
		c.JSON(204, gin.H{"id": id})
	}
}
