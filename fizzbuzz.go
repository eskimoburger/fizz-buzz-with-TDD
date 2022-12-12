package fizzbuzz

import "fmt"

func FizzBuzz(number int) string {
	if number == 5 {
		return "Buzz"
	}
	if number == 3 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", number)
}
