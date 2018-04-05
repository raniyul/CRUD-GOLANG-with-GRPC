package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "customer.bluebird.id"
	OnAdd     Status = 1
)

type Customer struct {
	CustomerId    int64
	Name          string
	CustomerType  int32
	Mobile        string
	Email         string
	Gender        string
	CallbackPhone string
	Status        int32
}
type Customers []Customer

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddCustomer(Customer) error
	ReadCustomerByMobile(string) (Customer, error)
	ReadCustomer() (Customers, error)
	UpdateCustomer(Customer) error
	ReadCustomerByEmail(string) (Customer, error)
}

type CustomerService interface {
	AddCustomerService(context.Context, Customer) error
	ReadCustomerByMobileService(context.Context, string) (Customer, error)
	ReadCustomerService(context.Context) (Customers, error)
	UpdateCustomerService(context.Context, Customer) error
	ReadCustomerByEmailService(context.Context, string) (Customer, error)
}
