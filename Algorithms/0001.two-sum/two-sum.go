package problem0001

/**
1. Create a map to store all numbers
2. Iterate through all numbers - num
3. check if target - num is available in map - if yes return 
4. Store number in map while iterating
5. If nothing matches return nil
**/
func twoSum(nums []int, target int) []int {
	// map to store nums while iterating
	index := make(map[int]int, len(nums))

	// Iterate over nums
	for i, b := range nums {
		// check if target - num available in map already
		if j, ok := index[target-b]; ok {
			// if available return it
			return []int{j, i}
		}

		// Add current number to map
		index[b] = i
	}

	return nil
}
