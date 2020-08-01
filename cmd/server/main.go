package main

import (
	"GolangGrpc/pkg/adder"
	"GolangGrpc/pkg/api"
    "google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

