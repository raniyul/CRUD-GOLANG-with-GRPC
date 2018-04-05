package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/mini/customer/server"
	kit "github.com/go-kit/kit/endpoint"
)

type CustomerEndpoint struct {
	AddCustomerEndpoint          kit.Endpoint
	ReadCustomerByMobileEndpoint kit.Endpoint
	ReadCustomerEndpoint         kit.Endpoint
	UpdateCustomerEndpoint       kit.Endpoint
	ReadCustomerByEmailEndpoint  kit.Endpoint
}

func NewCustomerEndpoint(service svc.CustomerService) CustomerEndpoint {
	addCustomerEp := makeAddCustomerEndpoint(service)
	readCustomerByMobileEp := makeReadCustomerByMobileEndpoint(service)
	readCustomerEp := makeReadCustomerEndpoint(service)
	updateCustomerEp := makeUpdateCustomerEndpoint(service)
	readCustomerByEmailEp := makeReadCustomerByEmailEndpoint(service)
	return CustomerEndpoint{AddCustomerEndpoint: addCustomerEp,
		ReadCustomerByMobileEndpoint: readCustomerByMobileEp,
		ReadCustomerEndpoint:         readCustomerEp,
		UpdateCustomerEndpoint:       updateCustomerEp,
		ReadCustomerByEmailEndpoint:  readCustomerByEmailEp,
	}
}

func makeAddCustomerEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		err := service.AddCustomerService(ctx, req)
		return nil, err
	}
}

func makeReadCustomerByMobileEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByMobileService(ctx, req.Mobile)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadCustomerEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadCustomerService(ctx)
		return result, err
	}
}

func makeUpdateCustomerEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		err := service.UpdateCustomerService(ctx, req)
		return nil, err
	}
}

func makeReadCustomerByEmailEndpoint(service svc.CustomerService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Customer)
		result, err := service.ReadCustomerByEmailService(ctx, req.Email)
		return result, err
	}
}
