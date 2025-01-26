class Solution:
    def isPalindrome(self, x: int) -> bool:
        numStringed = str(x)
        lhs = 0
        rhs = len(numStringed)-1
        for i in range(len(numStringed)//2):
            if numStringed[lhs] != numStringed[rhs]:
                return False
            
            lhs+=1
            rhs-=1
        
        return True
    
if __name__ == "__main__":
    sol = Solution()

    x = 121
    print(sol.isPalindrome(x))

    x = -121
    print(sol.isPalindrome(x))

    x = 10
    print(sol.isPalindrome(x))

    x = 1221
    print(sol.isPalindrome(x))


