package service

import (
	"fmt"
	"strconv"
	"strings"

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

func (service *BlogPostService) GetUserIds(blog *model.BlogPost) []int {
	userIds := []int{}
	userIds = append(userIds, blog.AuthorID)

	for i := 0; i < len(blog.Comments); i += 1 {
		userIds = append(userIds, blog.Comments[i].AuthorID)
	}
	for i := 0; i < len(blog.Ratings); i += 1 {
		userIds = append(userIds, blog.Ratings[i].AuthorID)
	}
	return userIds
}
func (service *BlogPostService) IdsToStr(ids []int) string {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.Itoa(id)
	}
	return strings.Join(strIds, ",")
}

func (service *BlogPostService) PopulateBlog(blog *model.BlogPost, users map[int]string) {
	blog.AuthorUsername = users[blog.AuthorID]
	for i := 0; i < len(blog.Comments); i += 1 {
		blog.Comments[i].AuthorUsername = users[blog.Comments[i].AuthorID]
	}
	for i := 0; i < len(blog.Ratings); i += 1 {
		blog.Ratings[i].AuthorUsername = users[blog.Ratings[i].AuthorID]
	}
}