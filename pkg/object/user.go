package object

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Email    string
	Role     string
	GitHub   GitHub
}

type GitHub struct {
	AccessToken string
	Username    string
}
