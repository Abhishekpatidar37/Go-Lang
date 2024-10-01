package service

import (
	"errors"
	"fmt"
	"golang-crud/custom_error"
	"golang-crud/models"
	"golang-crud/repository"
	"golang-crud/security"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.UserRepository // Keep the concrete type
}

func NewUserServiceImpl(repo repository.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set the hashed password back to the user model
	user.Password = string(hashedPassword)
	result, err := s.repo.Create(user)
	if err != nil {
		return nil, err // Ensure this line exists
	}
	return result, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all users: %w", err)
	}
	return users, nil
}

func (s *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			return nil, fmt.Errorf("user with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to retrieve user with ID %s: %w", id, err)
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUserDetails(user *models.User, data map[string]interface{}) error {
	if password, ok := data["password"]; ok {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		// Set the hashed password in the data map
		data["password"] = string(hashedPassword)
	}

	err := s.repo.Update(user, data)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			return fmt.Errorf("user with ID %v not found: %w", user.ID, err)
		}
		return fmt.Errorf("failed to update user details: %w", err)
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			return fmt.Errorf("user with ID %s not found: %w", id, err)
		}
		return fmt.Errorf("failed to delete user with ID %s: %w", id, err)
	}
	return nil
}

func (s *UserServiceImpl) PaginateUsers(page, pageSize int) ([]models.User, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("invalid page or page size")
	}

	offset := (page - 1) * pageSize
	users, err := s.repo.Paginate(offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to paginate users: %w", err)
	}
	return users, nil
}

func (s *UserServiceImpl) AuthenticateUser(email, password string) (string, error) {
	// Fetch the user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Println("Error when fetching user ", err)
		return "", errors.New("invalid email or password")
	}
	log.Println("User ", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Error when matching passwrod ", err)
		return "", errors.New("invalid password")
	}

	token, err := security.GenerateJWT(user)
	log.Println("generated token ", token, err)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	log.Println("user error in service ", err, email)
	log.Println("user  in service ", user)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			return nil, fmt.Errorf("user with Email %s not found: %w", email, err)
		}
		return nil, fmt.Errorf("failed to retrieve user with email %s: %w", email, err)
	}
	return user, nil
}
