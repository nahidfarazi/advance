package method

import "fmt"

var p = fmt.Println
type Person struct{
	Name string
	Age int
}
func(p Person) ShowData(){
	fmt.Printf("Name: %s, Age: %d\n",p.Name,p.Age)
}

type Arithmetic struct{
	X int
	Y int
	Op string
}
func(a Arithmetic) Operation(){
	if a.Op == "+"{
		p("Sum :",a.X+a.Y)
	}else if a.Op == "-"{
		p("Sub :",a.X-a.Y)
	}else if a.Op == "*"{
		p("Mul :",a.X*a.Y)
	}else if a.Op == "/"{
		if (a.X  == 0 && a.X<a.Y)||(a.Y  == 0 && a.Y>a.X) {
			p("Divided not possible. make sure a right value")
		}else{
			p("Divided :",a.X/a.Y)
		}
	}
}

