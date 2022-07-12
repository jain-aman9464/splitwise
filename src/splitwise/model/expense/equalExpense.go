package expense

import (
	"container/list"
	"github.com/google/uuid"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/split"
)

type EqualExpense struct {
	Expense
}

func NewEqualExpense(amount float64, expensePaidBy model.User, splits list.List, title string) *EqualExpense {
	return &EqualExpense{struct {
		id            string
		amount        float64
		expensePaidBy model.User
		splits        list.List
		title         string
	}{id: uuid.New().String(), amount: amount, expensePaidBy: expensePaidBy, splits: splits, title: title}}
}

func (e *EqualExpense) Validate() bool {
	splits := e.GetSplits()

	for f := splits.Front(); f != nil; f = f.Next() {
		if _, ok := f.Value.(split.EqualSplit); !ok {
			return false
		}
	}

	return true
}
