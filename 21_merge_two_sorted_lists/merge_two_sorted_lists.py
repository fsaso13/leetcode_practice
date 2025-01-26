from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 == None:
            return list2
        if list2 == None:
            return list1

        lhs = list1
        rhs = list2
        head = list1
        if list2.val < list1.val:
            head = list2
            rhs = rhs.next
        else:
            lhs = lhs.next
        curr = head

        while lhs != None or rhs != None:
            if rhs == None:
                curr.next = lhs
                curr = curr.next
                lhs = lhs.next
            elif lhs == None:
                curr.next = rhs
                curr = curr.next
                rhs = rhs.next
            elif lhs.val <= rhs.val:
                curr.next = lhs
                curr = curr.next
                lhs = lhs.next
            else:
                curr.next = rhs
                curr = curr.next
                rhs = rhs.next      

        return head
    
if __name__ == "__main__":
    sol = Solution()

    # list1 = [1,2,4]
    l3 = ListNode(4, None)
    l2 = ListNode(2, l3)
    l1 = ListNode(1, l2)
    # list2 = [1,3,4]
    r3 = ListNode(4, None)
    r2 = ListNode(3, r3)
    r1 = ListNode(1, r2)

    newHead = sol.mergeTwoLists(l1, r1)

    while newHead != None:
        print(newHead.val)
        newHead = newHead.next

    # # list1 = []
    l1 = None
    # # list2 = []
    r1 = None
    newHead = sol.mergeTwoLists(l1, r1)
    while newHead != None:
        print(newHead.val)
        newHead = newHead.next

    # # list1 = []
    l1 = None
    # # list2 = []
    r1 = ListNode(0)
    newHead = sol.mergeTwoLists(l1, r1)
    while newHead != None:
        print(newHead.val)
        newHead = newHead.next

    l1 = ListNode(2, None)
    r1 = ListNode(1, None)
    newHead = sol.mergeTwoLists(l1,r1)
    while newHead != None:
        print(newHead.val)
        newHead = newHead.next