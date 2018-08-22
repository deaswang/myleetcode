class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        for i, n in enumerate(nums):
            m[n] = i
        
        for i, n in enumerate(nums):
            if target - n in m and m[target - n] != i:
                return [m[target - n], i]
