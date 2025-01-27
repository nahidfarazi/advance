package main

import (
	"fmt"
	"time"
)

func printNumber(){
	for i := 1; i <=5; i++ {
		time.Sleep(250*time.Millisecond)
		fmt.Println(i)		
	}
}
func printLetter(){
	for i := 'a'; i <='e'; i++ {
		time.Sleep(350*time.Millisecond)
		fmt.Printf("%c\n",i)
	}
}
func main() {
	
	go printNumber()
   	go printLetter()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("main func done")
	
}