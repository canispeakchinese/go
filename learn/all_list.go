package main

import "fmt"

func allList(curr, n int, used map[int]bool) [][]int {
	res := [][]int{}
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}

		curr_res := []int{i}
		used[i] = true
		if curr < n {
			next_res := allList(curr+1, n, used)
			for _, item := range next_res {
				res = append(res, append(curr_res, item...))
			}
		} else {
			res = append(res, curr_res)
		}
		used[i] = false
	}
	return res
}

func main() {
	fmt.Println(allList(1, 5, map[int]bool{}))
}
