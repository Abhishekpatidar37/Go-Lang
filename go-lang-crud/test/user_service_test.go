package test

// // setup initializes and returns the mock repository and user service.
// func setup() (*repository.MockUserRepository, service.UserService) {
// 	repo := new(repository.MockUserRepository)
// 	svc := service.NewUserServiceImpl(repo)
// 	return repo, svc
// }

// func TestUserServiceImpl(t *testing.T) {

// 	t.Run("CreateUser Success", func(t *testing.T) {
// 		repo, service := setup()

// 		user := &models.User{ID: 1, Name: "John Doe"}
// 		repo.On("Create", user).Return(user, nil)

// 		result, err := service.CreateUser(user)

// 		assert.NoError(t, err)
// 		assert.Equal(t, user, result)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("CreateUser Failure", func(t *testing.T) {
// 		repo, service := setup()
// 		user := &models.User{ID: 1, Name: "John Doe"}
// 		expectedError := errors.New("some error")

// 		repo.On("Create", user).Return((*models.User)(nil), expectedError)

// 		createdUser, err := service.CreateUser(user)

// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, createdUser)

// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("GetAllUsers Success", func(t *testing.T) {
// 		repo, service := setup()
// 		users := []models.User{{ID: 1, Name: "John Doe"}}
// 		repo.On("FindAll").Return(users, nil)

// 		result, err := service.GetAllUsers()

// 		assert.NoError(t, err)
// 		assert.Equal(t, users, result)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("GetAllUsers Failure", func(t *testing.T) {
// 		repo, service := setup()
// 		expectedError := errors.New("some error")
// 		repo.On("FindAll").Return(([]models.User)(nil), expectedError)

// 		allUsers, err := service.GetAllUsers()

// 		assert.Equal(t, fmt.Errorf("failed to retrieve all users: %w", expectedError), err)
// 		assert.Nil(t, allUsers)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("GetUserById Success", func(t *testing.T) {
// 		repo, service := setup()
// 		user := &models.User{ID: 1, Name: "John Doe"}
// 		repo.On("FindById", "1").Return(user, nil)

// 		result, err := service.GetUserById("1")

// 		assert.NoError(t, err)
// 		assert.Equal(t, user, result)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("GetUserById Failure", func(t *testing.T) {
// 		repo, service := setup()
// 		expectedError := custom_error.ErrUserNotFound
// 		id := "1"
// 		repo.On("FindById", id).Return((*models.User)(nil), expectedError)

// 		foundUser, err := service.GetUserById(id)

// 		assert.Equal(t, fmt.Errorf("user with ID %s not found: %w", id, expectedError), err)
// 		assert.Nil(t, foundUser)

// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("UpdateUserDetails Success", func(t *testing.T) {
// 		repo, service := setup()
// 		user := &models.User{ID: 1, Name: "John Doe"}
// 		data := map[string]interface{}{"name": "Jane Doe"}
// 		repo.On("Update", user, data).Return(nil)

// 		err := service.UpdateUserDetails(user, data)

// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("UpdateUserDetails Failure", func(t *testing.T) {
// 		repo, service := setup()
// 		user := &models.User{ID: 1, Name: "John Doe"}
// 		data := map[string]interface{}{"Name": "Will Smith"}
// 		expectedError := errors.New("Error updating user details!")

// 		repo.On("Update", user, data).Return(expectedError)

// 		err := service.UpdateUserDetails(user, data)

// 		assert.Error(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("DeleteUser Success", func(t *testing.T) {
// 		repo, service := setup()
// 		repo.On("Delete", "1").Return(nil)

// 		err := service.DeleteUser("1")

// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("PaginateUsers Success", func(t *testing.T) {
// 		repo, service := setup()
// 		users := []models.User{{ID: 1, Name: "John Doe"}}
// 		repo.On("Paginate", 0, 10).Return(users, nil)

// 		result, err := service.PaginateUsers(1, 10)

// 		assert.NoError(t, err)
// 		assert.Equal(t, users, result)
// 		repo.AssertExpectations(t)
// 	})
// }
