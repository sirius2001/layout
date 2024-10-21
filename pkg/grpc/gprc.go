package grpc

import (
	"context"
	"loon/pkg/grpc/pb"
	"loon/pkg/kaf"
	"loon/pkg/log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RPCService struct
type RPCService struct {
	address    string
	grpcServer *grpc.Server
	listenner  net.Listener
	pb.UnimplementedAuditServiceServer
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
	pb.RegisterAuditServiceServer(r.grpcServer, r)
	reflection.Register(r.grpcServer)
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

// Notify 处理通知请求
func (r *RPCService) Upload(ctx context.Context, req *pb.AuditRecord) (*pb.AuditReply, error) {
	log.Info("recivce a new record", "type", "grpc")
	go kaf.Message(req)
	return &pb.AuditReply{Status: http.StatusOK}, nil
}
