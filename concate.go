package main

import "fmt"

//two strings take to concate
func ConcatString(str1, str2 string) string {
	return str1 + str2
}

func main() {
	
	result := ConcatString("Ishu, ", "Padsala")
	fmt.Println(result) 
}
