package services

import (
	"go-blog/models"
	"go-blog/repositories"
	"strconv"
)

// BlogData is the data transfer object for blog data
type BlogData struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var blogRepo repositories.BlogRepository

// Initialize the service with the repository
func init() {
	blogRepo = repositories.NewBlogRepository()
}

// CreateBlog creates a new blog entry
func CreateBlog(data BlogData) (int, error) {
	blog := models.Blog{
		Title:   data.Title,
		Content: data.Content,
	}
	return blogRepo.CreateBlog(blog)
}

// GetBlogByID returns a single blog by its ID
func GetBlogByID(id string) (BlogData, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return BlogData{}, err
	}
	blog, err := blogRepo.GetBlogByID(intID)
	if err != nil {
		return BlogData{}, err
	}
	return BlogData{
		ID:        strconv.Itoa(blog.ID),
		Title:     blog.Title,
		Content:   blog.Content,
		CreatedAt: blog.CreatedAt.Format("2006-01-02T15:04:05"),
		UpdatedAt: blog.UpdatedAt.Format("2006-01-02T15:04:05"),
	}, nil
}

// GetAllBlogs returns all blogs
func GetAllBlogs() ([]BlogData, error) {
	blogs, err := blogRepo.GetAllBlogs()
	if err != nil {
		return nil, err
	}
	var result []BlogData
	for _, blog := range blogs {
		result = append(result, BlogData{
			ID:        strconv.Itoa(blog.ID),
			Title:     blog.Title,
			Content:   blog.Content,
			CreatedAt: blog.CreatedAt.Format("2006-01-02T15:04:05"),
			UpdatedAt: blog.UpdatedAt.Format("2006-01-02T15:04:05"),
		})
	}
	return result, nil
}

// UpdateBlog updates an existing blog
func UpdateBlog(id string, data BlogData) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	blog := models.Blog{
		ID:      intID,
		Title:   data.Title,
		Content: data.Content,
	}
	return blogRepo.UpdateBlog(blog)
}

// DeleteBlog deletes a blog by its ID
func DeleteBlog(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return blogRepo.DeleteBlog(intID)
}
