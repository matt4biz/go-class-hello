package main

import (
	"fmt"
	"os"

	"hello"
)

func main() {
	// if there is only one argument, then this
	// slice expression will yield an empty slice

	fmt.Println(hello.Say(os.Args[1:]))
}
