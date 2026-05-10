package main

import (
	"fmt"

	"example.com/introduction/fileops"
)

func main() {
	CurrentBalance := fileops.ReadBalanceFiles()
	presentOptions()

	var operation int64
	for operation != 4 {
		print("Enter Your Opton : ")
		fmt.Scan(&operation)

		if operation == 1 {
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
		}
		if operation == 2 {
			CurrentBalance = fileops.DepositeMoney(CurrentBalance)
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
			fileops.WriteBalanceToFiles(CurrentBalance)
		}
		if operation == 3 {
			CurrentBalance = fileops.WithdrawMoney(CurrentBalance)
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
			fileops.WriteBalanceToFiles(CurrentBalance)
		}
	}

}
