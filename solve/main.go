package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, playground")
	//fmt.Println(calExpr("2*3+2"))
	fmt.Println(addOperators("1000000009", 9))
	fmt.Println(time.Since(t))
}
func addOperators(num string, target int) []string {
	exprCasesNum := 1 << ((len(num) - 1) * 2)
	var ans []string
	for i := 0; i < exprCasesNum; i++ {
		exp := convert(i, num)
		if isValidExpr(exp) {
			if calExpr(exp) == target {
				ans = append(ans, exp)
			}
		}
	}
	return ans
}

func convert(exprCasesNum int, e string) string {
	var n int = len(e)
	for exprCasesNum != 0 {
		n--
		if exprCasesNum%4 == 1 {
			e = e[:n] + "+" + e[n:]
		} else if exprCasesNum%4 == 2 {
			e = e[:n] + "-" + e[n:]
		} else if exprCasesNum%4 == 3 {
			e = e[:n] + "*" + e[n:]
		}
		exprCasesNum /= 4
	}
	return e
}

func isValidExpr(e string) bool {
	//if number is start with 0, it is invalid
	var lastTokenIndex int = -1
	for i := 0; i < len(e); i++ {
		if e[i] == '+' || e[i] == '-' || e[i] == '*' {
			lastTokenIndex = i
		}
		if e[i] == '0' && i-lastTokenIndex == 1 {
			if i+1 < len(e) {
				if e[i+1] >= '0' && e[i+1] <= '9' {
					return false
				}
			}
		}
	}
	return true
}

func calExpr(e string) int {
	nums, tokens := getAllNumberAndTokenInExpr(e)
	var stack []int
	stack = append(stack, nums[0])
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == '*' {
			stack[len(stack)-1] = stack[len(stack)-1] * nums[i+1]
		} else {
			stack = append(stack, nums[i+1])
		}
	}
	var tokenIndex int = 0
	var res int = stack[0]
	for i := 1; i < len(stack); i++ {
		for tokenIndex < len(tokens) && tokens[tokenIndex] == '*' {
			tokenIndex++
		}
		if tokens[tokenIndex] == '+' {
			res += stack[i]
		} else if tokens[tokenIndex] == '-' {
			res -= stack[i]
		}
		tokenIndex++
	}
	return res
}

func getAllNumberAndTokenInExpr(e string) ([]int, []byte) {
	var nums []int
	var tokens []byte
	var lastTokenIndex int = -1
	for i := 0; i < len(e); i++ {
		if e[i] == '+' || e[i] == '-' || e[i] == '*' {
			nums = append(nums, composeNum(e[lastTokenIndex+1:i]))
			tokens = append(tokens, e[i])
			lastTokenIndex = i
		}
	}
	nums = append(nums, composeNum(e[lastTokenIndex+1:]))
	return nums, tokens
}

func composeNum(e string) int {
	var res int
	for i := 0; i < len(e); i++ {
		res = res*10 + int(e[i]-'0')
	}
	return res
}
