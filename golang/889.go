/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
    var root *TreeNode
    root = new(TreeNode)
    root.Val = pre[0]
    l := len(pre)
    if l == 1 {
        return root
    }
    var cut int
    for i, v := range post{
        if v == pre[1] {
            cut = i
            break
        }
    }
    if cut == l-2 {
        root.Left = constructFromPrePost(pre[1:cut+2], post[:cut+1])
    } else {
        root.Left = constructFromPrePost(pre[1:cut+2], post[:cut+1])
        root.Right = constructFromPrePost(pre[cut+2:], post[cut+1:l-1])
    }
    
    return root
}
