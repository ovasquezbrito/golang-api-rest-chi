package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ovasquezbrito/golang-rest-api/errors"
	"github.com/ovasquezbrito/golang-rest-api/service"

	"github.com/ovasquezbrito/golang-rest-api/entity"
)

type controller struct{}

var (
	postService service.PostService
)

// PostController .
type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

// NewPostController .
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// GetPosts .
func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

// AddPost .
func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
