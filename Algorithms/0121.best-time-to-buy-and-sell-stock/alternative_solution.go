func maxProfit(prices []int) int {
    min := math.MaxInt64
    ans := 0
    
    for i:=0;i<len(prices);i++ {
        if prices[i] < min {
            min = prices[i]
        } else {
            ans = Max(ans, prices[i]-min)
        }
    }
    return ans
}

func Max(i, j int) int {
    if i > j {
        return i
    }
    return j
}
