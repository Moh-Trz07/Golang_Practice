package main

import "fmt"

func main() {
    var x int
    fmt.Print("Give a number: ")
    fmt.Scan(&x)
    
    if x < 0 {
    fmt.Println("Cant accept negative numbers")
    return
    }
    
    result := 1
    for i := 2; i <= x; i++{
    result *= i
    }
    
    fmt.Printf("Factorial of %d is: %d\n", x, result)
}
