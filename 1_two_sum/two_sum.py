from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashMap = {}

        for key, val in enumerate(nums):
            hashMap[val] = key
        
        for key2, val in enumerate(nums):
            if target-val in hashMap.keys() and key2 != hashMap[target-val]:
                return [key2, hashMap[target-val]]

        

if __name__ == "__main__":
    sol = Solution()
    nums = [2,7,11,15]
    target = 9
    print( sol.twoSum(nums, target))

    nums = [3,2,4]
    target = 6
    print( sol.twoSum(nums, target))

    nums = [3,3]
    target = 6
    print( sol.twoSum(nums, target))
