package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-with-go/cafe"
	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedCafeServer
}

type Product struct {
	Name  string
	Price int32
}

type Products struct {
	Values []*Product
}

var products = Products{Values: []*Product{
	&Product{Name: "coffee", Price: 100},
	&Product{Name: "late", Price: 110},
	&Product{Name: "mocha", Price: 120},
}}

func (pds *Products) getPrice(n string) (*int32, error) {
	var r = funk.Find(pds.Values, func(r *Product) bool {
		return r.Name == n
	})
	pd, ok := r.(*Product)
	if ok {
		return &pd.Price, nil
	} else {
		return nil, fmt.Errorf("There isn't %v in menus.", n)
	}
}

func (s *server) Order(ctx context.Context, in *pb.OrderRequest) (*pb.OrderReply, error) {
	log.Printf("Received: %v", in.GetName())
	p, err := (&products).getPrice(in.Name)
	if err != nil {
		log.Print(err)
		return &pb.OrderReply{}, nil
	} else {
		return &pb.OrderReply{Price: *p}, nil
	}

}

func (s *server) GetMenus(ctx context.Context, in *empty.Empty) (*pb.GetMenusReply, error) {
	log.Print("Which drink do you want?")
	var menus []*pb.Menu
	for _, value := range products.Values {
		menus = append(menus, &pb.Menu{Name: value.Name, Price: value.Price})
	}
	return &pb.GetMenusReply{Menus: menus}, nil
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCafeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
