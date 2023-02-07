func twoSum(numbers []int, target int) []int {
    lft, rgt := 0, len(numbers) - 1
    for lft < rgt{
        if numbers[lft]+numbers[rgt] == target{
            return []int{lft + 1, rgt +1}}
        if numbers[lft] + numbers[rgt] > target{
            rgt--
        }
        if numbers[lft] + numbers[rgt] < target{
            lft++
        }
    }
    return nil
}
