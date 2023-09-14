package maxnumltn

import "fmt"

func ExampleGetMaxDigitLtD() {
	fmt.Println(getMaxNumLTN([]int{1, 2, 9, 4}, 2533))
	fmt.Println(getMaxNumLTN([]int{1, 2, 5, 4}, 2543))
	fmt.Println(getMaxNumLTN([]int{1, 2, 5, 4}, 2541))
	fmt.Println(getMaxNumLTN([]int{1, 2, 9, 4}, 2111))
	fmt.Println(getMaxNumLTN([]int{5, 9}, 5555))

	// output:
	// 2499
	// 2542
	// 2525
	// 1999
	// 999
}
