package main

import (
	"fmt"
)

func main() {

	myCourses := []string{"Dogs", "Cats", "Fish"}

	fmt.Printf("Length of array is: %d \nCapacity is: %d\n",
		len(myCourses), cap(myCourses))

	mySlice := make([]int, 1, 4)
	fmt.Printf("\nLength is: %d \nCapacity is: %d\n",
		len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{10, 20, 30, 40, 50}
	slice3 := append(slice1, slice2...)
	fmt.Println("\n\n", slice3)
}
