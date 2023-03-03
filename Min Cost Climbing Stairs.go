func minCostClimbingStairs(cost []int) int {
    minCost := 0
    costLen := len(cost)
    lastOne := 0
    lastTwo := cost[costLen-1]
    for start := costLen - 2; start >= 0; start-- {
        minCost = cost[start] + min(lastOne, lastTwo)
        lastOne = lastTwo
        lastTwo = minCost
    } 
    return min(lastOne, lastTwo)
}

func min(i, y int)int{
    if i < y{
        return i
    }
    return y
}
