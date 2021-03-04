package fundamentals

import "fmt"

/* example scope:
amount: 5
coins:  [1,2,5]

// call stack

money:5 coin:5 index:2 ways: 1
money:3 coin:5 index:2 ways: 0
money:1 coin:5 index:2 ways: 0
money:5 coin:2 index:1 ways: 1
money:4 coin:5 index:2 ways: 0
money:2 coin:5 index:2 ways: 0
money:4 coin:2 index:1 ways: 1
money:3 coin:2 index:1 ways: 0
money:2 coin:2 index:1 ways: 1
money:1 coin:2 index:1 ways: 0
money:5 coin:1 index:0 ways: 4


*/

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
