package expense

import (
	"container/list"
	"github.com/google/uuid"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/split"
)

type ExactExpense struct {
	Expense
}

func NewExactExpense(amount float64, expensePaidBy model.User, splits list.List, title string) *ExactExpense {
	//uuid := uuid.New()

	return &ExactExpense{struct {
		id            string
		amount        float64
		expensePaidBy model.User
		splits        list.List
		title         string
	}{id: uuid.New().String(), amount: amount, expensePaidBy: expensePaidBy, splits: splits, title: title}}
}

func (e *ExactExpense) Validate() bool {
	splits := e.GetSplits()
	totalAmount := e.GetAmount()
	totalSplitAmount := 0.0

	for f := splits.Front(); f != nil; f = f.Next() {
		if _, ok := f.Value.(split.ExactSplit); !ok {
			return false
		}
		exactSplit := f.Value.(split.ExactSplit)
		totalSplitAmount += exactSplit.GetAmount()
	}

	return totalAmount == totalSplitAmount
}
