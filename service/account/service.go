package account

import "github.com/rakhmatullahyoga/tigade"

type userRepository interface {
	Insert(user tigade.User) (uint64, error)
}

type userFactory interface {
	CreateNewUser(email, password, name string) tigade.User
}

type service struct {
	uf userFactory
	ur userRepository
}

// NewService construct the account service
func NewService(uf userFactory, ur userRepository) tigade.AccountService {
	return service{uf, ur}
}

func (s service) Register(email, password, name string) error {
	user := s.uf.CreateNewUser(email, password, name)
	_, err := s.ur.Insert(user)
	return err
}
