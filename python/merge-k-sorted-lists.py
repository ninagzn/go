class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """


class MinHeapLinkedList(object):
    def __init__(self, lists):
        self.elements=[sorted(lists, key=lambda x: x.val)


    def heapify(self, i):
        left = self.elements[2*i]
        right = self.elements[2*i+1]

        return
        
    
    def getMin(self):
        if len(self.elements) == 0:
            return None
        x = self.elements[0].val;
        if self.elements[0].next is None:
            self.elements[0] = self.elements[len(self.elements)-1];
            del self.elements[-1]
        else:
            self.elements[0] = self.elements[0].next

        self.heapify(0)
        return x

