package main

import "fmt"

func sum(a,b int, ch chan int){
	ch <- a+b
}

func main() {
	ch:= make(chan int)
	go sum(2,3,ch)
	res := <-ch
	fmt.Println(res)
}