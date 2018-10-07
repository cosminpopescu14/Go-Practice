// test
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var age int //variable declaration. implicit value is 0
	age = 25
	fmt.Println("my age is", age)
	age = 26
	fmt.Println("my age is", age)

	var ageImplicit int = 27 // with initial value

	var ageTypeInference = 54

	fmt.Println("my implicit age is", ageImplicit)

	fmt.Println(ageTypeInference)

	//Multiple variable declaration
	var width, height int

	fmt.Println("the w and h is", width, height)

	fmt.Println("1 + 1 = ", 1+1)
}
