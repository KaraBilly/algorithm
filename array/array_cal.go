package array

//Leetcode #26
func removeDuplicates(nums []int) int {
	var length = len(nums)
	var point = 0
	for i := 1; i < length; i++ {
		if nums[i] != nums[point] {
			point++
			nums[point] = nums[i]
		}
	}
	return point + 1
}
