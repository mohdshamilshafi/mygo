package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/mohdshamilshafi/mygo/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Shamil, Hello!"))
	fmt.Println(cmp.Diff(" Hello Shamil,", "Hello!"))
}
