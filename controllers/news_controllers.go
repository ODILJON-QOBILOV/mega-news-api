package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
)

// GetNewsAPI godoc
// @Summary Get News API
// @Tags news
// @Success 200 {string} string "OK"
// @Router /news [get]
func GetNewsAPI(c *gin.Context) {
	var news []models.News
	if err := config.DB.Preload("Author").Find(&news).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, news)
}

// GetNewsDetailAPI godoc
// @Summary Get news details by ID
// @Tags news
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} models.News
// @Failure 404 {string} string "News not found"
// @Router /news/{id} [get]
func GetNewsDetailAPI(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := config.DB.Preload("Author").First(&news, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "News not found"})
		return
	}

	c.JSON(200, news)
}



// GetUserNewsAPI godoc
// @Summary Get Writer's Own News
// @Tags news
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.News
// @Failure 404 {string} string "not found"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /news/mine [get]
func GetUserNewsAPI(c *gin.Context) {
    userID := c.GetInt("user_id")
    fmt.Println("User ID:", userID) // Debugging line

    var news []models.News
    if err := config.DB.Where("author_id = ?", userID).Preload("Author").Find(&news).Error; err != nil {
        fmt.Println("Error executing query:", err)
        c.JSON(404, gin.H{"error": "No news found"})
        return
    }

    if len(news) == 0 {
        c.JSON(404, gin.H{"error": "No news found"})
        return
    }

    c.JSON(200, news)
}



// CreateNewsAPI godoc
// @Summary Create a new news article
// @Tags news
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "News title"
// @Param content formData string true "News content"
// @Param image formData file true "Primary image"
// @Param secondImage formData file false "Secondary image"
// @Success 201 {object} models.News
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /news [post]
func CreateNewsAPI(c *gin.Context) {
	userID := c.GetInt("user_id")

	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" || content == "" {
		c.JSON(400, gin.H{"error": "Title and content are required"})
		return
	}

	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to create upload directory"})
		return
	}

	var filePath1, filePath2 string

	if file1, err1 := c.FormFile("image"); err1 == nil {
		filePath1 = filepath.Join("uploads", file1.Filename)
		if err := c.SaveUploadedFile(file1, filePath1); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save image file"})
			return
		}
	}

	if file2, err2 := c.FormFile("secondImage"); err2 == nil {
		filePath2 = filepath.Join("uploads", file2.Filename)
		if err := c.SaveUploadedFile(file2, filePath2); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save second image file"})
			return
		}
	}

	news := models.News{
		Title:       title,
		Content:     content,
		Image:       filePath1,
		SecondImage: filePath2,
		AuthorID:    userID,
	}

	if err := config.DB.Create(&news).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create news"})
		return
	}

	c.JSON(201, news)
}
