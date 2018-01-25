from math import sqrt

def generalizedGCD(num, arr):
    minimum=min(arr)
    divisors=set(getDivisors(minimum))
    for i in range(num):
        for d in divisors:
            if arr[i]%d!=0:
                divisors.remove(d)
    return max(divisors)
    pass

def getDivisors(n):
    d=set(1,n)
    for i in range(2,sqrt(n)):
        if n%i==0:
            d.add(i)
            d.add(n/i)
    return d