package main

import "fmt"

type rec struct{
	wid float64
	hei float64
}

func (r rec) Area() float64{
	return r.hei * r.wid
}

func (r rec) Perimeter() float64{
	return (r.hei + r.wid) + 2
}

func main11(){
	rect := rec{}
	var x string
	fmt.Print("give the height :")
	fmt.Scan(&rect.hei)
	fmt.Print("\ngive the width :")
	fmt.Scan(&rect.wid)
    fmt.Print("\nwhat do you want to calculate (Area or Perimeter) : ")
    fmt.Scan(&x)
	switch x{
	case "a":
		fmt.Printf("\nArea of Rectangle is: %.2f", rect.Area())
	case "p":
		fmt.Printf("\nPerimeter of Rectangle is: %.2f", rect.Perimeter())
	default:
		fmt.Print("\nInvaled")
		return
	}

}
