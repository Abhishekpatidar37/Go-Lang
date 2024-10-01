package custom_error

import "errors"

// ErrUserNotFound represents an error when a user is not found.
var ErrUserNotFound = errors.New("user not found")