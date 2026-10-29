package main

import "fmt"

func main(){

	var a int
	var b int

	fmt.Println("Please Enter Your Favorites two Interges")
	fmt.Scan(&a, &b)
	
	fmt.Println("myXor",a ^ b)
	
	fmt.Println("myNOT", a, ^a , b ,^b)
	
	fmt.Println("myOR", a | b)
	
	fmt.Println("myAND",a & b)

	fmt.Println("myLeftShift",a <<2)

	fmt.Println("myRightShift",a >>2)
	

}