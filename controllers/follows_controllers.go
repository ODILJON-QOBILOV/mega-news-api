package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
)

// FollowWriterAPI godoc
// @Summary User follows a writer
// @Tags follows
// @Param writer_id path int true "ID of the writer"
// @Security BearerAuth
// @Success 200 {object} string "Successfully followed"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /follow/{writer_id} [post]
func FollowWriterAPI(c *gin.Context) {
	user_id := c.GetInt("user_id")
	writerIDStr := c.Param("writer_id")
    writer_id, err := strconv.Atoi(writerIDStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid writer ID"})
        return
    }

	if user_id == writer_id {
		c.JSON(400, gin.H{"error": "you cannot follow to yourself"})
		return
	}

	var existingFollow models.Follow
	if err := config.DB.Where("user_id = ? AND writer_id = ?", user_id, writer_id).First(&existingFollow).Error; err == nil {
		c.JSON(400, gin.H{"error": "you are already following this user"})
		return
	}

	follow := models.Follow{
		UserId: user_id,
		WriterId: writer_id,
	}
	if err := config.DB.Create(&follow).Error; err != nil {
		c.JSON(500, gin.H{"error": "couldn't follow"})
		return
	}
	c.JSON(200, gin.H{"message": "successfully followed"})
}


// UnfollowWriterAPI godoc
// @Summary User unfollows a writer
// @Tags follows
// @Param writer_id path int true "ID of the writer"
// @Security BearerAuth
// @Success 200 {object} string "Successfully unfollowed"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /unfollow/{writer_id} [post]
func UnfollowWriterAPI(c *gin.Context) {
    userID := c.GetInt("user_id")
    writerID := c.Param("writer_id")

    if err := config.DB.Where("user_id = ? AND writer_id = ?", userID, writerID).Delete(&models.Follow{}).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to unfollow writer"})
        return
    }

    c.JSON(200, gin.H{"message": "Successfully unfollowed"})
}
