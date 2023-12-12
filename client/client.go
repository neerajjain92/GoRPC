package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

func main() {

	// Address of this variable will be sent to RPC Server
	// Type of reply should be same as that specified on server
	var reply int64
	args := Args{}

	// DialHTTP connects to an HTTP RPC server at thr specified network
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	// Invoke the remote funtion, GiveMeServerTime attached to ServerTime pointer
	err = client.Call("ServerTime.GiveMeServerTime", args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	// Print the server time came in reply
	log.Printf("%d", reply)
}
