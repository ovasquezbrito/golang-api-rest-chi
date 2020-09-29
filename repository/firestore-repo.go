package repository

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
	"github.com/ovasquezbrito/golang-rest-api/entity"
)

var connStr = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

type repo struct{}

// NewFirestoreRepository .
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "golang-rest-api-f9ebf"
	collectionName string = "posts"
)

var opt = option.WithCredentialsFile(connStr)

// Save .
func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
