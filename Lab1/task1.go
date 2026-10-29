package main

import "fmt"

func main(){
	//fmt.Println("hello, world!")
	// fmt.Println("plus", a+b)
    // fmt.Println("")
	//var a, b int = 10, 3
	// var (
	// 	a int =10
	// 	b int =3
	// )

	var a int 
	var b int 
	fmt.Println("enter two intergers")
	fmt.Scan(&a,&b)

	result :=a
	fmt.Println("=",result)

	result +=b
	fmt.Println("+=",result)

	result -=b
	fmt.Println("-=",result)

	result *=b
	fmt.Println("*=",result)

	result /=b
	fmt.Println("/=",result)

	result %=b
	fmt.Println("%=",result)

}

// func enter() {
// 	var a int 
// 	var b int 
// 	fmt.Println("enter two intergers",)
// 	fmt.Scan(&a,&b)
// }