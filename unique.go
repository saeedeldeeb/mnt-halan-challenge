package main

func unique(arr []string) []string {
	var uqArr []string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		for j := 0; j < len(arr); j++ {
			if cur == arr[j] && i != j {
				break
			}
			if j == len(arr)-1 {
				uqArr = append(uqArr, cur)
			}
		}
	}
	return uqArr
}
