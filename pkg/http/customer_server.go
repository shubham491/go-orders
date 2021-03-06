package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
)

type CustomerServer struct {
}

const customerTableName = "Customers"

func (s *CustomerServer) GetCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Customer, error) {
	customer, err := db.GetCustomer(customerTableName, id.CustomerId)
	return customerToPb(customer), err
}

func (s *CustomerServer) PostCustomer(ctx stdctx.Context, customer *pb.CreateCustomer) (*pb.Customer, error) {
	newCustomer, err := db.InsertCustomer(customerTableName, pbToCreateCustomer(customer))
	return customerToPb(newCustomer), err
}

func (s *CustomerServer) PutCustomer(ctx stdctx.Context, customer *pb.UpdateCustomer) (*pb.Customer, error) {
	newCustomer, err := db.UpdateCustomer(customerTableName, pbToUpdateCustomer(customer))
	return customerToPb(newCustomer), err

}

func (s *CustomerServer) DeleteCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Empty, error) {
	err := db.DeleteKey(customerTableName, id.CustomerId)
	return &pb.Empty{}, err
}

func (s *CustomerServer) ListCustomers(ctx context.Context, empty *pb.Empty) (*pb.CustomerList, error) {
	customerList, err := db.GetAllCustomers(customerTableName)
	return customerListToPb(customerList), err
}
