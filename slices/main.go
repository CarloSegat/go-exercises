package main

import "fmt"

func main(){
    fmt.Printf("slices")


    // declare empty slice
    var emptySlice []string
    // what is it equal to? what length and capacity does it have?
    fmt.Println("uninit:", emptySlice, emptySlice == nil, len(emptySlice) == 0)


    // append some values to the empty slice
    emptySlice = append(emptySlice, "a", "b")
    // what is length and capacity of emptySlice now?
    fmt.Println(
	"capacity and length of the slice after appending and reassigning",
	cap(emptySlice),
	len(emptySlice),
    )

    // declare slice with 1,2,3
    s := []int{1,2,3}

    // instantiate slice with twice capacity as s and print it
    ss := make([]int, cap(s)*2)

    // reassign ss to a shorter version of itself and print it
    ss = ss[0:4]

    // what is the capacity of ss?
    // what is the length of ss?
    fmt.Println(cap(ss))
    fmt.Println(len(ss))
}
