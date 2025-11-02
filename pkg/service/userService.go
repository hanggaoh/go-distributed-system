package service

import "go-distributed-system/pkg/store"

type UserService struct {
	Store store.UserStore
}

func (s *UserService) GetUserName(userId int) (string, error) {
	name, error := s.Store.GetUser(userId)
	if error != nil {
		return "", error
	}
	return name, nil
}
