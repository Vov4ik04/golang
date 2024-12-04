package main

import (
	"fmt"

	"gopackages/city"
	"gopackages/digit"
	"gopackages/wordz"

	utils "github.com/Vov4ik04/myutils"
	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println("Privet")

	fmt.Println(xstrings.Shuffle(city.City()))
	digit.Digit()

	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice have string")
	}
}
