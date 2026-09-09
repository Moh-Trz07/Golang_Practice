package main

import "fmt"

func main2() {
   var numbers = []int{1, 2, 3, 5, 10, 12, 6, 7}
   num := numbers[0]
   for _,x := range numbers[1:]{
    if x > num{
        num = x
    } 
   }
   fmt.Println(num)
}
