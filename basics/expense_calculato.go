package main

import "fmt"
func sdmain(){
	var Revenue float64
	var Expense float64
	 
	fmt.Print("What is the Revenue :")
	fmt.Scan(&Revenue)

	fmt.Print("What is the Expense :")
	fmt.Scan(&Expense)
	profitBeforeTax := float64(Revenue) - float64(Expense)

	fmt.Println("Profit withoute text : ",profitBeforeTax)

	fmt.Println("%T",Revenue)
}