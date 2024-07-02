package main

import "fmt"

func main() {
	// Problem 01
	fmt.Println("----- problem 01 -----")
	fmt.Println("Is 'anna' Palindrome? ", Palindrome("anna"))
	fmt.Println("Is 'level' Palindrome? ", Palindrome("level"))
	fmt.Println("Is 'noon' Palindrome? ", Palindrome("noon"))
	fmt.Println("Is 'rotator' Palindrome? ", Palindrome("rotator"))
	fmt.Println("Is 'apple' Palindrome? ", Palindrome("apple"))
	fmt.Println()

	// Problem 02
	fmt.Println("----- problem 02 -----")
	fmt.Println("renaming file ...")
	fmt.Println("mv original.txt changed.txt")
	file := File{Name: "original.txt"}
	err := file.Rename("changed.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()

	// Problem 03
	fmt.Println("----- problem 03 -----")
	encoded := runLengthEncode("aaaaaaaaaabbbaxxxxyyyzyx")
	fmt.Println("Encoded: ", encoded)
	fmt.Println("Decoded: ", runLengthDecode(encoded))
	fmt.Println()

	// Problem 04
	fmt.Println("----- problem 04 -----")
	h := compose(square, inc)
	fmt.Println(h(6))
	fmt.Println()

	// Problem 05
	fmt.Println("----- problem 05 -----")
	uniques := unique([]string{"apple", "banana", "apple", "orange", "banana", "apple"})
	fmt.Println(uniques)
	uniques = unique([]string{"apples", "apples"})
	fmt.Println(uniques)
	fmt.Println()

	// Problem 06
	fmt.Println("----- problem 06 -----")
	fmt.Println(transpose([][]int{{1, 2}, {3, 4}}))
	fmt.Println(transpose([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}))
	fmt.Println(transpose([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println()

	// Problem 07
	fmt.Println("----- problem 07 -----")
	fmt.Println("Steps to fill 4 liters container with 5 and 3 liters containers:")
	fmt.Println("1. Fill 5 liters container")
	fmt.Println("2. Pour 5 liters container to 3 liters container")
	fmt.Println("thinking ...")
	fmt.Println("I don't know how to solve this problem. ):")
	fmt.Println()

	// Problem 08
	fmt.Println("----- problem 08 -----")
	fmt.Println("First duplicate index: ", indexOfFirstDuplicate([]int{5, 17, 3, 17, 4, -1}))
	fmt.Println("Time complexity: O(n)")
	fmt.Println("Space complexity: O(n)")
	fmt.Println()

	// Problem 09
	fmt.Println("----- problem 09 -----")
	node := Node{
		value: 1,
		children: []Node{
			{
				value: 2,
				children: []Node{
					{
						value: 3,
						children: []Node{
							{
								value: 4,
								children: []Node{
									{
										value:    5,
										children: []Node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("Sum of nodes: ", sum(node))
}
