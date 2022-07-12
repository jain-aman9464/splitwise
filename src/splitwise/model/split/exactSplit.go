package split

import "github.com/tokopedia/test/splitwise/src/splitwise/model"

type ExactSplit struct {
	Split
}

func NewExactSplit(user model.User, amount float64) ExactSplit {
	return ExactSplit{
		Split{
			user:   user,
			amount: amount,
		}}
}
