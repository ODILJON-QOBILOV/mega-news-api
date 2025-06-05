package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/controllers"
	"github.com/newsapi/v2/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.Static("/uploads", "./uploads")
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterAPI)
		auth.POST("/login", controllers.LoginAPI)
	}
	news := r.Group("/news")
	{
		news.GET("", controllers.GetNewsAPI)
		news.GET("/:id", controllers.GetNewsDetailAPI)
		news.GET("/mine", middleware.AuthMiddleware("writer"), controllers.GetUserNewsAPI)
		news.POST("", middleware.AuthMiddleware("writer", "admin"), controllers.CreateNewsAPI)
	}
	categories := r.Group("/categories")
	{
		categories.GET("", controllers.GetCategoriesAPI)
		categories.GET("/:id", controllers.GetCategoryDetailAPI)
		categories.POST("", middleware.AuthMiddleware("admin"),controllers.CreateCategoriesAPI)
		categories.DELETE("/:id", middleware.AuthMiddleware("admin"), controllers.DeleteCategoryAPI)
		categories.PATCH("/:id", middleware.AuthMiddleware("admin"), controllers.UpdateCategoryAPI)
	}
	follow := r.Group("/")
	{
		follow.POST("follow/:writer_id", middleware.AuthMiddleware("user"),controllers.FollowWriterAPI)
		follow.POST("unfollow/:writer_id", middleware.AuthMiddleware("user"), controllers.UnfollowWriterAPI)
	}
	banner := r.Group("/banners")
	{
		banner.POST("", middleware.AuthMiddleware("admin"), controllers.CreateBannerAPI)
		banner.GET("", controllers.GetBannersAPI)
		banner.PATCH("/:id", middleware.AuthMiddleware("admin"), controllers.UpdateBannerAPI)
		banner.DELETE("/:id", middleware.AuthMiddleware("admin"), controllers.DeleteBannerAPI)
		banner.GET("/:id", controllers.GetBannerDetailAPI)
	}
	banner_carousel := r.Group("/banners/carousel")
	{
		banner_carousel.GET("", controllers.GetBannersCarouselAPI)
		banner_carousel.POST("", middleware.AuthMiddleware("admin"), controllers.CreateBannerCarouselAPI)
		banner_carousel.GET("/:id", controllers.GetBannerCarouselDetailAPI)
		banner_carousel.DELETE("/:id", middleware.AuthMiddleware("admin"), controllers.DeleteBannerCarouselAPI)
	}
	advertisement := r.Group("/advertisements")
	{
		advertisement.GET("", controllers.GetAdvertisementsAPI)
		advertisement.GET("/:id", controllers.GetAdvertisementDetailAPI)
		advertisement.POST("", middleware.AuthMiddleware("admin"), controllers.CreateAdvertisementAPI)
		advertisement.DELETE("/:id", middleware.AuthMiddleware("admin") ,controllers.DeleteAdvertisementAPI)
	}
}