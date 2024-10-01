// repository/user_repository.go
package repository

import (
	"golang-crud/models"
	"log"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepositoryImpl.
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

// Implement the UserRepository interface
func (r *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	err := r.DB.Create(user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Posts").First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) Update(user *models.User, data map[string]interface{}) error {
	return r.DB.Model(user).Updates(data).Error
}

func (r *UserRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepositoryImpl) Paginate(offset, limit int) ([]models.User, error) {
	var users []models.User
	err := r.DB.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) MultipleUpdateSaveTransaction(user *models.User) (*models.User, error) {
	var result models.User

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Update("name", "John Doe").Error; err != nil {
			return err
		}
		log.Println("User name updated : ", user)
		if err := tx.Model(user).Update("email", "doe.john@email.com").Error; err != nil {
			return err
		}
		log.Println("User email updated : ", user)
		if err := tx.Model(user).Update("email", nil).Error; err != nil {
			return err
		}
		log.Println("User email updated 2nd time : ", user)
		if err := tx.First(&result, user.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err // Handle the transaction error here
	}

	return &result, nil // Return the updated user if successful
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	log.Println("error in repo ", err)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
