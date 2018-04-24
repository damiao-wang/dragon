package service

import (
	"context"
	"dragon/pb"
	"fmt"
)

// StringService provides operations on strings.
type AddService interface {
	// Sums two integers.
	Sum(context.Context, *pb.SumReq) (*pb.SumResp, error)
	// Concatenates two strings
	Concat(context.Context, *pb.ConcatReq) (*pb.ConcatResp, error)
	Hello(context.Context, *pb.HelloReq) (*pb.HelloResp, error)
}

// stringService is a concrete implementation of StringService
type Service struct{}

func (s *Service) Sum(ctx context.Context, req *pb.SumReq) (*pb.SumResp, error) {
	return &pb.SumResp{
		V: req.A + req.B,
	}, nil
}

func (s *Service) Concat(ctx context.Context, req *pb.ConcatReq) (*pb.ConcatResp, error) {
	return &pb.ConcatResp{
		V: req.A + req.B,
	}, nil
}

func (s *Service) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{
		V: fmt.Sprintf("Hello %v!", req.Name),
	}, nil
}
