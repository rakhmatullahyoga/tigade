package tigade

// User entity
type User struct {
	Email       string
	Password    string
	DisplayName string
}

type Token string

// AccountService contract
type AccountService interface {
	Register(email string) error
	ActivateUser(regToken Token) error
	UpdateProfile(name string) error
}

// AuthService contract
type AuthService interface {
	Login(email, password string) (Token, error)
	Logout(token Token) error
}
