func search(nums []int, target int) int {
    if len(nums) == 0 {
         return -1
    }

    var left=0
    var right=len(nums)-1
    
    for ;left <= right;{
        var mid=left+(right-left)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            left = mid + 1
        }else if nums[mid] > target{
            right = mid - 1
        }
    }
    return -1
}