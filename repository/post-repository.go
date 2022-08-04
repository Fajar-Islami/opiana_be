package repository

import (
	"opiana_code_test/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(b entity.Post) entity.Post
	UpdatePost(b entity.Post) entity.Post
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

func (db *postConnection) InsertPost(b entity.Post) entity.Post {
	db.connection.Save(&b)

	// untuk mengambil object user, liat post entity
	db.connection.Preload("User").Preload("PostType").Find(&b)
	return b
}

func (db *postConnection) UpdatePost(b entity.Post) entity.Post {

	db.connection.Save(&b)
	// untuk mengambil object user, liat post entity
	db.connection.Preload("User").Preload("PostType").Find(&b)
	return b
}

func (db *postConnection) DeletePost(b entity.Post) {
	db.connection.Delete(&b)
}

func (db *postConnection) AllPost() []entity.Post {
	var posts []entity.Post
	db.connection.Preload("User").Preload("PostType").Find(&posts)
	return posts
}

func (db *postConnection) FindPostByID(postID uint64) entity.Post {
	var post entity.Post
	db.connection.Preload("User").Preload("PostType").Find(&post, postID)

	return post
}
