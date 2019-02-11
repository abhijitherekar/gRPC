package main

import (
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/abhijitherekar/gRPC/customer"
	"golang.org/x/net/context"
)

const (
	addr = "127.0.0.1:40001"
)

//create wrappers for the client.CreateCustomer
func createCustomer(cc customer.CustomerClient, req *customer.CustomerRequest) error {
	resp, err := cc.CreateCustomer(context.Background(), req)
	if err != nil {
		log.Panicln("err getting the response")
		return err
	}
	if resp.Success {
		log.Println("create success")
	}
	return nil
}

//create wrappers for the client.GetCustomer
func getCustomer(cc customer.CustomerClient, c *customer.CustomerId) error {
	stream, err := cc.GetCustomer(context.Background(), c)
	if err != nil {
		log.Panicln("err getting the response")
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", cc, err)
		}
		log.Println("got the customer", resp)
	}
	return nil
}

func main() {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println("grpc dial error", err)
		return
	}
	client := customer.NewCustomerClient(cc)
	defer cc.Close()

	cust := &customer.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addr: []*customer.CustomerRequest_Address{
			&customer.CustomerRequest_Address{
				Street: "1 Mission Street",
				City:   "San Francisco",
				State:  "CA",
			},
			&customer.CustomerRequest_Address{
				Street: "Greenfield",
				City:   "Kochi",
				State:  "KL",
			},
		},
	}

	// Create a new customer
	createCustomer(client, cust)

	cust = &customer.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addr: []*customer.CustomerRequest_Address{
			&customer.CustomerRequest_Address{
				Street: "1 Mission Street",
				City:   "San Francisco",
				State:  "CA",
			},
		},
	}

	// Create a new customer
	createCustomer(client, cust)
	// Filter with an empty Keyword
	filter := &customer.CustomerId{Keyword: "Irene Rose"}
	getCustomer(client, filter)
}
