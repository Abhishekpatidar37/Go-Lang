// repository/user_repository_interface.go
package repository

import "golang-crud/models"

// UserRepository defines the methods for user repository operations.
type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindById(id string) (*models.User, error)
	Update(user *models.User, data map[string]interface{}) error
	Delete(id string) error
	Paginate(offset, limit int) ([]models.User, error)
	MultipleUpdateSaveTransaction(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}
