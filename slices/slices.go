package slices

import (
	"fmt"
)
var p = fmt.Println
func Slices(){
	var size int
	fmt.Printf("array size : ")
	fmt.Scanf("%d", &size)
	numbers := make([]int,size)
	p(len(numbers))
}