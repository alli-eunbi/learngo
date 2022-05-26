package main

import (
	"fmt"
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

func forArray(numbers ...int)(number int){
	// for _, n := range numbers{
	// 	number += n
	// }

	for i:=0; i<len(numbers); i++{
		number += numbers[i]
	}
	return 
}

func areYouRich(money int) bool{
	if koreanMoney := money*1200 ; koreanMoney> 8000 {
		return true
	}
	return false
}

func switchCases(money int)bool{
	switch koreanMoney := money*1200; {
	case koreanMoney > 80000 : 
		return true
	
	default: 
		return false
	}
}

func pointer(num int) (int, int, *int,int){
	a := num
	b := a
	c := &a
	d := *&a
	return a, b, c, d
}


func main(){
	// something.SayHello()
	// num := 1
	// fmt.Println(num)
	// fmt.Println(nameLen("alli"))
	// name, _:= nameLen("alli")
	// fmt.Println(name)
	// allNameLen("alli", "cj", "jerry", "eunbi")
	// fmt.Println(nakedReturn("alli"))
	// fmt.Println(forArray(1,2,3))
	// fmt.Println(areYouRich(6000))
	// fmt.Println(switchCases(20))
	
	fmt.Println(pointer(1))

}