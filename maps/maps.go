package main

import "fmt"
var p = fmt.Println
func main(){
	//Create a map
	//var mapName map [key_type]value_type
	var maps map[string]int
	p(maps)

	//initialization
	foodsMap:=map[string]int{
		"apple":5,
		"banana":10,
		"lemon":15,
		"cherry":20,
	}
	//accessing elements
	//p(foodsMap["lemon"])
	//modifying map elements
	foodsMap["cherry"]=50
	//p(foodsMap["cherry"])
	//add new elements 
	foodsMap["orange"]=90
	//p(foodsMap)
	//delete elements
	delete(foodsMap,"apple")
	//iterate
	for k,v := range foodsMap{
		fmt.Printf("key : %s -- value : %d\n",k,v)
	}
}