package main

import (
	"opiana_code_test/config"
	"opiana_code_test/controller"
	"opiana_code_test/middleware"
	"opiana_code_test/repository"
	"opiana_code_test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	postRepository repository.PostRepository = repository.NewPostRepository(db)
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	serve := gin.Default()

	authRoutes := serve.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := serve.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
	}

	postRoutes := serve.Group("api/post", middleware.AuthorizeJWT(jwtService))
	{
		postRoutes.GET("/", postController.All)
		postRoutes.POST("/", postController.Insert)
		postRoutes.GET("/:id", postController.FindByID)
		postRoutes.PUT("/:id", postController.Update)
		postRoutes.DELETE("/:id", postController.Delete)

	}

	serve.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
