package main

import "fmt"


func myClosure() func() int{
    i := 0
    return func() int {
	i++
	return i
    }
}


func main(){
    theFUnc := myClosure()
    i := theFUnc()
    fmt.Printf("i is %d\n", i)
    fmt.Printf("i is %d", theFUnc())
}
