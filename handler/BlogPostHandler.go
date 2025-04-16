package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MicroSOA-09/blog-service/model"
	"github.com/MicroSOA-09/blog-service/service"
	"github.com/gorilla/mux"
)

type BlogPostHandler struct {
	BlogPostService *service.BlogPostService
	Logger          *log.Logger
}

func (handler *BlogPostHandler) GetAll(writter http.ResponseWriter, req *http.Request) {
	blogs, err := handler.BlogPostService.FindAll()
	writter.Header().Set("Content-Type", "application/json")
	if err != nil {
		handler.Logger.Printf("Error while fetching blogs")
		writter.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(blogs); i++ {
		err = handler.populateBlog(&blogs[i])
		if err != nil {
			writter.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	var result model.PagedResult[model.BlogPost]
	result.Results = blogs
	result.TotalCount = len(blogs)

	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(result)
}

func (handler *BlogPostHandler) Get(writter http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	blog, err := handler.BlogPostService.FindBlog(id)
	writter.Header().Set("Content-Type", "application/json")
	if err != nil {
		writter.WriteHeader(http.StatusNotFound)
		return
	}
	err = handler.populateBlog(blog)
	if err != nil {
		handler.Logger.Printf("Error  while fetching Users")
		writter.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(blog)
}

func (handler *BlogPostHandler) Create(writter http.ResponseWriter, req *http.Request) {
	var blog model.BlogPost
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		handler.Logger.Printf("Error while parsing JSON")
		writter.WriteHeader(http.StatusBadRequest)
		return
	}
	println(blog.Title)
	err = handler.BlogPostService.Create(&blog)
	if err != nil {
		handler.Logger.Printf("Error while creating a new blog")
		writter.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writter.WriteHeader(http.StatusCreated)
	writter.Header().Set("Content-Type", "application/json")
}

func (handler *BlogPostHandler) populateBlog(blog *model.BlogPost) error {
	userIds := handler.BlogPostService.GetUserIds(blog)
	if len(userIds) > 0 {
		param := strings.Join(userIds, ",")
		users, err := handler.fetchUsersFromStakeholders(param)
		if err != nil {
			return err
		}
		handler.BlogPostService.PopulateBlog(blog, users)
	}
	return nil
}
func (handler *BlogPostHandler) fetchUsersFromStakeholders(param string) (map[string]string, error) {
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	resp, err := http.Get(fmt.Sprintf("http://%s/api/user/getUsernames/%s", userServiceURL, param))
	if err != nil {
		handler.Logger.Printf("Error with http request")
		return nil, err
	}

	defer resp.Body.Close()

	var response model.PagedResult[model.User]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		handler.Logger.Printf("Error with mapping JSON to Response: %v", err.Error())
		return nil, err
	}

	userMap := make(map[string]string)
	for _, user := range response.Results {
		userMap[user.ID] = user.Username
	}
	return userMap, nil
}
