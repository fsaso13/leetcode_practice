class Solution:
    def romanToInt(self, s: str) -> int:
        romanDict = {"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}

        sum = 0
        finished = False
        lhs = 0
        rhs = 1

        while not finished:
            if lhs == len(s)-1:
                sum += romanDict[s[lhs]]
                finished = True
            elif romanDict[s[lhs]] < romanDict[s[rhs]]:
                sum += romanDict[s[rhs]] - romanDict[s[lhs]]
                lhs+=2
                rhs+=2
            else:
                sum += romanDict[s[lhs]]
                lhs +=1
                rhs +=1

            if lhs >= len(s):
                finished = True

        return sum

if __name__ == "__main__":
    sol = Solution()

    s ="III"
    print(sol.romanToInt(s))

    s ="LVIII"
    print(sol.romanToInt(s))

    s ="MCMXCIV"
    print(sol.romanToInt(s))