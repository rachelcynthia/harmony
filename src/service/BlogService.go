package service

import (
	apiModel "harmony/src/models/api"
	"harmony/src/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type BlogService interface {
	GetAllBlogs(ctx *gin.Context) ([]apiModel.Blog, error)
}

type blogService struct {
	blogRepository repository.BlogRepository
}

func (s blogService) GetAllBlogs(ctx *gin.Context) ([]apiModel.Blog, error) {
	blogs, err := s.blogRepository.GetAllBlogs(ctx.Request.Context())

	if err != nil {
		log.Printf("error when accessing db: %v", err)
		return nil, err
	}
	return blogs, nil
}

func NewBlogService(blogRepository repository.BlogRepository) BlogService {
	return blogService{
		blogRepository: blogRepository,
	}
}
