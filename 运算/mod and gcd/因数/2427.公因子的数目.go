/*
 * @lc app=leetcode.cn id=2427 lang=golang
 * @lcpr version=21909
 *
 * [2427] 公因子的数目
 */

// @lc code=start
// 实质上是约数
package main

func commonFactors(a int, b int) (ans int) {
	g := gcd(a, b)
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g { // i*i==g时候，只能算1个
				ans++
			}
		}
	}
	return
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 分解质因数
// void divide(int x)
// {
//     for (int i = 2; i <= x / i; i ++ )
//     {
//         if(x % i == 0)// i 一定是质数
//         {
//             int s = 0;
//             while(x % i == 0)
//             {
//                 x /= i;
//                 s ++;
//             }

//             printf("%d %d\n", i, s);
//         }
//     }
//     if(x > 1) printf("%d %d\n", x, 1);

//     puts("");
// }

// @lc code=end

/*
// @lcpr case=start
// 12\n6\n
// @lcpr case=end

// @lcpr case=start
// 25\n30\n
// @lcpr case=end

*/
