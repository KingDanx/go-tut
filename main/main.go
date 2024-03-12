package main

import (
	"test/pg"
	"test/tcp"
)

func main() {
	// pg.Test()
	pg.TestPGX()
	go tcp.Connect("200.200.200.166", "10000")
	select {}
}
