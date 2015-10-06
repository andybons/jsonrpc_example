package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	if err := rpc.Register(new(Calculator)); err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	log.Println("Listening on localhost:1234")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		log.Printf("new connection established\n")
		go jsonrpc.ServeConn(conn)
	}
}
