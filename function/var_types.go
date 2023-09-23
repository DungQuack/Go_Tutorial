package main

import (
	"fmt"

)

var num1, num2 int = 7,8


func main(){
	var c, python, java , run  = "Welcome" , "tutorial !", "go" ,true
	// := used in place of a var declaration with implicit type.
	d := 3
	html, js := true, false

	fmt.Println(num1 , num2 , c, java, python, run, d, html ,js)
}