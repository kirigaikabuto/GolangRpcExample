package main

import (
	"GolangGrpc/pkg/api"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main(){
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments")
	}
	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	conn,err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewAdderClient(conn)
	resp, err := c.Add(context.Background(), &api.AddRequest{A:int32(x), B:int32(y)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.GetResult())
}
