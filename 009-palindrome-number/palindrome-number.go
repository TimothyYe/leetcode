package problem009

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}

	var reversed int
	tmpX := x
	for {
		tmp := tmpX % 10
		reversed = reversed*10 + tmp
		tmpX = tmpX / 10

		if tmpX == 0 {
			break
		}
	}

	if reversed == x {
		return true
	}

	return false
}
