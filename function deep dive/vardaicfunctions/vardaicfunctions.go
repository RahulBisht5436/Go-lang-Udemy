package vardiacfunctions

import "fmt"

func main() {
	sumup("rahul bisht", 6, 2, 4, 2, 3, 2, 321, 23, 4, 12, 124)
}

//Can send here any number of values
func sumup(name string, nums ...int) {
	var total int
	for _, value := range nums {
		total += value
	}
	fmt.Println(total)
}
