package main

// func Palindrome(s string) bool {
// 	return s == reverseString(s)
// }

// func reverseString(s string) string {
// 	var revStr string
// 	for i := len(s) - 1; i >= 0; i-- {
// 		revStr += string(s[i])
// 	}
// 	return revStr
// }

func Palindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
