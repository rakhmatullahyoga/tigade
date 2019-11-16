package tigade

// Service model/schema
type User struct {
	ID          uint64
	Email       string
	Password    string
	DisplayName string
	Active      bool
}

type Token string

// Service contract
type AccountService interface {
	Register(email, password, name string) error
	ActivateUser(id uint64) error
}

type AuthService interface {
	Login(email, password string) (Token, error)
}
