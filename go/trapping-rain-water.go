package main

import "math"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	amount := 0
	maxBefore := height[0]
	for i := 1; i < len(height)-1; i++ {
		if height[i] < maxBefore {
			amount += (maxBefore - height[i])
		} else {
			maxBefore = height[i]
		}
	}

	maxAfter := height[len(height)-1]
	for i := len(height) - 2; i > 0; i-- {
		if height[i] >= maxBefore || maxAfter >= maxBefore {
			break
		}

		amount -= (maxBefore - int(math.Max(float64(maxAfter), float64(height[i]))))
		if height[i] > maxAfter {
			maxAfter = height[i]
		}
	}

	return amount
}
