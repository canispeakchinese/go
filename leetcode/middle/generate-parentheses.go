package main

import "fmt"

func cal(num int, left int, right int) (ans []string) {
	if left == num && right == num {
		return []string{""}
	} else if left == num {
		next := cal(num, left, right+1)
		ans = make([]string, 0, len(next))
		for _, item := range next {
			ans = append(ans, ")"+item)
		}
	} else if left == right {
		next := cal(num, left+1, right)
		ans = make([]string, 0, len(next))
		for _, item := range next {
			ans = append(ans, "("+item)
		}
		return ans
	} else {
		next1 := cal(num, left+1, right)
		next2 := cal(num, left, right+1)
		ans = make([]string, 0, len(next1)+len(next2))
		for _, item := range next1 {
			ans = append(ans, "("+item)
		}
		for _, item := range next2 {
			ans = append(ans, ")"+item)
		}
	}
	return ans
}

func generateParenthesis(n int) []string {
	return cal(n, 0, 0)
}

func main() {
	fmt.Println(generateParenthesis(3))
}
