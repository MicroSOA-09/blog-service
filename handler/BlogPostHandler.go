package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MicroSOA-09/blog-service/model"
	"github.com/MicroSOA-09/blog-service/service"
	"github.com/gorilla/mux"
)

type BlogPostHandler struct {
	BlogPostService *service.BlogPostService
}

func (handler *BlogPostHandler) GetAll(writter http.ResponseWriter, req *http.Request) {
	blogs, err := handler.BlogPostService.FindAll()
	writter.Header().Set("Content-Type", "application/json")
	if err != nil {
		writter.WriteHeader(http.StatusNotFound)
		return
	}
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(blogs)
}

func (handler *BlogPostHandler) Get(writter http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	blog, err := handler.BlogPostService.FindBlog(id)
	writter.Header().Set("Content-Type", "application/json")
	if err != nil {
		writter.WriteHeader(http.StatusNotFound)
		return
	}
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(blog)
}

func (handler *BlogPostHandler) Create(writter http.ResponseWriter, req *http.Request) {
	var blog model.BlogPost
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing JSON")
		writter.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogPostService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writter.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writter.WriteHeader(http.StatusCreated)
	writter.Header().Set("Content-Type", "application/json")
}
