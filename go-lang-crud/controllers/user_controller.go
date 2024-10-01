// controllers/user_controller.go
package controllers

import (
	"golang-crud/models"
	"golang-crud/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser - Calls the CreateUser method in the service
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": createdUser})
}

// GetUsers - Calls GetAllUsers method in the service
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserById - Calls the GetUserById method in the service
func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.userService.GetUserById(id)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUserDetails - Calls the UpdateUserDetails method in the service
func (uc *UserController) UpdateUserDetails(c *gin.Context) {
	var userRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	data := map[string]interface{}{
		"name":  userRequest.Name,
		"email": userRequest.Email,
	}

	if err := uc.userService.UpdateUserDetails(user, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// DeleteUser - Calls the DeleteUser method in the service
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uc.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// PaginateUsers - Calls the PaginateUsers method in the service
func (uc *UserController) PaginateUsers(c *gin.Context) {
	var requestBody struct {
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Page <= 0 {
		requestBody.Page = 1
	}

	if requestBody.PageSize <= 0 {
		requestBody.PageSize = 10
	}

	users, err := uc.userService.PaginateUsers(requestBody.Page, requestBody.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var userLogin struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate the user using the service layer
	token, err := uc.userService.AuthenticateUser(userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Respond with the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// var googleOauthConfig = &oauth2.Config{
// clientID := os.Getenv("CLIENT_ID")
// clientSecret := os.Getenv("CLIENT_SECRET")
// RedirectURL := os.Getenv("Callback_URL")
// 	Scopes: []string{
// 		"https://www.googleapis.com/auth/userinfo.profile",
// 		"https://www.googleapis.com/auth/userinfo.email",
// 	},
// 	Endpoint: google.Endpoint,
// }

// var oauthStateString = "sjdsjcjbcbxbcmcbnxc"

// // Google Login Handler
// func (uc *UserController) HandleGoogleLogin(c *gin.Context) {
// 	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// }

// // Google OAuth Callback Handler
// func (uc *UserController) HandleGoogleCallback(c *gin.Context) {
// 	// Verify the state to protect against CSRF attacks
// 	if c.Query("state") != oauthStateString {
// 		fmt.Println("Invalid OAuth state")
// 		c.Redirect(http.StatusTemporaryRedirect, "/")
// 		return
// 	}

// 	// Exchange the authorization code for an access token
// 	code := c.Query("code")
// 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		fmt.Printf("Code exchange failed: %s\n", err.Error())
// 		c.Redirect(http.StatusTemporaryRedirect, "/")
// 		return
// 	}

// 	// Use the access token to fetch the user's information
// 	client := googleOauthConfig.Client(context.Background(), token)
// 	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		fmt.Printf("Error getting user info: %s\n", err.Error())
// 		c.Redirect(http.StatusTemporaryRedirect, "/")
// 		return
// 	}
// 	defer userInfoResp.Body.Close()

// 	// Display user info
// 	userInfo, err := io.ReadAll(userInfoResp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading user info:", err)
// 		c.String(http.StatusInternalServerError, "Internal Server Error")
// 		return
// 	}

// 	var googleUser map[string]interface{}
// 	err = json.Unmarshal(userInfo, &googleUser)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling user info:", err)
// 		c.String(http.StatusInternalServerError, "Internal Server Error")
// 		return
// 	}

// 	email := googleUser["email"].(string)
// 	//name := googleUser["name"].(string)

// 	user, err := uc.userService.GetUserByEmail(email)
// 	if err != nil {
// 		log.Println("Error while fetching user", err)

// 		// Check if the error is because the user was not found
// 		if err == custom_error.ErrUserNotFound { // Assuming this is a specific error from your service layer
// 			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
// 		} else {
// 			// If it's a different kind of error, return a 500 status code
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user information"})
// 		}
// 		return
// 	}

// 	jwtToken, err := security.GenerateJWT(user)
// 	if err != nil {
// 		log.Println("Error generating JWT token", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	// Return the user info and JWT token in the response
// 	c.JSON(http.StatusOK, gin.H{
// 		"user":  user,
// 		"token": jwtToken,
// 	})
// }

func (uc *UserController) Success(c *gin.Context) {
	html := `<html><body>Login Success</body></html>`
	c.Data(200, "text/html; charset=utf-8", []byte(html))
}
