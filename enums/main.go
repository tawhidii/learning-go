package main

import "fmt"

/*
important note :
	there is no default enums in go ,
	but we can make by given example

*/

type RoyaltyType int

const (
	Platinum RoyaltyType = 1
	Gold     RoyaltyType = 2
	Silver   RoyaltyType = 3
)

func getDiscount(royalty RoyaltyType) int {
	fmt.Println(royalty)
	switch royalty {
	case Platinum:
		return 50
	case Gold:
		return 30
	case Silver:
		return 10
	default:
		panic("No royalty found !")
	}
}

func main() {
	fmt.Println()
}
