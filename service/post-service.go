package service

import (
	"errors"
	"math/rand"

	"github.com/ovasquezbrito/golang-rest-api/entity"
	"github.com/ovasquezbrito/golang-rest-api/repository"
)

// PostService .
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

// NewPostService .
func NewPostService(reposi repository.PostRepository) PostService {
	repo = reposi
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post Title is empty")
		return err
	}

	if post.Text == "" {
		err := errors.New("The post Text is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
