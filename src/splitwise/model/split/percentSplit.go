package split

import (
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
)

type PercentSplit struct {
	Split
}

func NewPercentSplit(user model.User, percent float64) PercentSplit {
	return PercentSplit{
		Split{
			user:    user,
			percent: percent,
		}}
}
