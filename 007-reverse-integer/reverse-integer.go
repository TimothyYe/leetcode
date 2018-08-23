package problem007

func reverse(x int) int {
	arr := []int{}

	for {
		if x != 0 {
			arr = append(arr, x%10)
			x = x / 10
		} else {
			break
		}
	}

	var result int
	for i := 0; i < len(arr); i++ {
		t := 1
		for j := 0; j < len(arr)-1-i; j++ {
			t *= 10
		}

		result += t * arr[i]
	}

	if result < -2147483648 || result > 2147483647 {
		return 0
	}

	return result
}
