package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/twistedogic/spero/pb"
	"google.golang.org/grpc"
)

type PollServer struct {
	matchStream []chan *pb.Odd
}

func NewPollServer(ch chan *pb.Odd) *PollServer {
	channels := make([]chan *pb.Odd, 0)
	server := &PollServer{matchStream: channels}
	go server.Sub(ch)
	return server
}

func (p *PollServer) Sub(out chan *pb.Odd) {
	for {
		select {
		case v := <-out:
			for _, ch := range p.matchStream {
				ch <- v
			}
		}
	}
}

//TODO: channel blocked here
func (p *PollServer) Poll(in *pb.Empty, stream pb.OddService_PollServer) error {
	newCh := make(chan *pb.Odd)
	p.matchStream = append(p.matchStream, newCh)
	for odd := range newCh {
		if err := stream.Send(odd); err != nil {
			return err
		}
	}
	return nil
}

type RpcServer struct {
	instance *grpc.Server
	listener net.Listener
}

func NewRpcServer(ch chan *pb.Odd, grpcMetrics *grpc_prometheus.ServerMetrics, port int) (*RpcServer, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	pb.RegisterOddServiceServer(grpcServer, NewPollServer(ch))
	grpc_prometheus.Register(grpcServer)
	return &RpcServer{
		instance: grpcServer,
		listener: lis,
	}, nil
}

func (r *RpcServer) Start() {
	log.Println("Starting RpcServer")
	if err := r.instance.Serve(r.listener); err != nil {
		log.Fatal(err)
	}
}

func (r *RpcServer) Stop() error {
	log.Println("Stoping RpcServer")
	r.instance.Stop()
	return nil
}
