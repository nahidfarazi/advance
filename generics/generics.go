package main

import "fmt"

func Generics[V any](nums []V){
	fmt.Println(nums)
}

func main() {
	num1:= []int{1,2,3}
	src := []string{"golang","java","c++"}
	Generics(num1)
	Generics(src)
}