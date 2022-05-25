package main

import (
	"fmt"
	"main/something"
)

func nameLen(name string) (string, int){
	return name, len(name)
}

func allNameLen(name ...string){
	fmt.Println(name)
}

func nakedReturn(name string) (nameLen int){
	defer fmt.Println("done")
	nameLen = len(name)
	return
}

func main(){
	something.SayHello()
	num := 1
	fmt.Println(num)
	fmt.Println(nameLen("alli"))
	name, _:= nameLen("alli")
	fmt.Println(name)
	allNameLen("alli", "cj", "jerry", "eunbi")

	fmt.Println(nakedReturn("alli"))
}