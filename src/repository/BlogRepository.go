package repository

import (
	"context"
	"errors"
	apiModel "harmony/src/models/api"
	"harmony/src/models/db"
	"log"

	"github.com/jmoiron/sqlx"
)

type BlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]apiModel.Blog, error)
}

type blogRepository struct {
	db *sqlx.DB
}

func (r blogRepository) GetAllBlogs(ctx context.Context) ([]apiModel.Blog, error) {
	var blogs []db.Blog

	query := "SELECT * FROM BLOGS"

	err := r.db.SelectContext(ctx, &blogs, query)

	if err != nil {
		log.Printf("error when retrieving user details %v", err)
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, errors.New("no blogs added to this space yet")
	}

	var allBlogs []apiModel.Blog

	for _, blog := range blogs {

		allBlogs = append(allBlogs, apiModel.Blog{
			PublishedDate: blog.PublishedDate,
			Title:         blog.Title,
			Content:       blog.Content,
			Public:        blog.Public,
			Groups:        blog.Groups,
		})
	}

	return allBlogs, nil
}

func NewBlogRepository(db *sqlx.DB) BlogRepository {
	return &blogRepository{
		db: db,
	}
}
