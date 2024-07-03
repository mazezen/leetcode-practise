package main

import (
	"fmt"
	"math"
)

func main() {
	harshadNumber := sumOfTheDigitsOfHarshadNumber(100)
	fmt.Println(harshadNumber)
}

// 如果一个整数能够被其各个数位上的数字之和整除， 则称之为 哈沙德数（Harshad number）。
// 给你一个整数 x 。如果 x 是 哈沙德数 ，则返回 x 各个数位上的数字之和，否则，返回 -1 。
func sumOfTheDigitsOfHarshadNumber(x int) int {
	if x == 0 {
		return -1
	}

	if 9 >= x && x > 0 {
		return x
	}

	if x >= -9 && x <= -1 {
		return int(math.Abs(float64(x)))
	}

	var m int   // 定义m 存个位数
	var n = 1   // 定义n 存高位数
	var sum int // 定义sum 计算个位数之合
	// 开启循环
	for n > 0 {
		// 依次存高位数，直到最高的位数
		n = x / 10
		// 计算个位数
		m = x % 10

		// 计算个位数之和
		sum = sum + m
		// n 就是最高位了
		if n < 10 {
			// 把最高位的数也计算上
			sum = sum + n
			n = 0 // 当只剩一个高位数时，将n置为0
		} else {
			// 如果不是最高位，将n的结果赋值给x
			x = n
		}
	}

	if x%sum == 0 {
		return sum
	}

	return -1
}
