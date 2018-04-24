package service

import (
	"context"
	"dragon/pb"
)

// StringService provides operations on strings.
type AddService interface {
	// Sums two integers.
	Sum(context.Context, *pb.SumReq) (*pb.SumResp, error)
	// Concatenates two strings
	Concat(context.Context, *pb.ConcatReq) (*pb.ConcatResp, error)
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
