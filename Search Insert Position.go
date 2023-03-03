class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:  
        idx = 0;  
        for i in range(len(nums)):
            if target > nums[i]:
                idx = i + 1
                continue
            if target == nums[i]:
                idx = i 
                continue
        return idx
            
