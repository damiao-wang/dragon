package main

import (
	"net"
	"net/http"
	"ninja/base/misc/errors"
	"os"

	apiendpoint "dragon/endpoint"
	"dragon/pb"
	"dragon/service"
	"dragon/transport"

	"github.com/go-kit/kit/log"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	svc := &service.Service{}
	endpoints := apiendpoint.New(svc, logger)
	httpHandler := transport.NewHTTPHandler(endpoints, logger)
	grpcServer := transport.NewGRPCServer(endpoints, logger)

	ln, err := net.Listen("tcp", ":1400")
	if err != nil {
		errors.Trace(err)
	}

	m := cmux.New(ln)
	// start grpc
	{
		baseServer := grpc.NewServer()
		pb.RegisterAddServer(baseServer, grpcServer)
		match := cmux.HTTP2HeaderField("content-type", "application/grpc")
		ln := m.Match(match)
		go baseServer.Serve(ln)

	}

	// start webapi
	{
		ln := m.Match(cmux.Any())
		go http.Serve(ln, httpHandler)
	}

	logger.Log("Listen", "1400", "service", "start")
	logger.Log("err", m.Serve())
}
