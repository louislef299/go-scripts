package main

import "fmt"

func main() {
	nums := 10
	code := make([]int, nums)
	count := 0
	secondCount := 0
	code[0] = nums - 1
	var succeeded bool
	end := make(map[int]int)
	for {
		r := make(map[int]int)
		for i := 0; i < nums; i++ {
			r[code[i]]++
		}

		var j int
		for j = 0; j < nums; j++ {
			if r[j] != code[j] {
				succeeded = false
				break
			}
		}
		if j == nums {
			succeeded = true
		}

		if !succeeded {
			if code[count]-1 == -1 {
				count++
			}
			code[count] = code[count] - 1
			for secondCount%10 == count {
				secondCount += 1
			}
			code[secondCount%10] += 1
		} else {
			end = r
			break
		}
		fmt.Println(r)
	}

	fmt.Println(end)
}
