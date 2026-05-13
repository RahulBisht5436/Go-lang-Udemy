package recursion

import "fmt"

func main() {
	factor := factorial(5)
	fmt.Println(factor)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)

}
