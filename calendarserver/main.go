package main

import (
	"flag"
)

var ip = flag.String("a", "", "listen ip")
var port = flag.String("p", "16688", "listen port")

func main() {
	flag.Parse()

	server := NewServer(new(Model))

	server.Run(*ip + ":" + *port)
}
