package service

import (
	"fmt"

	"github.com/MicroSOA-09/blog-service/model"
	"github.com/MicroSOA-09/blog-service/repository"
)

type BlogPostService struct {
	BlogPostRepo *repository.BlogPostRepository
}

func (service *BlogPostService) FindAll() ([]model.BlogPost, error) {
	blogs, err := service.BlogPostRepo.FindAll()
	// HTTP REQ to ASP.NET application to get author info
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("there are no blogs"))
	}
	return blogs, nil
}

func (service *BlogPostService) FindBlog(id string) (*model.BlogPost, error) {
	blog, err := service.BlogPostRepo.FindById(id)
	// HTTP REQ to ASP.NET application to get author info
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog with id %s not found", id))
	}
	return &blog, nil
}

func (service *BlogPostService) Create(blog *model.BlogPost) error {
	err := service.BlogPostRepo.CreateBlogPost(blog)
	if err != nil {
		return err
	}
	return nil
}
