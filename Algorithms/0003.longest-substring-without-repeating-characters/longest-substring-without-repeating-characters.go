package problem0003

/**
1. Initalize a array to store chars as there are max 128 chars create a fixed array of 128 ints to store the location of the char in given string
2. Initilize maxLenght, start
3. Iterate through given string and extract char
4. If char is already present in charArray and the placement of the char is greater than start index then
5. We found a duplicate char so we exclude everything before the previous positon of this char and move the start index to prev position + 1
6. Also check if the maxLength is less than the current max 
7. Insert char position in array to current position
**/
func lengthOfLongestSubstring(s string) int {
    charArray := [128]int{}
    start := 0
    maxLength := 0
    var i int
    for i := range charArray {
	    charArray[i] = -1 // as defualt value is 0, init with -1
    }
    for i = 0; i<len(s); i++ {
        if charArray[s[i]] >= start {
            //found duplicate
            if (i - start) > maxLength {
                maxLength = i - start 
            }
            start = charArray[s[i]]+1
        }
        
        charArray[s[i]] = i
    }
    if (i - start) > maxLength {
        maxLength = i - start 
    }
    return maxLength
}
