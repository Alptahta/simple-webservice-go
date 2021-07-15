package internal

import "errors"

type User struct {
	ID   uint
	Name string
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("Name Cannot be empty")
	}
	if len(u.Name) >= 10 {
		return errors.New("Lenght of the Name cannot be more than 10")
	}
	return nil
}
