package calc

// retruns sum of all integers with wrror
func Add(numbers ...int) (string, int) {

	sum := 0

	if len(numbers) < 2 {
		return "Provide more than 2 arguments", sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return "", sum
	}
}
