package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MicroSOA-09/blog-service/handler"
	"github.com/MicroSOA-09/blog-service/repository"
	"github.com/MicroSOA-09/blog-service/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func startService(handler *handler.BlogPostHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/blog/blogpost", handler.GetAll).Methods("GET")
	router.HandleFunc("/api/blog/blogpost/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/api/blog/blogpost", handler.Create).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	println("Server listening on :%v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Could not load .env file, using defaults:", err)
	}

	logger := log.New(os.Stdout, "[blog-handler] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[blog-repo] ", log.LstdFlags)

	database := repository.InitDB()
	if database == nil {
		print("FAILED TO CONNECT")
		return
	}

	repo := &repository.BlogPostRepository{Db: database, Logger: storeLogger}
	service := &service.BlogPostService{BlogPostRepo: repo}
	handler := &handler.BlogPostHandler{BlogPostService: service, Logger: logger}

	startService(handler)
}
