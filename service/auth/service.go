package auth

import (
	"errors"
	"github.com/rakhmatullahyoga/tigade"
)

type userRepository interface {
	FindByEmail(email string) (*tigade.User, error)
}

type tokenFactory interface {
	GenerateJWTToken(user tigade.User) (tigade.Token, error)
}

type service struct {
	ur userRepository
	tf tokenFactory
}

// NewService construct the auth service
func NewService(ur userRepository, tf tokenFactory) tigade.AuthService {
	return service{ur, tf}
}

func (s service) Login(email, password string) (tigade.Token, error) {
	user, err := s.ur.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	} else {
		return s.tf.GenerateJWTToken(*user)
	}
}
