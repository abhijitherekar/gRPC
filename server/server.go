package main

import (
	"context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"github.com/abhijitherekar/gRPC/customer"
)

const (
	port = ":40001"
)

//Server db
type Server struct {
	entires []*customer.CustomerRequest
}

//CreateCustomer creates a customer and adds to the server entires
func (s *Server) CreateCustomer(ctxt context.Context, req *customer.CustomerRequest) (*customer.CustomerResponse, error) {
	s.entires = append(s.entires, req)
	return &customer.CustomerResponse{Success: true}, nil
}

//GetCustomer searches the db and returns the entry
func (s *Server) GetCustomer(cQuery *customer.CustomerId, stream customer.Customer_GetCustomerServer) error {
	for _, cust := range s.entires {
		if cQuery.Keyword != "" {
			if !strings.Contains(cust.Name, cQuery.Keyword) {
				continue
			}
		}
		err := stream.Send(cust)
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("cannot bind", err)
	}
	srv := grpc.NewServer()
	customer.RegisterCustomerServer(srv, &Server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln("cannot start the server", err)
	}
}
