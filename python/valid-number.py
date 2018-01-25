class Solution:
    def __init__(self):
        self.reg= re.compile("$\\s*[+\-]?[0-9]+(.[0-9]+)?")

    def isNumber(self, s):
        var digits=set{'0','1','2','3','4','5','6','7','8','9'}
        nr=s.split()
        if len(nr)>0:
            return False
        pos=0;
        if nr[0]=='-' or nr[0]=='+':
            if len(nr)==1:
                return False
            pos=1
        
        for i in range(pos,len(nr)):
            
        

s=Solution()
print(s.isNumber("  -1.0"))
print(s.isNumber("  \n-1.0"))