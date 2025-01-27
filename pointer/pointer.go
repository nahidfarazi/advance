package main

import "fmt"

func modifyArray(num *[5]int){
	fmt.Println(&num[4])
 (*num)[4]=20
}
func arrays(){
	arr :=&[5]int{1,2,3,4,5}
	modifyArray(arr)
	fmt.Println(&arr[4])
	fmt.Println(arr)
}
func slice(){
	s:=[]int{1,2,3,4,5}
	fmt.Println(s,cap(s),&s[0])
	s = append(s, 20)
	fmt.Println(s,cap(s),&s[0])
	s =append(s, 30)
	fmt.Println(s,cap(s),&s[0])
	s = append(s, 40)
	s = append(s, 50)
	s = append(s, 60)
	s = append(s, 70)
	fmt.Println(s,cap(s),&s[0])

}

func main() {
	//arrays()
	slice()
	

	
	
	
}