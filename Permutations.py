#Solution to Leetcode Exercice 46

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        if len(nums) == 1:
            return [nums[:]]
        for i in range(len(nums)):
            n = nums.pop(0)

            permut = self.permute(nums)
            for perm in permut:
                perm.append(n)
            res.extend(permut)
            nums.append(n)
        return res
