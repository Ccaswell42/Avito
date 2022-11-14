package models

import "errors"

//type Users struct {
//	Users map[int]int
//}

type Users map[int]int

func (u *Users) Replenishment(id, amount int) error {

	s := *u
	_, ok := s[id]
	if !ok {
		return errors.New("no such user")
	}

	s[id] += amount
	return nil
}
