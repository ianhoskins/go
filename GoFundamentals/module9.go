package main

import (
	"fmt"
)

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 100
	leagueTitles["NewCastle"] = 200

	recentGame := map[string]int{
		"Sunderland": 5,
		"NewCastle":  0,
	}

	fmt.Printf("\n Title: %v \nRecent: %v\n", leagueTitles, recentGame)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is %v", key, value)
	}
	fmt.Printf("\n-----")
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is %v", key, value)
	}
	fmt.Printf("\n-----")
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is %v", key, value)
	}
	fmt.Printf("\n-----")
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is %v", key, value)
	}

	fmt.Println("\n\n", testMap["C"])

	testMap["C"] = 100
	fmt.Println("\n\n", testMap["C"])

	delete(testMap, "E")
	fmt.Println("\n\n", testMap)
}
