// controllers/post_controller.go
package controllers

import (
	"golang-crud/models"
	"golang-crud/service"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := pc.postService.CreatePost(&post); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(201, gin.H{"message": "Post created successfully", "post": post})
}

func (pc *PostController) GetPosts(c *gin.Context) {
	uid := c.Param("id")
	posts, err := pc.postService.GetPostsByUserId(uid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (pc *PostController) GetPostById(c *gin.Context) {
	id := c.Param("id")
	post, err := pc.postService.GetPostById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}
