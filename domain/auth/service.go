package auth

import "github.com/rakhmatullahyoga/tigade"

type UserRepository interface {
	FindByEmail(email string)
}

type service struct {
	ur UserRepository
}

func NewService(ur UserRepository) tigade.AuthService {
	return service{ur}
}

func (s service) Login(email, password string) (tigade.Token, error) {
	return "", nil
}
