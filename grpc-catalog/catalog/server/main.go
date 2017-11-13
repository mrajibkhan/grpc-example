//go:generate protoc -I grpc-example/  --go_out=plugins=grpc:grpc-example/catalog grpc-example/*.proto

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	pb "github.com/mrajibkhan/grpc-example/grpc-catalog/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	catalogs []*pb.Catalog
}

var c = &pb.Catalog{}

// GetCatalog implements CatalogServiceServer.GetCatalog
func (s *server) GetCatalogs(in *pb.Empty, stream pb.CatalogService_GetCatalogsServer) error {
	for i, v := range s.catalogs {
		log.Printf("array value at [%d]=%v", i, v)
		stream.Send(v)
	}
	return nil
}

func (s *server) GetCatalogByName(ctx context.Context, in *pb.SearchRequest) (*pb.Catalog, error) {
	return c, nil
}

func (s *server) GetCatalogProductByName(ctx context.Context, in *pb.SearchRequest) (*pb.Product, error) {
	return c.CatalogItems[0].Product, nil
}

func loadCatalog() []pb.Catalog {
	var catalogFile string
	flag.StringVar(&catalogFile, "catalogFile", "catalog.json", "catalog (yaml) file path")

	flag.Parse()

	fmt.Printf("catalogFile: " + catalogFile)

	catalogs := pb.GetCatalogFromJsonFile(catalogFile)

	return catalogs
}

// loadFeatures loads features from a JSON file.
func (s *server) loadCatalogs(filePath string) {
	//catalogFile := flag.String("catalogFile", "catalog.json", "catalog (yaml) file path")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default catalogs: %v", err)
	}
	if err := json.Unmarshal(file, &s.catalogs); err != nil {
		log.Fatalf("Failed to load default catalogs: %v", err)
	}
}

func newServer(filePath string) *server {
	s := new(server)
	s.loadCatalogs(filePath)
	return s
}

func main() {

	var catalogFile string
	flag.StringVar(&catalogFile, "catalogFile", "catalog.json", "catalog (yaml) file path")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCatalogServiceServer(s, newServer(catalogFile))
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
