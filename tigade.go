package tigade

// User entity
type User struct {
	ID           uint64
	Email        string
	PasswordHash string
	DisplayName  string
	Active       bool
}

type Token string

// AccountService contract
type AccountService interface {
	Register(email, password, name string) error
}

// AuthService contract
type AuthService interface {
	Login(email, password string) (Token, error)
}
