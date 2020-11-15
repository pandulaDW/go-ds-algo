package main

func moveZeros(nums []int) {

	var numNonZeroes int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			numNonZeroes++
		}
	}

	var index, i int
	for index < len(nums) {
		if index >= numNonZeroes {
			nums[index] = 0
			index++
		}
		if i < len(nums) && nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
		i++
	}
}
