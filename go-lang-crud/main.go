package main

import (
	"golang-crud/controllers"
	"golang-crud/enum"
	"golang-crud/initializers"
	"golang-crud/middlewares"
	"golang-crud/repository"
	"golang-crud/service"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToGoogle()
}

func main() {

	// Set up repository and services
	repo := repository.NewCompanyRepository(initializers.DB) // Assume this is correctly implemented
	companyService := service.NewCompanyServiceImpl(repo)
	companyController := controllers.NewCompanyController(companyService)

	postRepo := repository.NewPostRepository(initializers.DB)
	postService := service.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	userRepo := repository.NewUserRepository(initializers.DB)
	userService := service.NewUserServiceImpl(userRepo) // Returns an implementation of UserService interface
	userController := controllers.NewUserController(userService)
	authCon := controllers.NewGoAuthController(userService)

	// Create a Gin router
	r := gin.Default()

	// Define the home route
	r.GET("/", authCon.HandleHome)

	// Google Login route
	r.GET("/login", authCon.SignInWithProvider)

	// Callback route
	r.GET("/callback", authCon.CallbackHandler)

	// Google Login route
	//r.GET("/", userController.HandleHome)

	// r.GET("/login", userController.HandleGoogleLogin)

	// // Callback route
	// r.GET("/callback", userController.HandleGoogleCallback)
	// r.GET("/success", userController.Success)

	//Company API's
	r.POST("/company", companyController.CreateCompany)
	r.GET("/getAllCompanies", companyController.GetAllCompanies)
	r.DELETE("/deleteCompany/:id", companyController.DeleteCompany)

	//Post API's
	r.POST("/post", postController.CreatePost)
	r.GET("/getAllPosts/:id", postController.GetPosts)
	r.GET("/getPost/:id", postController.GetPostById)

	//Users API's
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/", middlewares.RoleAuthorization(enum.Admin), userController.CreateUser)               // Create user
		userRoutes.GET("/", middlewares.RoleAuthorization(enum.Admin), userController.GetUsers)                  // Get all users
		userRoutes.GET("/:id", middlewares.RoleAuthorization(enum.Admin, enum.User), userController.GetUserById) // Get user by ID
		userRoutes.PUT("/:id", middlewares.RoleAuthorization(enum.User), userController.UpdateUserDetails)       // Update user details
		userRoutes.DELETE("/:id", middlewares.RoleAuthorization(enum.Admin), userController.DeleteUser)          // Delete user
		userRoutes.GET("/paginated", middlewares.RoleAuthorization(enum.Admin), userController.PaginateUsers)
		r.POST("/login", userController.LoginUser)
		// Paginated users
	}

	r.Run(":8081")
}
