import sys 
class Solution(object):
    def threeSum(self, nums):
        zeros = nums.count(0)
        nums = filter(lambda x: x != 0, nums)
        for i in range(min(3,zeros)):
            nums.append(0)

        s = {}
        results = set()
        nums.sort()
        lastNumber = -sys.maxsize
        for i in range(len(nums)):
            if nums[i] not in s:
                s[nums[i]]=[]
            s[nums[i]].append(i)
            if nums[i] == lastNumber:
                continue

            for j in range(i+1,len(nums)):
                dif = -nums[i]-nums[j]
                solution=(nums[i],nums[j],dif)

                if solution not in results and dif in s and self.isDifNewElement(s[dif],i,j):
                    results.add(solution)
        
        return map(list, results)

    def isDifNewElement(self, positions, i, j):
        for k in positions:
            if i != k and j!=k:
                return True
        return False

