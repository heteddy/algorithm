# -*- coding:utf-8 -*-
import sys


class Solution:
    def coinChange(self, coins: list, amount: int) -> int:
        # 初始化dp， 横轴代表金额，纵轴代表硬币
        dp = [[sys.maxsize for _ in range(1, amount + 1)] for _ in range(0, len(coins))]

        # dp[i][j] 表示 使用coins[0-i]表示amount[j]块钱，最少的硬币数

        for i, coin_val in enumerate(coins):
            for amount_val in range(1, amount + 1):
                j = amount_val - 1
                # print(i, coin_val, j, amount_val)
                if amount_val == coin_val:
                    dp[i][j] = 1
                elif amount_val < coin_val:
                    # dp[i][j] = sys.maxsize
                    pass
                else:  # amount_val>coin_val
                    if i == 0:
                        if amount_val % coin_val == 0:
                            dp[i][j] = amount_val // coin_val
                    else:
                        # 使用当前硬币，那么余额
                        left = amount_val - coin_val
                        s = sys.maxsize

                        # 余额的最少硬币
                        for _i in range(i + 1):
                            s = min(s, dp[_i][left - 1])
                        print("amount_val", amount_val, "coin", coin_val, "剩余金额", left, "硬币数", s)

                        dp[i][j] = min(1 + s, dp[i - 1][j])

                        print("dp[{0}][{1}]={2}".format(i, j, dp[i][j]))

        return dp[len(coins) - 1][amount - 1]