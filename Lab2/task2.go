package main

import"fmt"

func main(){

	// a := 10
	// b :=5

	var a int  
	 var b int

	

	fmt.Println("enter two intergers")
	fmt.Scan(&a,&b)
	


	fmt.Println("Both are positive",a>=0 && b>=0)

	fmt.Println("One is greater than other",a>b || b>a)

	fmt.Println("Not equal ",a!=b)

}