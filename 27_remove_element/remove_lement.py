from typing import List

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums) == 0:
            return 0
        
        if len(nums) == 1:
            if nums[0] != val:
                return 1
            else:
                return 0
        
        lhs = 0
        rhs = len(nums) - 1

        while nums[rhs] == val:
            rhs-= 1

        while lhs < rhs:
            if nums[lhs] != val:
                lhs+= 1
            else:
                nums[lhs],nums[rhs] = nums[rhs],nums[lhs]
                lhs+=1
                while nums[rhs] == val:
                    rhs-=1
        return rhs+1
            
if __name__=="__main__":
    sol = Solution()

    nums = [3,2,2,3]
    val = 3
    print(sol.removeElement(nums,val))
    print(nums)

    nums = [0,1,2,2,3,0,4,2]
    val = 2
    print(sol.removeElement(nums,val))
    print(nums)