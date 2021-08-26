package controllers

import (
	"github/thallesrangel/api-book-go/database"
	"github/thallesrangel/api-book-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllBooks(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(201, book)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBooks(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
