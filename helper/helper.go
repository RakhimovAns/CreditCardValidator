package helper

// Check this function check card by Luna's algorithm
func Check(number string) bool {
	var sum int = 0
	for i := 0; i < len(number); i++ {
		var digit int
		digit = int(number[i] - '0')
		if i%2 == 0 {
			digit *= 2
		}
		if digit > 10 {
			digit = digit%10 + digit/10
		}
		sum += digit
	}
	return sum%10 == 0
}
