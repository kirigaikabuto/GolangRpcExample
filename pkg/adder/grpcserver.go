package adder

import (
	"GolangGrpc/pkg/api"
	"context"
)

type GRPCServer struct {

}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse,error){
	return &api.AddResponse{Result: req.GetA() + req.GetB()},nil
}