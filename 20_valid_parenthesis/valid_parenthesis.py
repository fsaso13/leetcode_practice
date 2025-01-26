class Solution:
    def isValid(self, s: str) -> bool:
        if len(s)%2 != 0:
            return False
        
        sol = []
        for i in s:
            if i == "[" or i == "(" or i == "{":
                sol.append(i)
            else:
                if len(sol) == 0:
                    return False
                elif i == "]" and sol[-1] == "[" or i== ")" and sol[-1] == "(" or i == "}" and sol[-1] == "{":    
                    sol = sol[:len(sol)-1]
                else:
                    return False
        
        return len(sol) == 0
    

if __name__=="__main__":
    sol = Solution()

    s = "()"
    print(sol.isValid(s))

    s = "()[]{}"
    print(sol.isValid(s))

    s = "(]"
    print(sol.isValid(s))

    s = "([])"
    print(sol.isValid(s))

    s = "))"
    print(sol.isValid(s))