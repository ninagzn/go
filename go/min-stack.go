package main

type MinStack struct {
	list []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
		return
	}
	min := this.mins[len(this.mins)-1]
	if x < min {
		min = x
	}
	this.mins = append(this.mins, min)
}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	this.list = this.list[0 : len(this.list)-1]
	this.mins = this.mins[0 : len(this.mins)-1]
}

func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.mins[len(this.mins)-1]
}
