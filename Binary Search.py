class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums) - 1
        while (left < right):
            pivot = (right + left) // 2
            if nums[pivot] == target:
                return pivot
            elif nums[pivot] < target:
                left = pivot + 1
            elif nums[pivot] > target:
                right = pivot - 1

        if nums[left] == target:
            return left
        return -1 
