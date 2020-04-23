package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	courseInProg := []string{"Docker Deep Dive", "Docker Clustering", "Kubernetes"}

	for _, i := range courseInProg {
		fmt.Println(i)
	}

	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
