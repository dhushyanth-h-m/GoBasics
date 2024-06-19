package main

import "fmt"

func main() {
	var myName string = "Dhushyanth"
	myAge := 25
	myFloat := 3.14

	fmt.Printf("Hello, %s!\n", myName)
	fmt.Printf("Age, %d\n", myAge)
	fmt.Printf("Float, %f\n", myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	fmt.Printf("Friends Name, %s my bool %t and my other int %d\n", myFriendsName, myBool, myOtherInt)
}
