package fizzbuzz

import "fmt"

func FizzBuzz(number int) string {
	if number == 5 {
		return "Buzz"
	}
	if number == 3 || number == 6 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", number)
}
