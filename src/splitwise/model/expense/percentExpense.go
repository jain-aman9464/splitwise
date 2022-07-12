package expense

import (
	"container/list"
	"github.com/google/uuid"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/split"
)

type PercentExpense struct {
	Expense
}

func NewPercentExpense(amount float64, expensePaidBy model.User, splits list.List, title string) *PercentExpense {
	return &PercentExpense{struct {
		id            string
		amount        float64
		expensePaidBy model.User
		splits        list.List
		title         string
	}{id: uuid.New().String(), amount: amount, expensePaidBy: expensePaidBy, splits: splits, title: title}}
}

func (e *PercentExpense) Validate() bool {
	splits := e.GetSplits()
	totalSplitPercent := 0.0

	for f := splits.Front(); f != nil; f = f.Next() {
		if _, ok := f.Value.(split.PercentSplit); !ok {
			return false
		}
		percentSplit := f.Value.(split.PercentSplit)
		totalSplitPercent += percentSplit.GetPercent()
	}

	return 100.0 == totalSplitPercent
}
