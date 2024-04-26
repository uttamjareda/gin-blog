package routes

import (
	"go-blog/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	// CRUD routes for the blog
	router.POST("/blogs", controllers.CreateBlog)
	router.GET("/blogs/:id", controllers.GetBlogByID)
	router.GET("/blogs", controllers.GetAllBlogs)
	router.PUT("/blogs/:id", controllers.UpdateBlog)
	router.DELETE("/blogs/:id", controllers.DeleteBlog)
}
