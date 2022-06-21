package main

import "fmt"

func main() {
	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while
	counter := 0
	for counter < 10 {
		println(counter)
		counter++
	}

	//for forever
	//counterForever :=0

	/* 	for{
		fmt.Println(counterForever)
		counterForever++
	} */

	for j := 0; j > -100; j-- {
		fmt.Println(j)
	}
}
