/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) != 0{
        qlen := len(queue)
        stage := []int{}
        for index := 0; index < qlen; index++ {
            node := queue[0]
            queue = queue[1:]
        if node != nil {
            stage = append(stage, node. Val)
            queue = append(queue , node.Left, node.Right)
        }
        }
        if len(stage) != 0{
            res = append(res, stage)
        }
    }
    return res
}
