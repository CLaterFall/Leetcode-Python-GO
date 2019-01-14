func maxProfit(prices []int, fee int) int {
	cash := 0
	hold := -prices[0]
	length := len(prices)
	var i int
	for i = 1; i < length; i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
