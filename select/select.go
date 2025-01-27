package main

import (
	"fmt"
	"os"
	"time"
)

func serverX(ch1 chan string){
	time.Sleep(4 * time.Second)
	ch1 <- "server 1"
}
func serverY(ch2 chan string){
	time.Sleep(2 * time.Second)
	ch2 <- "server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go serverX(output1)
	go serverY(output2)

	select{
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)
	}
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}	
	defer file.Close()
}
