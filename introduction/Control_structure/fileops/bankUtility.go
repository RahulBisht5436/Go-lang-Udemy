package fileops

import (
	"fmt"
	"os"
	"strconv"
)

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
func WriteBalanceToFiles(CurrentBalance int) {

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

func ReadBalanceFiles() int {
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
