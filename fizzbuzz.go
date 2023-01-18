package fizzbuzz

import "fmt"

func FizzBuzz(number int) string {
	if isFizzBuzz(number) {
		return "FizzBuzz"
	}
	if isBuzz(number) {
		return "Buzz"
	}
	if isFizz(number) {
		return "Fizz"
	}
	return fmt.Sprintf("%d", number)

}

func isFizz(num int) bool {
	return num%3 == 0
}
func isBuzz(num int) bool {
	return num%5 == 0
}
func isFizzBuzz(num int) bool {
	return num%15 == 0
}
