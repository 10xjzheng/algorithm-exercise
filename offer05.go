package main

import "fmt"

/**
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

限制：

0 <= s 的长度 <= 10000

思路：直接将空格替换，用切片连接
*/

func ReplaceSpace(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			s = s[:i] + "%20" + s[i+1:]
		}
	}
	return s
}

func main() {
	fmt.Println(ReplaceSpace("ab dc d"))
}
