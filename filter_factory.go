package main

import "fmt"

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isBiggerThan4(integer int) bool {
	if integer > 4 {
		return true
	}
	return false
}

func filter_factory(f func(int) bool) func(s []int) (yes, no []int) {
	return func(s []int) (yes, no []int) {
		for _, value := range s {
			if f(value) {
				yes = append(yes, value)
			} else {
				no = append(no, value)
			}
		}
		return
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)
	odd_even_function := filter_factory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)

	bigger, smaller := filter_factory(isBiggerThan4)(s)
	fmt.Println("Bigger than 4: ", bigger)
	fmt.Println("Smaller than or equal to 4: ", smaller)
}
