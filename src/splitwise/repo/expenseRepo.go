package repo

import (
	"container/list"
	"fmt"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"math"
)

type ExpenseRepository struct {
	userMap      map[string]model.User
	balanceSheet map[string]map[string]float64
}

func NewExpenseRepository() ExpenseRepository {
	return ExpenseRepository{
		userMap:      make(map[string]model.User, 0),
		balanceSheet: make(map[string]map[string]float64, 0),
	}
}

func (e *ExpenseRepository) AddUser(user model.User) {
	e.userMap[user.GetUserName()] = user
	e.balanceSheet[user.GetUserName()] = make(map[string]float64, 0)
}

func (e *ExpenseRepository) GetUser(username string) model.User {
	return e.userMap[username]
}

func (e *ExpenseRepository) AddExpense(
	paidBy string,
	users map[model.User]float64) {

	for user, amt := range users {
		paidTo := user.GetUserName()
		balances := e.balanceSheet[paidBy]

		if _, ok := balances[paidTo]; !ok {
			balances[paidTo] = 0.0
		}

		balances[paidTo] = balances[paidTo] + amt

		balances = e.balanceSheet[paidTo]
		if _, ok := balances[paidBy]; !ok {
			balances[paidBy] = 0.0
		}

		balances[paidBy] = balances[paidBy] - amt
	}
}

func (e *ExpenseRepository) GetBalance(username string) *list.List {
	balances := list.New()

	userBalance := e.balanceSheet[username]
	if len(userBalance) != 0 {
		for user, bal := range userBalance {
			balances.PushBack(e.checkSign(username, user, bal))
		}
	}

	return balances
}

func (e *ExpenseRepository) GetBalances() *list.List {
	balances := list.New()

	allBalances := e.balanceSheet
	for user1, val := range allBalances {
		if len(val) > 0 {
			for user2, balance := range val {
				balances.PushBack(e.checkSign(user1, user2, balance))
			}
		}
	}

	return balances
}

func (e *ExpenseRepository) checkSign(user1, user2 string, amount float64) string {
	user1Name := e.userMap[user1].GetUserName()
	user2Name := e.userMap[user2].GetUserName()

	if amount < 0 {
		return user1Name + " owes " + user2Name + ": " + fmt.Sprintf("%f", math.Abs(amount))
	} else if amount > 0 {
		return user2Name + " owes " + user1Name + ": " + fmt.Sprintf("%f", math.Abs(amount))
	}

	return ""
}
