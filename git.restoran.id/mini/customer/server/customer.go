package server

import (
	"context"
)

type customer struct {
	writer ReadWriter
}

// FUNGSI DIBAWAH BERDIRI SENDIRI / LOOSE COUPLING (TANPA MENUNGGU LAGI KAPAN SELESAINYA SUATU PROSES)
func NewCustomer(writer ReadWriter) CustomerService {
	return &customer{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (c *customer) AddCustomerService(ctx context.Context, customer Customer) error {
	//fmt.Println("customer")
	err := c.writer.AddCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customer) ReadCustomerByMobileService(ctx context.Context, mob string) (Customer, error) {
	cus, err := c.writer.ReadCustomerByMobile(mob)
	//fmt.Println(cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *customer) ReadCustomerService(ctx context.Context) (Customers, error) {
	cus, err := c.writer.ReadCustomer()
	//fmt.Println("customer", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *customer) UpdateCustomerService(ctx context.Context, cus Customer) error {
	err := c.writer.UpdateCustomer(cus)
	if err != nil {
		return err
	}
	return nil
}

func (c *customer) ReadCustomerByEmailService(ctx context.Context, email string) (Customer, error) {
	cus, err := c.writer.ReadCustomerByEmail(email)
	//fmt.Println("customer:", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}
