package main

// Generic functions

func addNumbers(a, b int) int {
	return a + b
}

func ValidateUsername(s string) bool {
	l := len(s)
	var ans bool

	if l < 5 {
		ans = false
	} else if l > 10 {
		ans = false
	} else if l == 0 {
		ans = false
	} else {
		ans = true
	}

	return ans
}