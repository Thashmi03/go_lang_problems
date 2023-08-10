package main

import (
	"fmt"
)

func main() {
	var str1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	len := len(str1)
	array := []rune(str1)
	for i := 2; i < len-1; i += 4 {
		temp := array[i]
		//fmt.Println("temp =", temp)
		array[i] = array[i+1]
		array[i+1] = temp

	}
	fmt.Println(string(array))

}
