class Solution:
    def climbStairs(self, n: int) -> int:
        staircase = [1, 2]
        for i in range(2, n):
            staircase.append(staircase[i-1] + staircase[i-2])
        return staircase[n-1]
