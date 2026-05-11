package main

import "fmt"

type userMap map[string]float64

func main() {
	// Here make create the array with 2 block level of Space
	userName := make([]string, 2, 5)
	userName = append(userName, "Rahul Bisht")
	userName = append(userName, "Kamal Bisht")
	fmt.Println(userName)
	fmt.Printf("This is the first element %v", userName[2])

	// In same way we can use make for Maps
	// only differnce is we dont enter slots , just reserve the max space
	userRating := make(userMap, 10)
	userRating["Rahul Bisht"] = 10
	userRating["Sheetal Bisht"] = 6
	userRating["Kamal Bisht"] = 23
	userRating["Pareshwari Bisht"] = 11

	for key, value := range userRating {
		fmt.Printf("The key %v value is  : %v \n", key, value)
	}
}
