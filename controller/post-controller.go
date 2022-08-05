package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"opiana_code_test/dto"
	"opiana_code_test/entity"
	"opiana_code_test/helper"
	"opiana_code_test/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type postController struct {
	postService service.PostService
	jwtService  service.JWTService
}

func NewPostController(postService service.PostService, jwtService service.JWTService) PostController {
	return &postController{
		postService: postService,
		jwtService:  jwtService,
	}
}

func (c *postController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims, ok := aToken.Claims.(jwt.MapClaims)
	if !ok {
		panic("Failed to claim token")
	}
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *postController) All(context *gin.Context) {
	var post []entity.Post = c.postService.All()
	res := helper.BuildResponse(true, "OK", post)
	context.JSON(http.StatusOK, res)
}

func (c *postController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	var post entity.Post = c.postService.FindById(id)

	if (post == entity.Post{}) {
		res := helper.BuildErrorResponse("Data not found", "No database given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", post)
		context.JSON(http.StatusOK, res)
	}

}

func (c *postController) Insert(context *gin.Context) {
	var postCreateDTO dto.PostCreateDTO

	errDTO := context.ShouldBind(&postCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {

		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			postCreateDTO.UserID = convertedUserID
		}

		result, err := c.postService.Insert(postCreateDTO)

		if err != nil {
			resp := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *postController) Update(context *gin.Context) {
	var postUpdateDTO dto.PostUpdateDTO
	errDTO := context.ShouldBind(&postUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	postUpdateDTO.ID = id

	if c.postService.IsAllowedToEdit(userID, postUpdateDTO.ID) {
		userid, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			postUpdateDTO.UserID = userid
		}

		result, err := c.postService.Update(postUpdateDTO)
		if err != nil {
			resp := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}

}

func (c *postController) Delete(context *gin.Context) {
	var post entity.Post
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	post.ID = id

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if c.postService.IsAllowedToEdit(userID, uint64(post.ID)) {
		c.postService.Delete(post)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}

}
