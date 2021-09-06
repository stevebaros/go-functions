package main

import (
	"fmt"
	"github.com/stevebaros/go-functions/extras"
)

func main() {
a,_,_,_ := getAllMathOperations(4,8)
fmt.Println(a)/*
fmt.Println(s)
fmt.Println(d)
fmt.Println(m)*/

val := extras.NotModifying(a)

fmt.Println(val)
fmt.Println(a)
fmt.Println(extras.Modifying(&a))
fmt.Println(a)


}

func getAllMathOperations(x ,y int) (addition, subtraction, division, multiplication int) {
	return x + y, y - x, y / x, x * y

}
