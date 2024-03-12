package main

import (
	"fmt"
	"test/pg"
)

func main() {
	fmt.Println("hello, world!")
	pg.Test()
	pg.TestPGX()
}
