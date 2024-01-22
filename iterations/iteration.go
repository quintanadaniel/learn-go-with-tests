package iterations

import "fmt"

const repeatedCont = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCont; i++ {
		repeated += character
	}
	return repeated
}

func ListNames(names []string) {
	for _, name := range names {
		fmt.Printf("name: %v", name)
	}
}