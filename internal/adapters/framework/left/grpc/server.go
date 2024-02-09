package grpc

import (
	log "github.com/sirupsen/logrus"
	googlegrpc "google.golang.org/grpc"
	"hexarch/internal/adapters/framework/left/grpc/pb"
	"hexarch/internal/ports"
	"hexarch/utils"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("listen on port 9000 failure: %v", err)
	}

	grpcServer := googlegrpc.NewServer()
	arithmeticServiceServer := grpca
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("serve gRPC server failure: %v", err)
	}
}
