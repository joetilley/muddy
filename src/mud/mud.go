package mud

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"world"
	"worldloader"
)

type MudConnection struct {
	net.Conn
}

var gameData *world.GameWorld

func Initialize() {
	gameData = worldloader.LoadWorld()
}

func StartConnection(c net.Conn) {
	mc := MudConnection{c}
	mc.PrintLine("What is your name?")
	name := mc.ReadLine()
	mc.PrintLine("Welcome %s", name)
	curRoom := gameData.StartRoom
	mc.PrintLine(curRoom.Desc)
	mc.PrintLine("You go south")
	curRoom = curRoom.S
	mc.PrintLine(curRoom.Desc)
}

func (c MudConnection) ReadLine() string {
	r := bufio.NewReader(c)
	buf := bytes.NewBuffer(make([]byte, 50))
	for {
		data, prefix, err := r.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		buf.Write(data)
		if !prefix {
			break
		}
	}
	return buf.String()
}

func (c MudConnection) PrintLine(format string, args ...interface{}) {
	fmt.Fprintf(c, format, args...)
	newLine(c)
}

func newLine(c net.Conn) {
	var esc = []byte{27}
	c.Write(esc)
	fmt.Fprintf(c, "[1B")
	c.Write(esc)
	fmt.Fprintf(c, "[255D")
}
