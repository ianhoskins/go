package main

import (
	"fmt"
	"os"
)

func main() {
	firstRank := 39
	secondRank := 614

	if firstRank < secondRank {
		fmt.Println("\nfirstRank is less than secondRank")
	} else if firstRank > secondRank {
		fmt.Println("\nfirstRank is greater than secondRank")
	} else {
		fmt.Println("\nNot Greater Than or Less Than")
	}

	if thirdRank, fourthRank := 40, 50; thirdRank < fourthRank {
		fmt.Println("\nthirdRank is less than fourthRank")
	} else if thirdRank > fourthRank {
		fmt.Println("\nthirdRank is greater than fourthRank")
	} else {
		fmt.Println("\nNot Greater Than or Less Than")
	}

	switch "docker" {
	case "linux":
		fmt.Println("\nlinux")
	case "docker":
		fmt.Println("\ndocker")
	case "windows":
		fmt.Println("\nwindows")
	}

	_, err := os.Open("c:\\temp\\ian.txt")
	if err != nil {
		fmt.Println("\nError was:", err)
	}
}
