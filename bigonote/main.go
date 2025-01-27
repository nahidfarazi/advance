package main

import (
	"fmt"
	"time"
)
func sumFor(n int)int{
	sum := 0
	for i := 1; i <= n; i++{
		sum += i
	}
	return sum
}

func sum(n int)int{
	res:= n * (n+1) / 2
	return res
}
func main() {
	t1 := time.Now()
	fmt.Println(sumFor(20000))
	t2 := time.Now()
	
	// Calculate the duration
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)
}