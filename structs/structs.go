package main

import "fmt"

type Dog struct{
	Name string
	Age int
	Color string
}
func(d Dog) Barks(){
	fmt.Printf("%s says : woof!\n",d.Name)
}
func(d *Dog) ChangeName(name string){
	d.Name = name
}
// constructor

func newDog(name, color string, age int)  Dog{
	myDog := Dog{
		Name: name,
		Color: color,
		Age: age,
	}
	return myDog
}

func main(){
	myDog := newDog("fix","red",6)
	myDog.ChangeName("Rex")
	fmt.Println(myDog)
	
}