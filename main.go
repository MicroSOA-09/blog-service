package main

import (
	"log"
	"net/http"

	"github.com/MicroSOA-09/blog-service/handler"
	"github.com/MicroSOA-09/blog-service/repository"
	"github.com/MicroSOA-09/blog-service/service"
	"github.com/gorilla/mux"
)

func startService(handler *handler.BlogPostHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/blog/blogpost", handler.GetAll).Methods("GET")
	router.HandleFunc("/api/blog/blogpost/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/api/blog/blogpost", handler.Create).Methods("POST")

	println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	database := repository.InitDB()
	if database == nil {
		print("FAILED TO CONNECT")
		return
	}

	repo := &repository.BlogPostRepository{Db: database}
	service := &service.BlogPostService{BlogPostRepo: repo}
	handler := &handler.BlogPostHandler{BlogPostService: service}

	startService(handler)
}
