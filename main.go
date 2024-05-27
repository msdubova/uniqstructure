package main

import (
	"fmt"
	"uniqstruct/uniqstruct"
)

func main() {
	var s = []uniqstruct.Structure{
		{ID: 3},
		{ID: 2},
		{ID: 1},
		{ID: 2},
	}
	fmt.Println(uniqstruct.Uniqstruct(s))

}
