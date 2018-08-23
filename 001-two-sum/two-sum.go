package problem001

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		result = result[:0]
		rest := target - nums[i]
		result = append(result, i)

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == rest {
				result = append(result, j)
				return result
			}
		}
	}

	return result
}
