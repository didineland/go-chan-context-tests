package openmeteo

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/didineland/meteo/protofiles/meteo-streaming"
)

type server struct {
	pb.UnimplementedMeteoStremingServiceServer
}

func (s server) GetCurrentStreaming(req *pb.Empty, srv pb.MeteoStremingService_GetCurrentStreamingServer) error {
	log.Println("Fetch data streaming")

	receiverChan := RegisterListerner()

	for {
		select {
		case current := <-receiverChan:

			resp := pb.CurrentResponse{
				Temperature2M: float32(current.Temperature2M),
			}

			if err := srv.Send(&resp); err != nil {
				DeregisterListener(receiverChan)
				log.Println("error generating response")
				return err
			}
		}
	}

	return nil
}

func ConnectGrpcServer(ctx context.Context) {
	// create listener
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic("error building server: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()
	pb.RegisterMeteoStremingServiceServer(s, server{})

	log.Println("start server")

	select {
	case <-ctx.Done():
		fmt.Println("grpc server cancelled")
		s.Stop()
		return
	default:
		if err := s.Serve(listener); err != nil {
			panic("error building server: " + err.Error())
		}
	}

	// if err := s.Serve(listener); err != nil {
	// 	panic("error building server: " + err.Error())
	// }

}
