package controllers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
)


// GetCategoriesAPI godoc
// @Summary Get Categories API
// @Tags categories
// @Success 200 {string} string "OK"
// @Router /categories [get]
func GetCategoriesAPI(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(200, categories)
}


// GetCategoryDetailAPI godoc
// @Summary Get category details by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {string} string "Category not found"
// @Router /categories/{id} [get]
func GetCategoryDetailAPI(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, category)
}


// CreateCategoriesAPI godoc
// @Summary Create a new category
// @Tags categories
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Category title"
// @Param image formData file true "Image file"
// @Success 201 {object} models.Category
// @Failure 401 {string} string "unauthorized"
// @Router /categories [post]
func CreateCategoriesAPI(c *gin.Context) {
	title := c.PostForm("title")

	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")
	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}
	category := models.Category{
		Title: title,
		Image: filePath,
	}
	config.DB.Create(&category)
	c.JSON(201, category)
}


// DeleteCategoryAPI godoc
// @Summary Delete category details by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Security BearerAuth
// @Success 200 {object} models.Category
// @Failure 404 {string} string "Category not found"
// @Router /categories/{id} [delete]
func DeleteCategoryAPI(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.Delete(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "deleted"})
}


// UpdateCategoryAPI godoc
// @Summary Update category details by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Param title formData string false "Category Title"
// @Param image formData file false "Category Image"
// @Security BearerAuth
// @Success 200 {object} models.Category
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Category not found"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [patch]
func UpdateCategoryAPI(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")

	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")
	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}

	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	if title != "" {
		category.Title = title
	}
	if filePath != "" {
		category.Image = filePath
	}

	if err := config.DB.Save(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(200, category)
}
