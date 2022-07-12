package service

import (
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/repo"
)

type UserService struct {
	expenseRepo repo.ExpenseRepository
}

func NewUserService(repository repo.ExpenseRepository) UserService {
	return UserService{
		expenseRepo: repository,
	}
}

func (u *UserService) AddUser(user model.User) {
	u.expenseRepo.AddUser(user)
}

func (u UserService) GetUser(username string) model.User {
	return u.expenseRepo.GetUser(username)
}
