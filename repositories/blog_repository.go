package repositories

import (
	"database/sql"
	"fmt"
	"go-blog/db"
	"go-blog/models"
	"time"
)

type BlogRepository interface {
	CreateBlog(blog models.Blog) (int, error)
	GetBlogByID(id int) (models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
	UpdateBlog(blog models.Blog) error
	DeleteBlog(id int) error
}

type blogRepository struct {
	conn *sql.DB
}

func NewBlogRepository() BlogRepository {
	return &blogRepository{conn: db.Connect()}
}

func (r *blogRepository) CreateBlog(blog models.Blog) (int, error) {
	query := `INSERT INTO blogs (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.conn.QueryRow(query, blog.Title, blog.Content, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not execute insert operation: %w", err)
	}
	return id, nil
}

func (r *blogRepository) GetBlogByID(id int) (models.Blog, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM blogs WHERE id = $1`
	var blog models.Blog
	err := r.conn.QueryRow(query, id).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return models.Blog{}, fmt.Errorf("could not fetch blog with id %d: %w", id, err)
	}
	return blog, nil
}

func (r *blogRepository) GetAllBlogs() ([]models.Blog, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM blogs`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch blogs: %w", err)
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		err = rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not read blog data: %w", err)
		}
		blogs = append(blogs, blog)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while reading blogs: %w", err)
	}
	return blogs, nil
}

func (r *blogRepository) UpdateBlog(blog models.Blog) error {
	query := `UPDATE blogs SET title = $1, content = $2, updated_at = $3 WHERE id = $4`
	_, err := r.conn.Exec(query, blog.Title, blog.Content, time.Now(), blog.ID)
	if err != nil {
		return fmt.Errorf("could not update blog with id %d: %w", blog.ID, err)
	}
	return nil
}

func (r *blogRepository) DeleteBlog(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete blog with id %d: %w", id, err)
	}
	return nil
}
