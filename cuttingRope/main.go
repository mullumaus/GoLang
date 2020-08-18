/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jian-sheng-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func cuttingRope(n int) int {
	cr := make([]int, n+1)
	cr[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			tmp := max(j, cr[j]) * max(i-j, cr[i-j])
			cr[i] = max(cr[i], tmp)
		}
	}
	return cr[n]
}

func main() {

}
