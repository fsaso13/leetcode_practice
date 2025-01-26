from typing import List

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            if target > nums[0]:
                return 1
            else:
                return 0
        lhs = 0
        rhs = len(nums)-1
        midpoint = (lhs+rhs)//2
        done = False

        while not done:
            if target == nums[midpoint]:
                done = True
            elif rhs-lhs <= 1:
                if target > nums[rhs]:
                    midpoint+=1
                elif target < nums[lhs]:
                    midpoint-=1
                midpoint+=1
                done = True
            elif target < nums[midpoint]:
                rhs = midpoint
                midpoint = (lhs+rhs)//2
            else:
                lhs = midpoint
                midpoint = (lhs+rhs)//2

        return midpoint
    
if __name__=="__main__":
    sol = Solution()

    nums = [1,3,5,6]
    target = 5
    print(sol.searchInsert(nums,target))

    nums = [1,3,5,6]
    target = 2
    print(sol.searchInsert(nums,target))

    nums = [1,3,5,6]
    target = 7
    print(sol.searchInsert(nums,target))

    nums = [1,3,4,5,6]
    target = 7
    print(sol.searchInsert(nums,target))



        
