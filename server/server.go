package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type ServerTime int64

func main() {
	serverTime := new(ServerTime)

	// Register the serverTime Object upon which the GiveMeServerTime
	// function will be called on the RPC Server(from the client)
	rpc.Register(serverTime)

	// Registers an HTTP Handler for RPC messages
	// sets up an HTTP handler to handle Remote Procedure Calls (RPC)
	// over HTTP. It allows you to expose RPC services via HTTP, making them accessible to clients over the web.
	rpc.HandleHTTP()

	// Start listening on the PORT
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("Listener error: ", err)
	} else {
		log.Println("Listening on port 1234.........")
	}

	// Accepts incoming traffic on the listener
	http.Serve(listener, nil)

}

func (t *ServerTime) GiveMeServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}
