package main

import "fmt"

var (
	conins = 300
	names  = []string{
		"Tom", "Sily",
		"Jerry", "Bob",
		"Mike", "Lily",
		"Jack", "Jim",
		"Mary", "Lucy",
	}

	distribution = make(map[string]int, len(names))
)

func dispatchCoins() (ret int) {
	count := 0
	for _, name := range names {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				count += 1
				distribution[name]++
			case 'i', 'I':
				count += 2
				distribution[name] += 2
			case 'o', 'O':
				count += 3
				distribution[name] += 3
			case 'u', 'U':
				count += 4
				distribution[name] += 4
			}
		}
	}

	for k, v := range distribution {
		fmt.Println(k, v)
	}

	ret = conins - count
	return
}

func main() {
	ret := dispatchCoins()
	fmt.Println("剩余金币", ret)
}
