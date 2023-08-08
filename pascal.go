package main

import "fmt"

func main() {
	var row int
	fmt.Scanln(&row)
	a := make([][]int, row)
	
	for i := range a {
		a[i] = make([]int, i+1)
	}
	fmt.Println(a)
	for j := 0; j <row; j++ {
		a[j][0] = 1//row start always 1
	}
	k := 0
	for i := 1; i < row; i++ {
		for k = 1; k < i; k++ {
			if i != 1 {
				a[i][k] = a[i-1][k-1] + a[i-1][k]//adding the before 2 val
			}
		}
		a[i][k] = 1//placing 1 at last of every row
	}
	for i := 0; i <row; i++ {

		for b := 1; b < row-i; b++ {
			fmt.Print(" ")//space
		}
		for k = 0; k <= i; k++ {
			fmt.Printf("%v ", a[i][k])
		}
		fmt.Println()
	}
}
