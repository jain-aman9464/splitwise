package service

import (
	"fmt"
	"github.com/tokopedia/test/splitwise/src/splitwise/repo"
)

type SplitService struct {
	expenserepo repo.ExpenseRepository
}

func NewSplitService(repository repo.ExpenseRepository) SplitService {
	return SplitService{expenserepo: repository}
}

func (s *SplitService) ShowBalance(username string) {
	fmt.Println("")
	balances := s.expenserepo.GetBalance(username)
	if balances.Len() == 0 {
		fmt.Println("No Balances")
	} else {
		for element := balances.Front(); element != nil; element = element.Next() {
			balance := element.Value.(string)
			fmt.Println(balance)
		}
	}
}

func (s *SplitService) ShowBalances() {
	balances := s.expenserepo.GetBalances()
	fmt.Println("")
	if balances.Len() == 0 {
		fmt.Println("No Balances")
	} else {
		for element := balances.Front(); element != nil; element = element.Next() {
			balance := element.Value.(string)
			fmt.Println(balance)
		}
	}
}
