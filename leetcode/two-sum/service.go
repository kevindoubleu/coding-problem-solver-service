package twosum

func solve(nums []int, target int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
