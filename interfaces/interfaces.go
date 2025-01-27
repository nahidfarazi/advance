package interfaces

import "fmt"

var p = fmt.Println
type Walker interface{
	Walk()
}
type Dog struct{
	Name string
	Age int
	Color string
}
func(d Dog) Walk(){
	p("Dog with name",d.Name,"is walking")
}
type Cat struct{
	Name string
	Age int
	Color string
}
func (c Cat) Walk(){
	p("Cat with name",c.Name,"is walking")
}
func MakeWalk(w Walker){
	w.Walk()
}
func Interfaces(){
	dog1 := Dog{Name: "Rex",Age: 2,Color: "Brown"}
	MakeWalk(dog1)
	cat1 := Cat{Name: "Wizex",Age: 3,Color: "White"}
	MakeWalk(cat1)
}