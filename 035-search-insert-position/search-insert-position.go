package problem035

func searchInsert(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
        
        if nums[i] < target && i+1 <= len(nums)-1 && target < nums[i + 1] {
            return i + 1
        }
        
        if nums[i] < target && i == len(nums) -1 {
            return i + 1
        }
    }
    
    return 0
}
