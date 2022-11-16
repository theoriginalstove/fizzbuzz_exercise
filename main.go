package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

// fizzBuzz will determine if the string to print out is
// divisable by 3 or 5 and return Fizz or Buzz respectively.
// if it's divisble by both then it will return FizzBuzz
// else it will return n as a string
func fizzBuzz(n int) string {
	res := ""
	if n%3 == 0 {
		res += "Fizz"
	}
	if n%5 == 0 {
		res += "Buzz"
	}
	if res != "" {
		return res
	}

	return fmt.Sprintf("%d", n)
}
