package main

import "fmt"

func main() {

	var name string = "Mesut"
	var name2 = "Oezdil" /* Short declaration */
	name3 := "Mest"      /* Short declaration */

	/*
		more shorter:

		var (
			name		string 	= "Mesut"
			age 		int 	= 30
			isMarried	bool	= false
			weight		float32 = 80
		)
	*/

	var age int = 30
	var isMarried bool = false
	var weight float32 = 80

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

}
