package service

import "dependency-injection/repo"

type UserService interface {
	//Dummy call
	DummyFunc() string
}

type UserServiceImpl struct {
	db repo.Database
}

func NewUserService(db repo.Database) UserService {
	return UserServiceImpl{db: db}
}

func (u UserServiceImpl) DummyFunc() string {
	return "dummy function called"
}
