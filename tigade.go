package tigade

// User entity
type User struct {
	ID          uint64
	Email       string
	Password    string
	DisplayName string
	Active      bool
}

type Token string

// AccountService contract
type AccountService interface {
	Register(email, password, name string) error
	ActivateUser(id uint64) error
}

// AuthService contract
type AuthService interface {
	Login(email, password string) (Token, error)
}
