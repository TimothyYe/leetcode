package problem026

func removeDuplicates(nums []int) int {
    var j = 0
    for i := 1; i < len(nums) ; i++ {
        if nums[j] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }
    
    return j+1
}
