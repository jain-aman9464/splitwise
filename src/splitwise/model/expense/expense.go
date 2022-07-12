package expense

import (
	"container/list"
	"github.com/google/uuid"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
)

type Expense struct {
	id            string
	amount        float64
	expensePaidBy model.User
	splits        list.List
	title         string
}

type Expenser interface {
	Validate() bool
}

func NewExpense(amount float64, expensePaidBy model.User, splits list.List, title string) Expense {
	return Expense{
		id:            uuid.New().String(),
		amount:        amount,
		expensePaidBy: expensePaidBy,
		splits:        splits,
		title:         title,
	}
}

func (s *Expense) SetExpensePaidBy(user model.User) {
	s.expensePaidBy = user
}

func (s Expense) getExpensePaidBy() model.User {
	return s.expensePaidBy
}

func (s *Expense) SetAmount(amount float64) {
	s.amount = amount
}

func (s Expense) GetAmount() float64 {
	return s.amount
}

func (s *Expense) SetTitle(title string) {
	s.title = title
}

func (s Expense) GetTitle() string {
	return s.title
}

func (s *Expense) SetSplits(splits list.List) {
	s.splits = splits
}

func (s Expense) GetSplits() list.List {
	return s.splits
}
