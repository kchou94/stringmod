package main

import (
	"fmt"

	qt "github.com/bernylinville/stringmod/quotes"
	str "github.com/bernylinville/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
