package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main(){
	var input string

	fmt.Print("Input Your Favorite Text : ")
	fmt.Scan(&input)

	text:=[]byte(input)
	fmt.Printf("Input in binary : %b\n",text)

	hex:=hex.EncodeToString(text)
	fmt.Println("Input in Hexadecimal : ",hex)

	base64:=base64.StdEncoding.EncodeToString(text)
	fmt.Println("Input in Base64 : ",base64)


}