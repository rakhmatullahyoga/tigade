package tigade

// Service model/schema
type User struct {
	Email       string
	Password    string
	DisplayName string
}

type Token string

// Service contract
type AccountService interface {
	Register(email string) error
	ActivateUser(regToken Token) error
	UpdateProfile(name string) error
}

type AuthService interface {
	Login(email, password string) (Token, error)
	Logout(token Token) error
}
