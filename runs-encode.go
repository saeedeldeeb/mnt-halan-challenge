package main

import (
	"strconv"
)

func runLengthEncode(s string) string {
	var encodedStr string
	var prev rune
	count := 0

	for _, r := range s {
		if r != prev && count > 0 {
			encodedStr += string(prev) + strconv.Itoa(count)
			count = 0
		}
		prev = r
		count++
	}

	if count > 0 {
		encodedStr += string(prev) + strconv.Itoa(count)
	}

	return encodedStr
}

func runLengthDecode(s string) string {
	var decodedStr string
	var countStr, cur string

	for _, r := range s {
		if r <= '9' {
			countStr += string(r)
		} else if cur == "" {
			cur = string(r)
		} else {
			count, _ := strconv.Atoi(countStr)
			for i := 0; i < count; i++ {
				decodedStr += string(cur)
			}
			countStr = ""
			cur = string(r)
		}
	}

	if countStr != "" {
		count, _ := strconv.Atoi(countStr)
		for i := 0; i < count; i++ {
			decodedStr += string(cur)
		}
	}

	return decodedStr
}
