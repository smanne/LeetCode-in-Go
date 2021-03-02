

func findSubstring(s string, words []string) []int {
    wordMap := map[string]int{}
    for _, v := range words {
        wordMap[v]++   
    }
    wordLen := len(words[0])
    totalWords := len(words)
    totalWordLen := wordLen*totalWords
    ans  := []int{}
    for i:=0; i<=len(s)-totalWordLen;i++ {
        substr := s[i:i+totalWordLen]
        tempMap := map[string]int{}
        isValid := true
        for k:=0; k<=len(substr)-wordLen; k+=wordLen {
            ov, ok := wordMap[substr[k:k+wordLen]]
            tempMap[substr[k:k+wordLen]]++
            v, _ := tempMap[substr[k:k+wordLen]]
            if !ok || v > ov {
                isValid = false
                break;
            }
        }
        if isValid {
            ans = append(ans, i)
        }
    }
    return ans
}
