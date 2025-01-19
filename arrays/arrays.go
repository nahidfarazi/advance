package arrays

import (
	"fmt"
)
var p = fmt.Println

func Arrays(){
	// array declaration --> method 1
	/* var keyword --> array name [size] type */
	var numbers [4] int
	// add element in array
	/* array name [index number] = value */
	numbers [1] = 10
	p(numbers)

/* <-----------------------------------------------------------------------------------------------------------------------------> */

	// array declaration --> method 2
	/* array name := [size] type {elements} */
	foods := [4]string{"apple","mango","banana","lemon"}
	p(foods)
	
}