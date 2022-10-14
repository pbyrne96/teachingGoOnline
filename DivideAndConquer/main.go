package main

import "fmt";

func max(values ...int) int{
	maxVal := values[0];
	for _, current := range values {
		if (current > maxVal) {
			maxVal = current
		}
	}
	return maxVal
}

func maxSubArray(nums []int) int {
	right := len(nums) -1
	return findMaxSubArray(nums, 0, right);
}

func maxCrossArray(nums []int, left int, right int, mid int) int {
	leftSum := -1<<31
	rightSum := -1<<31
	sums:=0
	for i := mid; i >= left; i++ {
		sums += nums[i];
		leftSum = max(leftSum, sums);
	}
	sums = 0;
	for i:= mid +1; i <= right; i++ {
		sums += nums[i];
		rightSum = max(right, sums)
	}

	return leftSum + rightSum
}


func findMaxSubArray(nums []int, left int, right int ) int {
	if (nums == nil || left == right) {
		return nums[left];
	}
	mid := (left+right)/2
	var leftMax = findMaxSubArray(nums, left, mid);
	var rightMax = findMaxSubArray(nums, mid+1, right);
	var crossMax = maxCrossArray(nums, left, right, mid);

	return max(leftMax, rightMax, crossMax);
}


func main() {
	problem := []int{-2,1,-3,4,-1,2,1,-5,4};
	answer := findMaxSubArray(problem, 0, len(problem)-1)
	fmt.Println(answer);
}