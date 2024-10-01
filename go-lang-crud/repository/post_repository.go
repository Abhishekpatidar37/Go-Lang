// repository/post_repository.go
package repository

import (
	"golang-crud/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) FindByUserId(userId string) ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindById(id string) (*models.Post, error) {
	var post models.Post
	err := r.DB.Preload("User").First(&post, "id = ?", id).Error
	return &post, err
}
