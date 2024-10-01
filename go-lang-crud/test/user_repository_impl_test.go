package test

// func setupTestDB() (*gorm.DB, error) {
// 	// Configure GORM logger
// 	newLogger := logger.New(
// 		log.New(log.Writer(), "\r\n", log.LstdFlags), // IO writer (stdout)
// 		logger.Config{
// 			SlowThreshold: time.Second, // Log SQL queries that take longer than this
// 			LogLevel:      logger.Info, // LogLevel: Silent, Error, Warn, Info
// 			Colorful:      true,        // Enable/Disable color
// 		},
// 	)
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
// 		Logger: newLogger,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&models.User{})
// 	return db, nil
// }

// func TestUserRepositoryImpl_Create(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewUserRepository(db)
// 	user := &models.User{Name: "John Doe"}

// 	// Test Create method
// 	createdUser, err := repo.Create(user)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "John Doe", createdUser.Name)
// 	assert.NotZero(t, createdUser.ID)
// }

// func Test_FindAll(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewUserRepository(db)
// 	user := &models.User{Name: "John", Email: "john@gmail.com"}
// 	expectedUsers := []*models.User{user}

// 	repo.Create(user)

// 	result, err := repo.FindAll()

// 	log.Println("expectedUsers[0].Name : ", expectedUsers[0].Name)
// 	log.Println("result[0].Name : ", result[0].Name)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUsers[0].Name, result[0].Name)
// 	assert.NotZero(t, result)

// }

// func TestMultipleUpdateSaveTransaction_Error(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	// Create a new user
// 	user := &models.User{Name: "Steve", Email: "steve@email.com"}
// 	db.Create(user)

// 	repo := repository.UserRepositoryImpl{DB: db}

// 	// Attempt the transaction which is expected to fail
// 	_, err = repo.MultipleUpdateSaveTransaction(user)

// 	// Assertions
// 	assert.Error(t, err)

// 	// Fetch the user again to ensure no updates were applied
// 	var result models.User
// 	db.First(&result, user.ID)
// 	assert.Equal(t, "Steve", result.Name)            // Should remain unchanged
// 	assert.Equal(t, "steve@email.com", result.Email) // Should remain unchanged
// }
