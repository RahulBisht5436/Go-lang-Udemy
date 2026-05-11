package main

import "fmt"

func main() {
	hobbyArray := []string{
		"coding", "anime", "travelling",
	}

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	printHobby(hobbyArray)
	fmt.Println(hobbyArray[0])
	fmt.Println(hobbyArray[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	newSlice := hobbyArray[:3]
	fmt.Printf("This is the Slice Data : %v \n", newSlice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	resizedSlice := newSlice[1:]
	fmt.Printf("Resized Array From the Slice : %v \n", resizedSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	careerGoals := []string{
		"crack the MNCS or product based Company",
		"Start the fitness Journey and become fit",
	}
	fmt.Printf("Dynamic Araay with career Goals : %v \n", careerGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	careerGoals[0] = "crack the MNCS or product based Company V2"
	careerGoals = append(careerGoals, "Find a BEautiful Wife and have kids")
	fmt.Printf("This is the final array : %v \n", careerGoals)

}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
func printHobby(hobbdyArray []string) {
	fmt.Println(hobbdyArray[:])

}

// Time to practice what you learned!

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
