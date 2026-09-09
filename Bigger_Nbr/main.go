package main

import "fmt"

func main() {
    var count int
    fmt.Print("Enter how many numbers: ")
    fmt.Scan(&count)
    
    numbers := make([]int, count)
    
    fmt.Printf("Enter %d numbers separated by spaces: ", count)
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i])
    }
    
    num := numbers[0]
    for _, x := range numbers[1:] {
        if x > num {
            num = x
        }
    }
    fmt.Println("Largest number:", num)
}
