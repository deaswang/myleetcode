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
    var left, right [][]int
    if root.Left != nil {
        left = levelOrder(root.Left)
    }
    if root.Right != nil {
        right = levelOrder(root.Right)
    }
    
    var ret = make([][]int, 0)
    ret = append(ret, []int{root.Val})
    if left == nil && right == nil{
        return ret
    } else if left == nil {
        ret = append(ret, right...)
    } else if right == nil {
        ret = append(ret, left...)
    } else {
        ll := len(left)
        lr := len(right)
        i := 0
        for ; i < ll && i < lr; i++ {
            ret = append(ret, append(left[i], right[i]...))
        }
        if i == ll && i < lr {
            ret = append(ret, right[i:]...)
        } else if i < ll && i == lr {
            ret = append(ret, left[i:]...)
        }
    }
    return ret
}

// Better
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
    
    var ret = new([][]int)
    levelDepth(ret, root, 0)
    return *ret
}

func levelDepth(ret *[][]int, node *TreeNode, depth int) {
    if len(*ret) < depth+1 {
        *ret = append(*ret, []int{})
    }
    (*ret)[depth] = append((*ret)[depth], node.Val)
    if node.Left != nil {
        levelDepth(ret, node.Left, depth+1)
    }
    if node.Right != nil {
        levelDepth(ret, node.Right, depth+1)
    }
}

