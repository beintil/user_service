package client

// AuthData Authorization data of the user
type AuthData struct {
	// Email Address of the user
	Email string `json:"email" validate:"required"`
	// Password of the user
	Password string `json:"password" validate:"required"`
}
