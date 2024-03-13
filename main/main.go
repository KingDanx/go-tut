package main

import (
	"test/pg"
	"test/rest"
	"test/tcp"
)

func main() {
	// pg.Test()
	go pg.TestPGX()
	go tcp.Connect("192.168.1.152", "10000")
	rest.RestServer()
	select {}
}
