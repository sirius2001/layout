package grpc

import (
	"net"

	"google.golang.org/grpc"
)

// RPCService struct
type RPCService struct {
	address    string
	grpcServer *grpc.Server
	listenner  net.Listener
}

// ServiceAddr implements core.ServiceInner.
func (r *RPCService) ServiceAddr() string {
	return r.address
}

// ServiceName implements core.ServiceInner.
func (r *RPCService) ServiceName() string {
	return "GrpcService"
}

// StartService implements core.ServiceInner.
func (r *RPCService) StartService() error {
	//TODO:regiser your rpc service

	return r.grpcServer.Serve(r.listenner)
}

// StopService implements core.ServiceInner.
func (r *RPCService) StopService() error {
	r.grpcServer.Stop()
	return r.listenner.Close()
}

func NewRPCServer(addr string) (*RPCService, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &RPCService{
		address:    addr,
		grpcServer: grpc.NewServer(),
		listenner:  listener,
	}, nil
}
