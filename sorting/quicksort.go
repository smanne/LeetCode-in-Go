//Quick sort
//Ref: https://www.youtube.com/watch?v=0SkOjNaO1XY&ab_channel=CSDojo
func sortArray(nums []int) []int {
    partition(nums, 0, len(nums)-1)
    return nums
}

func partition(arr []int, left, right int) {
    if left >= right {
        return 
    }
    k := left - 1
    pivotValue := arr[right]
    for i:=left; i<right; i++ {
        if arr[i] < pivotValue {
            k++
            arr[k], arr[i] = arr[i], arr[k]
        }
    }
    k++
    arr[k], arr[right] = arr[right], arr[k]
    partition(arr, left, k-1)
    partition(arr, k+1, right)
}
