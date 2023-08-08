package main

import ("github.com/akriventsev/go-card"
"fmt")

func main(){
	// Initialize a new card:
card,err := card.NewCard("4716339239466898", "334", 12, 2023, "Ivanov Ivan")

if err!= nil {
    fmt.Print(err)
}else{
	fmt.Println(card)
}

}