package main

import (
    "fmt"
)

func main(){
    // declare a strign2string map with 2 entries
    m := map[string]string{"ciao":"bella", "come":"stai"}
    fmt.Println(m)

    fmt.Println("printing leys and values")

    // iterate keys and values
    for v, k := range(m){
	fmt.Println(k)
	fmt.Println(v)
	for _, char := range(k) {
	    fmt.Println("the unicode of the char is", char)
	    }
    }
}
