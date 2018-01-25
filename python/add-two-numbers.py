class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        result=None
        cur=None
        while l1 is not None and l2 is not None:
            p=ListNode(l1.val+l2.val)
            l1=l1.next
            l2=l2.next
            if result is None:
                result=p
                cur=result
            else:
                cur.next=p
                self.balanceNode(cur)
                cur=cur.next
        
        l=None   
        if l1 is None:
            l=l2
        else:
            l=l1
            
        while l is not None:
            p=ListNode(l.val)
            l=l.next
            if result is None:
                result=p
                cur=result
            else:
                cur.next=p
                self.balanceNode(cur)
                cur=cur.next

        if cur.val>9:
            p=ListNode(1)
            cur.val%=10
            cur.next=p

        return result

    def balanceNode(self,p):
        if p.next is not None and p.val>9:
                p.next.val+=1
                p.val%=10

s=Solution()
l1=ListNode(8)
l1.next=ListNode(6)
l2=ListNode(4)
l2.next=ListNode(8)
l2.next.next=ListNode(9)
r=s.addTwoNumbers(l1, l2)
while r is not None:
     print (r.val)
     r=r.next