// service/post_service.go
package service

import (
	"golang-crud/models"
	"golang-crud/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) GetPostsByUserId(userId string) ([]models.Post, error) {
	return s.repo.FindByUserId(userId)
}

func (s *PostService) GetPostById(id string) (*models.Post, error) {
	return s.repo.FindById(id)
}
