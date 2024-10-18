package main

import (
	"fmt"
)

func main() {
	ans := countGoodIntegers(5, 6)
	fmt.Println("ans", ans)
	fmt.Println(composeNum(10, 3))
	fmt.Println(DigitsCount(4, [10]int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0}))
}

func countGoodIntegers(n int, k int) int64 {
	half := n / 2
	var ans int64
	if n == 1 {
		for i := 1; i < 10; i++ {
			if i%k == 0 {
				ans++
			}
		}
		return ans
	}
	var mapVisited = make(map[[10]int]bool)
	if n%2 == 1 {
		for num := powerTen(half - 1); num < powerTen(half); num++ {
			fmt.Println("num", num)
			for j := 0; j < 10; j++ {
				realNum := num*powerTen(half+1) + j*powerTen(half) + reverseNum(num)
				if mapVisited[numToDigits(realNum)] {
					continue
				}
				if realNum%k == 0 {
					mapVisited[numToDigits(realNum)] = true
					ans += DigitsCount(n, numToDigits(realNum))
				}
			}
		}
	} else {
		for num := powerTen(half - 1); num < powerTen(half); num++ {
			realNum := num*powerTen(half) + reverseNum(num)
			if mapVisited[numToDigits(realNum)] {
				continue
			}
			if realNum%k == 0 {
				ans += DigitsCount(n, numToDigits(realNum))
				mapVisited[numToDigits(realNum)] = true
			}
		}
	}
	return ans
}

func numToDigits(num int) [10]int {
	var digits [10]int
	for num > 0 {
		digits[num%10]++
		num /= 10
	}
	return digits
}

func DigitsCount(n int, digits [10]int) int64 {
	var ret int64 = 1
	var left = n
	for i := 0; i < 10; i++ {
		if digits[i] == 0 {
			continue
		}
		if i == 0 {
			ret *= composeNum(left-1, digits[i])
		} else {
			ret *= composeNum(left, digits[i])
		}
		left -= digits[i]
	}
	return ret
}

func composeNum(a int, b int) int64 {
	var ret = int64(1)
	for i := int64(0); i < int64(b); i++ {
		ret *= (int64(a) - i)
	}
	for i := int64(0); i < int64(b); i++ {
		ret /= (i + 1)
	}
	return ret
}

func powerTen(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

func nextIncNum(n int) int {
	n++
	var numZero int
	for n%10 == 0 {
		numZero++
		n /= 10
	}
	lastGoodDigit := n % 10
	for i := 0; i < numZero; i++ {
		n = n*10 + lastGoodDigit
	}
	return n
}

func reverseNum(n int) int {
	var res int
	for n > 0 {
		res = res*10 + n%10
		n /= 10
	}
	return res
}
