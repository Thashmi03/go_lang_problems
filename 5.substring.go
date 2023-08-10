package main

import "fmt"

func main() {
	//var str1 string = "go"
	var str2 string
	fmt.Scanln(&str2)
	occ := 0
	len := len(str2)
	//fmt.Println(len,str1)
	for i := 0; i < len-1; i++ {
		if (str2[i] == 'g' && str2[i+1] == 'o') || (str2[i] == 'G' && str2[i+1] == 'O') {
			occ++
		}
	}
	fmt.Println("occurence =", occ)
}
