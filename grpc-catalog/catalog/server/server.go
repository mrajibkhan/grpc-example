//go:generate protoc -I grpc-example/  --go_out=plugins=grpc:grpc-example/catalog grpc-example/*.proto

package main

import (
	"context"
	"encoding/json"
	"flag"
	pb "github.com/mrajibkhan/grpc-example/grpc-catalog/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"errors"
	"fmt"
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
	log.Printf("[GetCatalogs] total catalogs = %d ", len(s.catalogs))
	for i, v := range s.catalogs {
		log.Printf("straming catalog at [%d]=%v", i, v)
		stream.Send(v)
	}
	return nil
}

func (s *server) GetCatalogByName(ctx context.Context, in *pb.SearchRequest) (*pb.Catalog, error) {
	log.Printf("[GetCatalogByName] search catalog for: %v", in)
	for _, c := range s.catalogs {
		if c.CatalogName == in.CatalogName {
			return c, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *server) GetCatalogItemByName(in *pb.SearchRequest, stream pb.CatalogService_GetCatalogItemByNameServer) error {
	log.Printf("[GetCatalogItemByName] search catalog items for: %v", in)
	var catalog = &pb.Catalog{}
    match := false
	if in.CatalogName != "" {
		catalog, _ = s.GetCatalogByName(nil, &pb.SearchRequest{CatalogName: in.CatalogName})
	}

	for _, c := range catalog.CatalogItems {
		if c.Product.Name == in.ProductName {
			log.Printf("found match for product ", in.ProductName)
			match = true;
			stream.Send(c)
		}
	}


	if !match {
		return errors.New(fmt.Sprintf("Product %s not found in catalog %s", in.ProductName, in.CatalogName))
	}

	return nil
}


// loadFeatures loads features from a JSON file.
func (s *server) loadCatalogs(filePath string) {
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
	flag.StringVar(&catalogFile, "catalogFile", "grpc-catalog/testdata/catalog.json", "catalog (json) file path")

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
