// controllers/user_controller.go
package controllers

import (
	"golang-crud/security"
	"golang-crud/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type GoAuthController struct {
	userService service.UserService
}

func NewGoAuthController(userService service.UserService) *GoAuthController {
	return &GoAuthController{userService: userService}
}

func (uc *GoAuthController) HandleHome(c *gin.Context) {
	html := `<html><body><a href="/login">Google Login</a></body></html>`
	c.Data(200, "text/html; charset=utf-8", []byte(html))
}

func (uc *GoAuthController) SignInWithProvider(c *gin.Context) {
	provider := "google"
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// CallbackHandler processes the authentication response from Google.
func (uc *GoAuthController) CallbackHandler(c *gin.Context) {
	provider := "google"
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Access the user repository to manage users
	// userRepository := repository.NewUserRepository(initializers.DB)
	userData, err := uc.userService.GetUserByEmail(user.Email)

	if err != nil {
		// User does not exist, handle accordingly
		c.JSON(http.StatusBadRequest, gin.H{"error": "User doesn't exist with email: " + user.Email})
		return
	}

	// Generate JWT token for the user
	tokenString, err := security.GenerateJWT(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error generating token"})
		return
	}

	// // Optionally, return the token to the frontend or store it in session/local storage
	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"user":    userData,
		"token":   tokenString,
	})
}
