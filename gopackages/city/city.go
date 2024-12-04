package city

import "gopackages/wordz"

func City() string {
	wordz.Words = []string{"Moskow", "Minsk", "Brest", "Grodno"}
	wordz.Prefix = ""
	return wordz.Random()
}
