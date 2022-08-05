package repository

import (
	"errors"
	"opiana_code_test/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(b entity.Post) (entity.Post, error)
	UpdatePost(b entity.Post) (entity.Post, error)
	DeletePost(b entity.Post)
	AllPost() []entity.Post
	FindPostByID(postID uint64) entity.Post
}

type postConnection struct {
	connection *gorm.DB
}

// NewPostRepository creates an isntance PostRepository
func NewPostRepository(dbConn *gorm.DB) PostRepository {
	return &postConnection{
		connection: dbConn,
	}
}

func (db *postConnection) InsertPost(b entity.Post) (entity.Post, error) {
	if db.connection.Save(&b).Error != nil {
		return entity.Post{}, errors.New("Failed to insert data")
	}

	// untuk mengambil object user, liat post entity
	db.connection.Preload("User").Preload("PostType").Find(&b)
	return b, nil
}

func (db *postConnection) UpdatePost(b entity.Post) (entity.Post, error) {

	if db.connection.Save(&b).Error != nil {
		return entity.Post{}, errors.New("Failed to update data")
	}

	// untuk mengambil object user, liat post entity
	db.connection.Preload("User").Preload("PostType").Find(&b)
	return b, nil
}

func (db *postConnection) DeletePost(b entity.Post) {
	db.connection.Delete(&b)
}

func (db *postConnection) AllPost() []entity.Post {
	var posts = []entity.Post{}
	db.connection.Find(&posts)

	return posts
}

func (db *postConnection) FindPostByID(postID uint64) entity.Post {
	var post entity.Post
	db.connection.Preload("User").Preload("PostType").Find(&post, postID)

	return post
}
