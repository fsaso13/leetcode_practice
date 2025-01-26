from typing import List

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:        
        prefixLength = 0
        for i in range(len(strs[0])):
            for word in strs[1:]:
                if i == len(word) or word[i] != strs[0][i]:
                    return strs[0][:prefixLength]
            prefixLength+=1

        return strs[0][:prefixLength]

if __name__ == "__main__":
    sol = Solution()

    strs = ["flower","flow","flight"]
    print(sol.longestCommonPrefix(strs))

    strs = ["dog","racecar","car"]
    print(sol.longestCommonPrefix(strs))