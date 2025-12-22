package main

// EvenOrOdd returns "even" if the number is even, otherwise "odd".
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}