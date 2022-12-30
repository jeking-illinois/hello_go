package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/jeking-illinois/hello_go/morestrings"
)

func main() {

	monologue := `We've stepped into a war with the Cabal, 
		          so lets get to taking out their leadership one by one.`

	fmt.Println("Eyes up, Guardian.")
	fmt.Println(morestrings.ReverseRunes(monologue))
	fmt.Println(cmp.Diff("Little Light", "Widdle Wight"))
}
