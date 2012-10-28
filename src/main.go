// +build ignore

package main

import (
	"log"
	"mud"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	mud.Initialize()

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go mud.StartConnection(c)
	}
}
