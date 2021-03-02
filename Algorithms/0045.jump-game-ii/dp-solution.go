package problem0045

func jump(nums []int) int {
    pathArr := make([]int, len(nums)) 
    for i:=1; i<len(nums); i++ {
        pathArr[i] = len(nums)+1
    }
    for i:=1;i<len(nums);i++ {
        for j:=0; j<i; j++ {
            if j+nums[j] >= i {
                pathArr[i] = min(pathArr[i], pathArr[j]+1)
            }
        }
    }
    return pathArr[len(nums)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
