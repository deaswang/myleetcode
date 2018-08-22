/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := new(ListNode)
    l := ret
    p := l
    var val int
    var add bool = false
    for l1 != nil && l2 != nil {
        val = l1.Val + l2.Val
        if add {
            val += 1
            add = false
        }
        if val > 9 {
            l.Val = val - 10
            add = true
        } else {
            l.Val = val
        }
        l1 = l1.Next
        l2 = l2.Next
        l.Next = new(ListNode)
        p = l
        l = l.Next
    }
    if l1 == nil && l2 != nil{
        p.Next = l2
        l = p.Next
    } else if l2 == nil && l1 != nil{
        p.Next = l1
        l = p.Next
    }
    for ;add; {
        l.Val += 1
        add = false
        if l.Val > 9 {
            add = true
            l.Val -= 10
        }
        p = l
        if l.Next == nil {
            l.Next = new(ListNode)
        }
        l = l.Next
    }
    for ;l.Next != nil; {
        p = l
        l = l.Next
    }
    if l.Val == 0 {
        p.Next = nil
    }
    return ret
}
