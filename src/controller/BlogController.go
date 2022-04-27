package controller

import (
	"harmony/src/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController interface {
	GetAllBlogs(ctx *gin.Context)
}

type blogController struct {
	blogService service.BlogService
}

func (c blogController) GetAllBlogs(ctx *gin.Context) {
	blogs, err := c.blogService.GetAllBlogs(ctx)

	if err != nil {
		log.Printf("error in service %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func NewBlogController(blogService service.BlogService) BlogController {
	return blogController{
		blogService: blogService,
	}
}
