package main

import (
	"log"
	"flag"
	pb "github.com/mrajibkhan/grpc-example/grpc-catalog/catalog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCatalogServiceClient(conn)

	stream, err := client.GetCatalogs(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("[client:] could not retrieve catalog from server: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got Catalogs %s ", in.CatalogName)
		}
	}()

	stream.CloseSend()
	<-waitc
}
