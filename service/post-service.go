package service

import (
	"fmt"
	"log"

	"opiana_code_test/dto"
	"opiana_code_test/entity"
	"opiana_code_test/repository"

	"github.com/mashingan/smapping"
)

// PostService
type PostService interface {
	Insert(b dto.PostCreateDTO) entity.Post
	Update(b dto.PostUpdateDTO) entity.Post
	Delete(b entity.Post)
	All() []entity.Post
	FindById(postID uint64) entity.Post
	IsAllowedToEdit(userID string, postID uint64) bool
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (service *postService) Insert(b dto.PostCreateDTO) entity.Post {

	post := entity.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := service.postRepository.InsertPost(post)
	return res
}

func (service *postService) Update(b dto.PostUpdateDTO) entity.Post {
	// post := entity.Post{
	// 	ID:          b.ID,
	// 	Title:       b.Title,
	// 	Description: b.Description,
	// 	UserID:      b.UserID,
	// 	PostTypeID:  b.PostTypeID,
	// }
	fmt.Println("b.ID=======", b.ID)
	post := entity.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := service.postRepository.UpdatePost(post)
	// res := entity.Post{}
	return res
}

func (service *postService) Delete(b entity.Post) {
	service.postRepository.DeletePost(b)
}

func (service *postService) All() []entity.Post {
	return service.postRepository.AllPost()
}

func (service *postService) FindById(postID uint64) entity.Post {
	return service.postRepository.FindPostByID(postID)
}

func (service *postService) IsAllowedToEdit(userID string, postID uint64) bool {
	b := service.postRepository.FindPostByID(postID)

	id := fmt.Sprintf("%v", b.UserID)

	return userID == id
}
