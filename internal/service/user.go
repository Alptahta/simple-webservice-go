package service

import "github.com/Alptahta/simple-webservice-go/internal"

type UserRepository interface {
	Create(name string) error
	Find(id uint) (internal.User, error)
}

//User struct
type User struct {
	repo UserRepository
}

//Create method creates new user record
func (u User) Create(name string) error {
	err := u.repo.Create(name)
	if err != nil {
		return err
	}
	return nil
}

func (u User) Find(id uint) (internal.User, error) {
	user, err := u.repo.Find(id)
	if err != nil {
		return internal.User{}, err
	}
	return user, nil
}
