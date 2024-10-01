package enum

// Role type for defining user roles
type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
	Guest Role = "guest"
	// Add new roles here in the future
)
