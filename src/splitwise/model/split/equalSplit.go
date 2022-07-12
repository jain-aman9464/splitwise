package split

import "github.com/tokopedia/test/splitwise/src/splitwise/model"

type EqualSplit struct {
	Split
}

func NewEqualSplit(user model.User) EqualSplit {
	return EqualSplit{
		Split{
			user: user,
		}}
}
