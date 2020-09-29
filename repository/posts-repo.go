package repository

import (
	"github.com/ovasquezbrito/golang-rest-api/entity"
)

// PostRepository .
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
