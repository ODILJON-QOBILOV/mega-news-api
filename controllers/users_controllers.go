package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/models"
	"golang.org/x/crypto/bcrypt"
)


// RegisterAPI godoc
// @Summary Register
// @Description Register a new user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegisterInput true "User credentials"
// @Success 201 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
func RegisterAPI(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
        Role     string `json:"role"`
    }

    if err := c.BindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    // Check if username already exists
    var existingUser models.User
    if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
        // If a user with the given username exists, return an error
        c.JSON(400, gin.H{"error": "Username already exists"})
        return
    }

    hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

    user := models.User{
        Username: input.Username,
        Password: string(hash),
        Role:     input.Role,
    }

    result := config.DB.Create(&user)
    if result.Error != nil {
        c.JSON(400, gin.H{"error": result.Error.Error()})
        return
    }

    token, _ := config.GenerateJWT(uint(user.Id), user.Username, user.Role)

    c.JSON(201, gin.H{
        "access_token": token,
    })
}



// LoginAPI godoc
// @Summary Register
// @Description Login user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserLoginInput true "User credentials"
// @Success 202 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/login [post]
func LoginAPI(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BindJSON(&input)

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := config.GenerateJWT(uint(user.Id), user.Username, user.Role)

	c.JSON(202, gin.H{
		"access_token": token,
	})
}
