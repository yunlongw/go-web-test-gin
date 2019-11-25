package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, a(i))
	}

	ad2 := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, ad2 = ad2(i)
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, sum)
	}
}
