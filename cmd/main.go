package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/tokopedia/test/splitwise/src/splitwise/model"
	"github.com/tokopedia/test/splitwise/src/splitwise/model/split"
	"github.com/tokopedia/test/splitwise/src/splitwise/repo"
	"github.com/tokopedia/test/splitwise/src/splitwise/service"
	"os"
	"strconv"
	"strings"
)

func main() {
	user1 := model.NewUser(1, "Aman", "u1@gmai.com", "1111111111")
	user2 := model.NewUser(2, "Tiwari", "u2@gmai.com", "2222222222")
	user3 := model.NewUser(3, "Vatsal", "u3@gmai.com", "3333333333")
	user4 := model.NewUser(4, "Annu", "u4@gmai.com", "4444444444")

	expenseRepo := repo.NewExpenseRepository()
	userService := service.NewUserService(expenseRepo)

	userService.AddUser(user1)
	userService.AddUser(user2)
	userService.AddUser(user3)
	userService.AddUser(user4)

	splitService := service.NewSplitService(expenseRepo)
	expenseService := service.NewExpenseService(expenseRepo)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		vals := strings.Split(text, " ")

		switch vals[0] {
		case "EXPENSE":
			username := vals[1]
			amountSpent, _ := strconv.ParseFloat(vals[2], 64)
			totalMembers, _ := strconv.Atoi(vals[3])
			splits := list.List{}
			expenseIndex := 3 + totalMembers + 1
			expense := vals[expenseIndex]

			switch expense {
			case "EQUAL_SPLIT":
				for i := 0; i < totalMembers; i++ {
					splits.PushBack(split.NewEqualSplit(userService.GetUser(vals[i+4])))
				}
				expenseService.CreateExpense("EQUAL_SPLIT", amountSpent, username, splits, "GoaFlight")

			case "EXACT_SPLIT":
				for i := 0; i < totalMembers; i++ {
					amountFloat64, _ := strconv.ParseFloat(vals[expenseIndex+i+1], 64)
					splits.PushBack(split.NewExactSplit(userService.GetUser(vals[i+4]), amountFloat64))
				}
				expenseService.CreateExpense("EXACT_SPLIT", amountSpent, username, splits, "CabTickets")
			case "PERCENT_SPLIT":
				for i := 0; i < totalMembers; i++ {
					percentFlota64, _ := strconv.ParseFloat(vals[expenseIndex+i+1], 64)
					splits.PushBack(split.NewPercentSplit(userService.GetUser(vals[i+4]), percentFlota64))
				}
				expenseService.CreateExpense("PERCENT_SPLIT", amountSpent, username, splits, "Dinner")
			}
		case "SHOW":
			if len(vals) == 1 {
				splitService.ShowBalances()
			} else {
				splitService.ShowBalance(vals[1])
			}
		case "QUIT":
			fmt.Println("QUITTING ! ! !!")
			return
		default:
			fmt.Println("No Expected Argument Found")
		}

	}
}
