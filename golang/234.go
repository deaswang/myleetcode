/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var s = make([]int, 0)
    for n := head; n != nil; n=n.Next {
        s = append(s, n.Val)
    }
    var l = len(s)
    for i, j := 0, l-1; i<j; {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
