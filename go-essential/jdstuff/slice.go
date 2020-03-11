package main

import "fmt"

func main() {
	fiveElementStringArray := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("type=%T value=%v\n", fiveElementStringArray, fiveElementStringArray)
	modify(fiveElementStringArray)
	fmt.Printf("type=%T value=%v\n", fiveElementStringArray, fiveElementStringArray)

}

func modify(myArray []string) {
	myArray[1] = "boo"
	fmt.Printf("type=%T value=%v\n", myArray, myArray)

}
