package main

import "fmt"

func main(){
    fmt.Printf("slices")

    // declare empty slice
    var []string emptySlice

    // what is it equal to? what length and capacity does it have?
    if emptySlice == nil {
	fmt.Println(len(emptySlice))
	fmt.Println(cap(emptySlice))
    }

    // append some values to the empty slice
    emptySlice = append(emptySlice, "a", "b")

    // what is length and capacity of emptySlice now?
    fmt.Println(cap(emptySlice), len(emptySlice))

    // declare slice with 1,2,3

    var s := []int{1,2,3}

    // instantiate slice with twice capacity as s and print it

    ss := make([]int, cap(s) * 2)
    fmt.Println(ss)

    // reassign ss to a shorter version of itself and print it

    ss = ss[1:3]

    // what is the capacity of ss?
    // what is the length of ss?
}
