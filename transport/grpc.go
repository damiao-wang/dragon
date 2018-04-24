package transport

import (
	"context"
	"dragon/pb"
	"fmt"

	addendpoint "dragon/endpoint"

	"github.com/go-kit/kit/log"
	oldcontext "golang.org/x/net/context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	sum    grpctransport.Handler
	concat grpctransport.Handler
	hello  grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints addendpoint.Set, logger log.Logger) pb.AddServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		sum: grpctransport.NewServer(
			endpoints.SumEndpoint,
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
			options...,
		),
		concat: grpctransport.NewServer(
			endpoints.ConcatEndpoint,
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
			options...,
		),
		hello: grpctransport.NewServer(
			endpoints.HelloEndpoint,
			decodeGRPCHelloRequest,
			encodeGRPCHelloResponse,
			options...,
		),
	}
}

func (s *grpcServer) Sum(ctx oldcontext.Context, req *pb.SumReq) (*pb.SumResp, error) {
	_, rep, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	return rep.(*pb.SumResp), nil
}

func (s *grpcServer) Concat(ctx oldcontext.Context, req *pb.ConcatReq) (*pb.ConcatResp, error) {
	_, rep, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ConcatResp), nil
}

func (s *grpcServer) Hello(ctx oldcontext.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	_, rep, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResp), nil
}

func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*pb.SumReq), nil
}

// decodeGRPCConcatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC concat request to a user-domain concat request. Primarily useful in a
// server.
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*pb.ConcatReq), nil
}

func decodeGRPCHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*pb.HelloReq), nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.SumResp), nil
}

// encodeGRPCConcatResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain concat response to a gRPC concat reply. Primarily useful in a
// server.
func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ConcatResp), nil
}

func encodeGRPCHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.HelloResp), nil
}
