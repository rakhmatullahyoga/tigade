package account

import "github.com/rakhmatullahyoga/tigade"

type UserRepository interface {
	Insert(user tigade.User) (uint64, error)
	FindByID(id uint64) (*tigade.User, error)
	Update(id uint64, user tigade.User) error
	FindByToken(token tigade.Token) (*tigade.User, error)
	ActivateByID(id uint64) error
}

type UserFactory interface {
	CreateNewUser(email, password, name string) tigade.User
}

type TokenFactory interface {
	CreateActivationToken(ID uint64) tigade.Token
	CreateAuthToken(user tigade.User) tigade.Token
}

type service struct {
	tf TokenFactory
	uf UserFactory
	ur UserRepository
}

func NewService(tf TokenFactory, uf UserFactory, ur UserRepository) tigade.AccountService {
	return service{tf, uf, ur}
}

func (s service) Register(email, password, name string) error {
	user := s.uf.CreateNewUser(email, password, name)
	id, err := s.ur.Insert(user)
	if err != nil {
		return err
	}
	s.tf.CreateActivationToken(id)
	return err
}

func (s service) ActivateUser(id uint64) error {
	user, err := s.ur.FindByID(id)
	if err != nil {
		return err
	}
	err = s.ur.ActivateByID(user.ID)
	return err
}
