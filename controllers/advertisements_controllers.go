package controllers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
)

// CreateAdvertisementAPI godoc
// @Summary Create a new advertisement
// @Tags advertisements
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Ad title"
// @Param link formData string true "Ad link"
// @Param image formData file true "Image file"
// @Success 201 {object} models.Advertisement
// @Failure 401 {string} string "unauthorized"
// @Router /advertisements [post]
func CreateAdvertisementAPI(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	
	os.MkdirAll("uploads", os.ModePerm)
	var filePath string
	file, err := c.FormFile("image")
	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}
	advertisement := models.Advertisement{
		Title: title,
		Link: link,
		Image: filePath,
	}
	config.DB.Create(&advertisement)
	c.JSON(201, advertisement)
}



// DeleteAdvertisementAPI godoc
// @Summary Delete advertisement by ID
// @Tags advertisements
// @Produce json
// @Param id path int true "Advertisement ID"
// @Security BearerAuth
// @Success 200 {object} models.Advertisement
// @Failure 404 {string} string "Advertisement not found"
// @Router /advertisements/{id} [delete]
func DeleteAdvertisementAPI(c *gin.Context) {
	id := c.Param("id")
	var ad models.Advertisement
	if err := config.DB.Delete(&ad, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "deleted"})
}

// GetAdvertisementsAPI godoc
// @Summary Get Advertisements API
// @Tags advertisements
// @Success 200 {string} string "OK"
// @Router /advertisements [get]
func GetAdvertisementsAPI(c *gin.Context) {
	var ad []models.Advertisement
	config.DB.Find(&ad)
	c.JSON(200, ad)
}


// GetAdvertisementDetailAPI godoc
// @Summary Get ad details by ID
// @Tags advertisements
// @Produce json
// @Param id path int true "Advertisement ID"
// @Success 200 {object} models.Advertisement
// @Failure 404 {string} string "Advertisement not found"
// @Router /advertisements/{id} [get]
func GetAdvertisementDetailAPI(c *gin.Context) {
	id := c.Param("id")
	var ad models.Advertisement
	if err := config.DB.First(&ad, id).Error; err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, ad)
}