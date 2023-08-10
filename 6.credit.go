package main

import (
	"fmt"

	"github.com/akriventsev/go-card"
)

func main() {
	// Initialize a new card:
	card, err := card.NewCard("4716339239466898", "334", 12, 2023, "Ivanov Ivan")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("valid\n",card)
	}
}
