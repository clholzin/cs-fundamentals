package fundamentals

func change(amount int, coins []int) int {
	memo := make(map[string]int)
	return changeComp(amount, coins, 0, memo)
}

func changeComp(amount int, coins []int, index int, memo map[string]int) int {
	if amount == 0 {
		return 1
	}
	if index >= len(coins) {
		return 0
	}
	key := fmt.Sprintf("%d-%d", amount, index)

	if ways, ok := memo[key]; ok {
		//fmt.Println(key,"**ways:",ways)
		return ways
	}
	var addingbalence int
	var ways int
	var remaining int
	for addingbalence <= amount {
		remaining = amount - addingbalence
		ways += changeComp(remaining, coins, index+1, memo)
		addingbalence += coins[index]
	}
	// fmt.Println(key,"ways:",ways)
	memo[key] = ways
	return ways
}
