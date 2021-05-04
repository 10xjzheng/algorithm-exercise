package main

import (
	"fmt"
)

/**
剑指 Offer 17. 打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


说明：

用返回一个整数列表来代替打印
n 为正整数
*/
func printNumbers(n int) []string {
	arr := []string{}
	for j := 0; j < 10; j++ {
		arr = append(arr, fmt.Sprintf("%v", j))
	}
	var delzero func(str string) string
	delzero = func(str string) string {
		if len(str) == 0 {
			return str
		}

		if str[0] == '0' {
			return delzero(str[1:])
		}

		return str
	}

	handle := func(arr []string) []string {
		resarr := []string{}
		for i := 0; i < 10; i++ {
			for _, v := range arr {
				s := fmt.Sprint(i) + v
				resarr = append(resarr, delzero(s))
			}
		}
		return resarr
	}

	for i := 1; i < n; i++ {
		arr = handle(arr)
	}

	return arr[1:]
}

func main() {
	fmt.Println(printNumbers(3))
}
