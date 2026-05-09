package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	CurrentBalance := readBalanceFiles()
	fmt.Println("Welcome to the Go Bank")
	fmt.Println("What do Want to Do ??")
	fmt.Println("1: Check the Balance")
	fmt.Println("2: Deposite Money ??")
	fmt.Println("3: Withdraw Money ??")
	fmt.Println("4: Exit")

	var operation int64
	for operation != 4 {
		print("Enter Your Opton : ")
		fmt.Scan(&operation)

		if operation == 1 {
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
		}
		if operation == 2 {
			CurrentBalance = DepositeMoney(CurrentBalance)
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
			writeBalanceToFiles(CurrentBalance)
		}
		if operation == 3 {
			CurrentBalance = WithdrawMoney(CurrentBalance)
			fmt.Printf("Your Current Balanace is : %v \n", CurrentBalance)
			writeBalanceToFiles(CurrentBalance)
		}
	}

}

func DepositeMoney(CurrentBalance int) int {
	var depositeAmount int
	print("Enter the deposite Amount : ")
	fmt.Scan(&depositeAmount)
	if depositeAmount < 0 {
		println("Wrong Amount Entered , Kindly Enter Correct Value")
		return CurrentBalance
	}
	return depositeAmount + CurrentBalance
}

func WithdrawMoney(CurrentBalance int) int {
	var withdrawAmount int
	print("Enter the withdraw Amount : ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > CurrentBalance || withdrawAmount < 0 {
		println(" Sorry TransSaction can be Processed: withDraw Amount is larger than Balance")
		return CurrentBalance
	}
	return CurrentBalance - withdrawAmount
}

// writeBalanceToFiles writes the current balance into a file
func writeBalanceToFiles(CurrentBalance int) {

	// Convert the integer balance into a string
	// Example: 500 -> "500"
	var CurrentBalanceByte = fmt.Sprint(CurrentBalance)

	// WriteFile expects data in byte format ([]byte)
	// so we convert the string into bytes
	//
	// "balanceData.txt" -> file name
	// []byte(CurrentBalanceByte) -> file content in byte format
	// 0644 -> file permission
	//
	// 0 = special mode
	// 6 = owner can read + write
	// 4 = group can only read
	// 4 = others can only read
	os.WriteFile("balanceData.txt", []byte(CurrentBalanceByte), 0644)
}

func readBalanceFiles() int {
	data, err := os.ReadFile("balanceData.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	balance, errBalance := strconv.Atoi(string(data))
	if errBalance != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println(balance)
	return balance
}
