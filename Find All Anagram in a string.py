#Find All Anagram in a string 
#leetcode 424 - Medium Level

class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(p) > len(s):
            return []
        countP , countS = {}, {}
        for i in range(len(p)):
            countP[p[i]] = 1 + countP.get(p[i], 0)
            countS[s[i]] = 1 + countS.get(s[i], 0)
        ans = [0] if countS == countP else []
        y = 0
        for row in range(len(p), len(s)):
            countS[s[row]] = 1 + countS.get(s[row], 0)
            countS[s[y]] -= 1
            if countS[s[y]] == 0:
                countS.pop(s[y])
            y += 1
            if countS == countP:
                ans.append(y)
        return ans
