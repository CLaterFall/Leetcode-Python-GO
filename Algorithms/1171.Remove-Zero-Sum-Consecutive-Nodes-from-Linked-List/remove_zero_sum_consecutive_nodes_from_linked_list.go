/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    records := make(map[int]ListNode)
    p := head
    pre_sum := 0
    for {
        if p == nil {
            break
        }
        pre_sum += p.Val
        // 总数为0 则前面的都可弃掉，直接从next开始
        if pre_sum == 0 {
            head = p.Next
        } else {
            records[pre_sum] = *p
        }
        p = p.Next
    }
    
    p = head
    pre_sum = 0
    for {
        if p == nil {
            break
        }
        pre_sum += p.Val
        p.Next = records[pre_sum].Next
        p = p.Next
    }
    return head
}
