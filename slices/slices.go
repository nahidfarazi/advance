package main

import (
	"fmt"
)
var p = fmt.Println

type Person struct{
	Name string
	Age int
}

func sliceStruct(){
	slice := []*Person{}
	p1:=Person{"Nahid",22}
	p2:= Person{"Hasib",24}
	p2.Name = "kasem"
	slice =append(slice,&p1,&p2)
	for _,p:=range slice{
		fmt.Println(p.Name)
	}
	p(p2.Name)
}





func main(){
sliceStruct()
}