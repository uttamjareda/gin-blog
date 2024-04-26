package controllers

import (
	"go-blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBlog handles POST requests for creating new blogs
func CreateBlog(c *gin.Context) {
	var blog services.BlogData
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := services.CreateBlog(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetBlogByID handles GET requests to retrieve a blog by ID
func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	blog, err := services.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// GetAllBlogs handles GET requests to list all blogs
func GetAllBlogs(c *gin.Context) {
	blogs, err := services.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// UpdateBlog handles PUT requests to update a blog
func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog services.BlogData
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateBlog(id, blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteBlog handles DELETE requests to remove a blog
func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
