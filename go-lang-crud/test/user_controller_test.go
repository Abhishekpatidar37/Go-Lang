package test

// // Setup the gin engine and user controller
// func setupTestController() (*gin.Engine, *controllers.UserController, *service.MockUserService) {
// 	mockUserService := new(service.MockUserService)
// 	userController := controllers.NewUserController(mockUserService)
// 	router := gin.Default()

// 	router.POST("/user", userController.CreateUser)
// 	router.GET("/getUsers", userController.GetUsers)
// 	router.GET("/getUserById/:id", userController.GetUserById)
// 	router.PUT("/updateUser/:id", userController.UpdateUserDetails)
// 	router.DELETE("/deleteUser/:id", userController.DeleteUser)
// 	router.GET("/paginatedUser", userController.PaginateUsers)

// 	return router, userController, mockUserService
// }

// // Test cases

// func TestCreateUser_Success(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router, _, mockUserService := setupTestController()

// 	mockUser := &models.User{Name: "John Doe", Email: "john@example.com"}
// 	mockUserService.On("CreateUser", mockUser).Return(mockUser, nil)

// 	// Prepare the request payload
// 	userPayload, _ := json.Marshal(mockUser)
// 	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userPayload))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a test HTTP response recorder
// 	w := httptest.NewRecorder()

// 	// Call the handler
// 	router.ServeHTTP(w, req)

// 	// Assert the response
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	mockUserService.AssertExpectations(t)
// }

// func TestGetUsers_Success(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router, _, mockUserService := setupTestController()

// 	mockUsers := []models.User{{Name: "John Doe", Email: "john@example.com"}, {Name: "Jane Doe", Email: "jane@example.com"}}
// 	mockUserService.On("GetAllUsers").Return(mockUsers, nil)

// 	req, _ := http.NewRequest(http.MethodGet, "/getUsers", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockUserService.AssertExpectations(t)
// }

// func TestGetUserById_NotFound(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router, _, mockUserService := setupTestController()

// 	// Mock an error case where the user is not found
// 	mockUserService.On("GetUserById", "123").Return((*models.User)(nil), errors.New("User not found")) // Correctly return typed nil

// 	req, _ := http.NewRequest(http.MethodGet, "/getUserById/123", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)
// 	mockUserService.AssertExpectations(t)
// }

// func TestUpdateUserDetails_Success(t *testing.T) {
// 	router, _, mockUserService := setupTestController()

// 	// Mock user and data
// 	mockUser := &models.User{Name: "John Doe", Email: "john@example.com"}
// 	updateData := map[string]interface{}{"name": "John Updated", "email": "johnupdated@example.com"}

// 	mockUserService.On("GetUserById", "123").Return(mockUser, nil)
// 	mockUserService.On("UpdateUserDetails", mockUser, updateData).Return(nil)

// 	userPayload := `{"name": "John Updated", "email": "johnupdated@example.com"}`
// 	req, _ := http.NewRequest(http.MethodPut, "/updateUser/123", bytes.NewBuffer([]byte(userPayload)))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockUserService.AssertExpectations(t)
// }

// func TestDeleteUser_Success(t *testing.T) {
// 	router, _, mockUserService := setupTestController()

// 	// Mock success response
// 	mockUserService.On("DeleteUser", "123").Return(nil)

// 	req, _ := http.NewRequest(http.MethodDelete, "/deleteUser/123", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockUserService.AssertExpectations(t)
// }
