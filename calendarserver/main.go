package main

import (
	"flag"
	log "github.com/golang/glog"
)

var ip = flag.String("a", "", "listen ip")
var port = flag.String("p", "16688", "listen port")

func main() {
	flag.Parse()
	defer log.Flush()

	server := NewServer(new(Model))

	server.Run(*ip + ":" + *port)
}
