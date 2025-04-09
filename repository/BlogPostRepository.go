package repository

import (
	"github.com/MicroSOA-09/blog-service/model"
	"gorm.io/gorm"
)

type BlogPostRepository struct {
	Db *gorm.DB
}

func (repo *BlogPostRepository) FindAll() ([]model.BlogPost, error) {
	blogs := []model.BlogPost{}
	dbResult := repo.Db.Find(&blogs)
	if dbResult != nil {
		return blogs, dbResult.Error
	}
	return blogs, nil
}

func (repo *BlogPostRepository) FindById(id string) (model.BlogPost, error) {
	blog := model.BlogPost{}
	dbResult := repo.Db.First(&blog, "id = ?", id)
	if dbResult != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (repo *BlogPostRepository) CreateBlogPost(blog *model.BlogPost) error {
	dbResult := repo.Db.Create(blog)
	if dbResult != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
