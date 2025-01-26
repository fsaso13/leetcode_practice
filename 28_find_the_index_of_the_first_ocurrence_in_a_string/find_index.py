# import string

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        found = haystack.find(needle)
        if found != -1:
            return haystack.index(needle)
        else:
            return -1
        
if __name__ == "__main__":
    sol = Solution()

    haystack = "sadbutsad"
    needle = "sad"
    print(sol.strStr(haystack, needle))

    haystack = "leetcode"
    needle = "leeto"
    print(sol.strStr(haystack, needle))