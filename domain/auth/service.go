package auth

import "github.com/rakhmatullahyoga/tigade"

type userRepository interface {
	FindByEmail(email string)
}

type service struct {
	ur userRepository
}

// NewService construct the auth service
func NewService(ur userRepository) tigade.AuthService {
	return service{ur}
}

func (s service) Login(email, password string) (tigade.Token, error) {
	return "", nil
}
