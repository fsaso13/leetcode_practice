from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        pos = [0]
        last = nums[0]
        for key, val in enumerate(nums):
            if val != last:
                pos.append(key)
                last = val
        
        for key, val in enumerate(pos):
            nums[key], nums[val] = nums[val],nums[key]

        return len(pos)

if __name__ == "__main__":
    sol = Solution()

    nums = [1,1,2]
    print(sol.removeDuplicates(nums))
    print(nums)

    nums = [0,0,1,1,1,2,2,3,3,4]
    print(sol.removeDuplicates(nums))
    print(nums)