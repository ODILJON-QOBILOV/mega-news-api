package controllers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
)



// CreateBannerAPI godoc
// @Summary Create a new banner
// @Tags banners
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Banner title"
// @Param description formData string true "Banner description"
// @Param image formData file true "Image file"
// @Success 201 {object} models.Banner
// @Failure 401 {string} string "unauthorized"
// @Failure 400 {string} string "bad request"
// @Router /banners [post]
func CreateBannerAPI(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	
	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")

	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}
	banner := models.Banner{
		Title:       title,
		Description: description,
		Image:       filePath,
	}
	config.DB.Create(&banner)
	c.JSON(201, banner)
}


// GetBannersAPI godoc
// @Summary Get Banners API
// @Tags banners
// @Success 200 {string} string "OK"
// @Router /banners [get]
func GetBannersAPI(c *gin.Context) {
	var banner []models.Banner
	config.DB.Find(&banner)
	c.JSON(200, banner)
}

// UpdateBannerAPI godoc
// @Summary Update banner details by ID
// @Tags banners
// @Produce json
// @Param id path int true "Banner ID"
// @Param title formData string false "Banner Title"
// @Param description formData string false "Banner description"
// @Param image formData file false "Banner Image"
// @Security BearerAuth
// @Success 200 {object} models.Banner
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Banner not found"
// @Failure 500 {string} string "Internal server error"
// @Router /banners/{id} [patch]
func UpdateBannerAPI(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	description := c.PostForm("description")

	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")
	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}

	var banner models.Banner
	if err := config.DB.First(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Banner not found"})
		return
	}

	if title != "" {
		banner.Title = title
	}
	if description != "" {
		banner.Description = description
	}
	if filePath != "" {
		banner.Image = filePath
	}

	if err := config.DB.Save(&banner).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(200, banner)
}


// DeleteBannerAPI godoc
// @Summary Delete banner details by ID
// @Tags banners
// @Produce json
// @Param id path int true "Banner ID"
// @Security BearerAuth
// @Success 200 {object} models.Banner
// @Failure 404 {string} string "Banner not found"
// @Router /banners/{id} [delete]
func DeleteBannerAPI(c *gin.Context) {
	id := c.Param("id")

	var banner models.Banner
	if err := config.DB.Delete(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "deleted"})
}


// GetBannerDetailAPI godoc
// @Summary Get Baner API
// @Tags banners
// @Param id path int true "Banner ID"
// @Success 200 {string} string "OK"
// @Router /banners/{id} [get]
func GetBannerDetailAPI(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner
	if err := config.DB.First(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, banner)
}


// CreateBannerCarouselAPI godoc
// @Summary Create a new banner-carousel
// @Tags banners-carousel
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Banner title"
// @Param description formData string true "Banner description"
// @Param image formData file true "Image file"
// @Success 201 {object} models.BannerCarousel
// @Failure 401 {string} string "unauthorized"
// @Failure 400 {string} string "bad request"
// @Router /banners/carousel [post]
func CreateBannerCarouselAPI(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	
	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")

	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}
	banner := models.BannerCarousel{
		Title:       title,
		Description: description,
		Image:       filePath,
	}
	config.DB.Create(&banner)
	c.JSON(201, banner)
}


// GetBannerCarouselDetailAPI godoc
// @Summary Get Baner Carousel API
// @Tags banners-carousel
// @Param id path int true "Banner Carousel ID"
// @Success 200 {string} string "OK"
// @Router /banners/carousel/{id} [get]
func GetBannerCarouselDetailAPI(c *gin.Context) {
	id := c.Param("id")
	var banner models.BannerCarousel
	if err := config.DB.First(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, banner)
}

// DeleteBannerCarouselAPI godoc
// @Summary Delete banner carousel details by ID
// @Tags banners-carousel
// @Produce json
// @Param id path int true "Banner Carousel ID"
// @Security BearerAuth
// @Success 200 {object} models.BannerCarousel
// @Failure 404 {string} string "Banner Carousel not found"
// @Router /banners/carousel/{id} [delete]
func DeleteBannerCarouselAPI(c *gin.Context) {
	id := c.Param("id")

	var banner models.BannerCarousel
	if err := config.DB.Delete(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "deleted"})
}

// GetBannersCarouselAPI godoc
// @Summary Get Banners Carousel API
// @Tags banners-carousel
// @Success 200 {string} string "OK"
// @Router /banners/carousel [get]
func GetBannersCarouselAPI(c *gin.Context) {
	var banner []models.Banner
	config.DB.Find(&banner)
	c.JSON(200, banner)
}


// UpdateBannerCarouselAPI godoc
// @Summary Update banner Carousel details by ID
// @Tags banners-carousel
// @Produce json
// @Param id path int true "Banner ID"
// @Param title formData string false "Banner Carousel Title"
// @Param description formData string false "Banner Carousel description"
// @Param image formData file false "Banner Carousel Image"
// @Security BearerAuth
// @Success 200 {object} models.BannerCarousel
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Banner not found"
// @Failure 500 {string} string "Internal server error"
// @Router /banners/carousel/{id} [patch]
func UpdateBannerCarouselAPI(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	description := c.PostForm("description")

	os.MkdirAll("uploads", os.ModePerm)

	var filePath string
	file, err := c.FormFile("image")
	if err == nil {
		filePath = filepath.Join("uploads", file.Filename)
		c.SaveUploadedFile(file, filePath)
	}

	var banner models.BannerCarousel
	if err := config.DB.First(&banner, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Banner not found"})
		return
	}

	if title != "" {
		banner.Title = title
	}
	if description != "" {
		banner.Description = description
	}
	if filePath != "" {
		banner.Image = filePath
	}

	if err := config.DB.Save(&banner).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(200, banner)
}