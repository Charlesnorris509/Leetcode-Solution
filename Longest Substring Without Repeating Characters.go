func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    ans := 0
    for l, r := 0, 0; r < len(s); r++{
        if idx, pt := mp[s[r]]; pt{
            l = max(l, idx)
        }
        ans = max(ans,r-l+1)
        mp[s[r]] = r + 1
    }
    return ans
}

func max(n, m int) int{
    if n > m{
        return n
    }
    return m
}
