// studentName -- John Doe, grade -- 77, isPassed -- true
// type it with 3 differents types and see outputs


package main

import "fmt"

func main() {

/*	var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true */ 

/*	var studentName = "John Doe"
	var grade = 77
	var isPassed = true */ 

/*	studentName := "John Doe"
	grade := 77
	isPassed := true */ 

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}


var studentName, grade, isPassed = "John Doe", 77, true

studentName, grade, isPassed := "John Doe", 77, true
