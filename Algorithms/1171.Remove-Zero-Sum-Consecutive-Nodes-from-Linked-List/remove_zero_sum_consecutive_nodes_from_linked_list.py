# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeZeroSumSublists(self, head: ListNode) -> ListNode:
        p = head
        records = {}
        pre_sum = 0
        while p:
            pre_sum += p.val
            if pre_sum == 0:
                head = p.next
                records.clear()
                p = p.next
            else:
                if pre_sum in records:
                    records[pre_sum].next = p.next
                    p = head
                    pre_sum = 0
                    records.clear()
                else:
                    records[pre_sum] = p
                    p = p.next
        return head
                
