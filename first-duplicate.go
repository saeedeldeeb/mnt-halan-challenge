package main

func indexOfFirstDuplicate(arr []int) int {
	var dupIndex int
	var hashMap = make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if hashMap[arr[i]] != 0 {
			dupIndex = i
			break
		} else {
			hashMap[arr[i]] = 1
		}
	}

	return dupIndex
}
