package main

import "fmt"

func main(){

	var a int
	var b int 
	var choice int 

	for {
		fmt.Println("1.Add, 2.Sub, 3.Mul, 4.Div, 5.Exit")
		fmt.Print("Choose : ")
		fmt.Scan(&choice)

		if choice <1 || choice >5{
			fmt.Println("Invalid choice")
			continue	
		}else if choice == 5 {
			break;
		}
		fmt.Println("Plese Enter Your Favorite Two Number")
		fmt.Scan(&a, &b)

		if choice == 1{
			fmt.Println("Result=",a+b)
		}else if choice == 2{
			fmt.Println("Result=",a-b)
		}else if choice == 3{
			fmt.Println("Result=",a*b)
		}else if choice == 4{
				fmt.Printf("Result= %.2f\n",float32(a)/float32(b))
			}
		}
	} 

