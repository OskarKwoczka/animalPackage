package main

import (
	"fmt"

	"github.com/OskarKwoczka/puppy"
)

func main() {
	var s1 string
	var s2 string
	s1 = puppy.Bark()
	s2 = puppy.BigBark()
	// s1 := puppy.Bark()
	// s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
}
