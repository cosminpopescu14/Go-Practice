// test
package main

import (
	"fmt"
)

var global string = "global var"

func main() {

	const PI float32 = 3.14159265359
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
	//fm.Println("0 / 0 = ", 0/0) //don't do this
	//with proper error handling

	//Strings and variables

	fmt.Println(len("cosmin"))
	fmt.Println("Hello World"[1]) //result -> 101  because the character is represented by a byte

	name := "cosmin" + " " + "popescu"
	fmt.Println(name)

	fmt.Println("Value of expression is: ", (true && false) || (false && true) || !(false && false))

	/*var x string
	x = "something"
	fmt.Println(x)*/

	var x string = "hello"
	var y string = "hello"
	//or short sintax for var
	//delta := "delta"
	fmt.Println(x == y)

	fmt.Println(global)

	//PI = 3.14159265359 //it is a mistake to modify a constant. it will result in a compile time error

	myfunction()

	//fmt.Fscan()
}

func myfunction() {
	fmt.Println("printed from function: ", global) //has access to global variable
}

func multipleVariables() {
	var (
		a = 15
		b = 16
		c = 17
	)
}
