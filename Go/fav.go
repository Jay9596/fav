/*
	Inspired by @Maktm (Michael "Maktm" A.)
	This is his code just ported to GO.

	By: Jay9596
*/

package main

import (
	"fmt"
	"os"
)

func usage(er string) {
	fmt.Printf("fav: %s\n", er)
	os.Exit(1)
}

func combine() {
	for a := range aToZ() {
		for b := range aToZ() {
			for c := range aToZ() {
				fmt.Printf("%s%s%s\t", a, b, c)
			}
		}
	}
}

func main() {
	var recursive int
	if len(os.Args) == 2 {
		if os.Args[1] == "-r" {
			recursive = 1
		} else {
			usage("fav [-r (optional)]")
		}
	}

recurse:
	combine()
	if recursive == 1 {
		goto recurse
	}
	fmt.Println("")
}

// A range generator to make a channel of string from "a" to "z"
func aToZ() (ch chan string) {
	ch = make(chan string)

	go func() {
		a := rune('a')
		z := rune('z')
		for i := a; i <= z; i++ {
			ch <- string(i)
		}
		close(ch)
	}()
	return
}
