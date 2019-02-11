package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"testing"

	"github.com/abhijitherekar/gRPC/customer"
	"github.com/golang/protobuf/proto"
)

var cust = &customer.CustomerRequest{
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

func BenchmarkMarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(cust)
		if err != nil {
			log.Panicln("cannot marshall")
		}
	}
}
func BenchmarkMarshalXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(cust)
		if err != nil {
			log.Panicln("cannot marshall")
		}
	}
}
func BenchmarkMarshalproto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(cust)
		if err != nil {
			log.Panicln("cannot marshall")
		}
	}
}

func BenchmarkUnMarshalJSON(b *testing.B) {
	data, err := json.Marshal(cust)
	if err != nil {
		log.Panicln("cannot marshall")
	}
	for i := 0; i < b.N; i++ {
		var cust customer.CustomerRequest
		err := json.Unmarshal(data, &cust)
		if err != nil {
			log.Panicln("cannot UNmarshall")
		}
	}
}
func BenchmarkUnMarshalXML(b *testing.B) {
	data, err := xml.Marshal(cust)
	if err != nil {
		log.Panicln("cannot marshall")
	}
	for i := 0; i < b.N; i++ {
		var cust customer.CustomerRequest
		err := xml.Unmarshal(data, &cust)
		if err != nil {
			log.Panicln("cannot UNmarshall")
		}
	}
}
func BenchmarkUnMarshalproto(b *testing.B) {
	data, err := proto.Marshal(cust)
	if err != nil {
		log.Panicln("cannot marshall")
	}
	for i := 0; i < b.N; i++ {
		var cust customer.CustomerRequest
		err := proto.Unmarshal(data, &cust)
		if err != nil {
			log.Panicln("cannot UNmarshall")
		}
	}
}
