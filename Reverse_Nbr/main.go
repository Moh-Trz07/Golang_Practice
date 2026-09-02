package main

import "fmt"

func main4(){
	var n int 
	fmt.Print("give a nbr: ")
	fmt.Scan(&n)
	org := n
	rev := 0
	for n > 0{
		digit := n % 10
		rev = rev*10 + digit
		n /= 10
	}
	fmt.Printf("Reverse of %d is : %d\n", org, rev)
}
