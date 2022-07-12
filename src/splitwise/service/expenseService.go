package service

import (
	"container/list"
	"fmt"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/expense"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/split"
	"github.com/tokopedia/test/splitwise/src/splitwise/repo"
	"math"
)

type ExpenseService struct {
	expenseRepo repo.ExpenseRepository
}

func NewExpenseService(repository repo.ExpenseRepository) ExpenseService {
	return ExpenseService{
		expenseRepo: repository,
	}
}

func (e *ExpenseService) CreateExpense(expenseType string,
	amount float64,
	expensePaidBy string,
	splits list.List,
	title string) {

	var expenserInterface expense.Expenser
	usersAmountMap := make(map[model.User]float64, 0)

	expensePaidByUser := e.expenseRepo.GetUser(expensePaidBy)
	switch expenseType {
	case "EXACT_SPLIT":
		for element := splits.Front(); element != nil; element = element.Next() {
			split := element.Value.(split.ExactSplit)
			usersAmountMap[split.GetUser()] = split.GetAmount()
		}

		expenserInterface = expense.NewExactExpense(amount, expensePaidByUser, splits, title)

	case "PERCENT_SPLIT":
		for element := splits.Front(); element != nil; element = element.Next() {
			split := element.Value.(split.PercentSplit)
			split.SetAmount((amount * split.GetPercent()) / 100.0)
			usersAmountMap[split.GetUser()] = split.GetAmount()
		}

		expenserInterface = expense.NewPercentExpense(amount, expensePaidByUser, splits, title)

	case "EQUAL_SPLIT":
		totalSplits := splits.Len()
		splitAmount := math.Round((amount*100.0)/float64(totalSplits)) / 100.0

		for element := splits.Front(); element != nil; element = element.Next() {
			split := element.Value.(split.EqualSplit)
			split.SetAmount(splitAmount)
			usersAmountMap[split.GetUser()] = split.GetAmount()
		}

		expenserInterface = expense.NewEqualExpense(amount, expensePaidByUser, splits, title)
	}


	isValidate := expenserInterface.Validate()
	if !isValidate {
		fmt.Println("INVALID Distribution. Please check your input again")
		return
	}

	e.expenseRepo.AddExpense(expensePaidBy, usersAmountMap)
	fmt.Println("\nExpense Added !!")
}
