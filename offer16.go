package main

import (
	"fmt"
)

/**
剑指 Offer 16. 数值的整数次方
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104

思路：
使用指数迭代的方式
log n 时间复杂度
*/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -1 * n
	}
	var res float64 = 1
	for n > 0 {
		fmt.Println(x, n, res)
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(myPow(2, 10))
}
