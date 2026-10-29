package main

import (
	"encoding/hex"
	"fmt"
)

func XorEncrypt(text string , key byte) string {

	data := []byte(text)

	for i:=0; i<len(data); i++{

		data[i] = data[i] ^ key
	}
		return string (data)
}
func  main()  {

	var text string
	fmt.Print("Enter your message : ")
	fmt.Scan(&text)

	var keyInput int
	
	fmt.Print("Enter a key (0-255) : ")
	fmt.Scan(&keyInput)
	
	if(keyInput <0 || keyInput >255){
		fmt.Println("The key is invalid",keyInput)
	}

	key := byte(keyInput)
	

	encrypted := XorEncrypt(text, key)
	fmt.Println("Encrypted (hex) : ", hex.EncodeToString([]byte(encrypted)))

	decrypted := XorEncrypt(encrypted, key)
	fmt.Println("Decrypted : ", decrypted)


}