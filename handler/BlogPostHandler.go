package handler

import (
	"encoding/json"
	"fmt"
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

	for i := 0; i < len(blogs); i++ {
		err = populateBlog(&blogs[i], handler)
		if err != nil {
			println("Error while fetching Users")
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
	err = populateBlog(blog, handler)
	if err != nil {
		println("Error while fetching Users")
		writter.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(blog)
}

func (handler *BlogPostHandler) Create(writter http.ResponseWriter, req *http.Request) {
	println("KREIRANJE")
	var blog model.BlogPost
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing JSON")
		writter.WriteHeader(http.StatusBadRequest)
		return
	}
	println(blog.Title)
	err = handler.BlogPostService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writter.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writter.WriteHeader(http.StatusCreated)
	writter.Header().Set("Content-Type", "application/json")
}

func populateBlog(blog *model.BlogPost, handler *BlogPostHandler) error {
	userIds := handler.BlogPostService.GetUserIds(blog)
	if len(userIds) > 0 {
		param := handler.BlogPostService.IdsToStr(userIds)
		users, err := fetchUsersFromStakeholders(param)
		if err != nil {
			return err
		}
		handler.BlogPostService.PopulateBlog(blog, users)
	}
	return nil
}
func fetchUsersFromStakeholders(param string) (map[int]string, error) {
	resp, err := http.Get(fmt.Sprintf("http://explorer:80/api/user/getUsernames/%s", param))
	if err != nil {
		println("Error with http request")
		return nil, err
	}

	defer resp.Body.Close()

	var response model.PagedResult[model.User]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		println("Error with mapping JSON to Response:", err.Error())
		return nil, err
	}

	userMap := make(map[int]string)
	for _, user := range response.Results {
		userMap[user.ID] = user.Username
	}
	return userMap, nil
}
