package digit

import (
	"fmt"

	"gopackages/wordz"
)

func Digit() {
	wordz.Words = []string{"One", "Two", "Three", "Four", "Five"}
	fmt.Println(wordz.Random())
}
