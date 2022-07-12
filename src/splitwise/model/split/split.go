package split

import "github.com/tokopedia/test/splitwise/src/splitwise/model"

type Split struct {
	user    model.User
	amount  float64
	percent float64
}

func NewSplit(user model.User, amount float64, percent float64) Split {
	return Split{
		user:    user,
		amount:  amount,
		percent: percent,
	}
}

func (s *Split) SetUser(user model.User) {
	s.user = user
}

func (s Split) GetUser() model.User {
	return s.user
}

func (s *Split) SetAmount(amount float64) {
	s.amount = amount
}

func (s Split) GetAmount() float64 {
	return s.amount
}

func (s *Split) SetPercent(percent float64) {
	s.percent = percent
}

func (s Split) GetPercent() float64 {
	return s.percent
}
