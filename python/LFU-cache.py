class CacheNode:
    def __init__(self,value):
        self.value=value
        self.frequency=0

class LFUCache:
    
    def __init__(self, capacity):
        self.capacity=capacity
        self.cache=dict()
        self.frequencies=[]
        

    def get(self, key):
        if key in self.cache:
            pos=self.cache[key]
            node=self.frequencies[pos]
            node.frequency+=1

            while pos<len(self.frequencies)-1 and self.frequencies[pos+1].frequency<=node.frequency:
                self.frequencies[pos]=self.frequencies[pos+1]
                pos+=1
            self.frequencies[pos]=node
            self.cache[key]=pos
            return self.frequencies[pos].value
        return -1
        

    def put(self, key, value):
        if key in self.cache:
            self.get(key)
            return
        cacheNode=CacheNode(value)
        self.cache[key]=0
        if len(self.frequencies)<self.capacity:
            self.frequencies.insert(0,cacheNode)
        else:
            self.frequencies[0]=cacheNode
        