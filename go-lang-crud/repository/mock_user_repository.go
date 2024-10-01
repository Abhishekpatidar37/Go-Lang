package repository

import (
	"golang-crud/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FindById(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *models.User, data map[string]interface{}) error {
	args := m.Called(user, data)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) Paginate(offset, pageSize int) ([]models.User, error) {
	args := m.Called(offset, pageSize)
	return args.Get(0).([]models.User), args.Error(1)
}

func (r *MockUserRepository) MultipleUpdateSaveTransaction(user *models.User) (*models.User, error) {
	return user, nil
}
