package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.bluebird.id/mini/customer/server"

	pb "MiniProject/git.bluebird.id/mini/customer/grpc"

	util "MiniProject/git.bluebird.id/mini/util/grpc"
	disc "MiniProject/git.bluebird.id/mini/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.CustomerService"
)

func NewGRPCCustomerClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.CustomerService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addCustomerEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddCustomerEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer) //retryMax menjalankan fungsi circuit braker
		addCustomerEp = retry
	}

	var readCustomerByMobileEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadCustomerByMobileEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readCustomerByMobileEp = retry
	}

	var readCustomerEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadCustomerEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readCustomerEp = retry
	}

	var updateCustomerEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateCustomer, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateCustomerEp = retry
	}

	var readCustomerByEmailEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadCustomerByEmail, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readCustomerByEmailEp = retry
	}
	return CustomerEndpoint{AddCustomerEndpoint: addCustomerEp, ReadCustomerByMobileEndpoint: readCustomerByMobileEp,
		ReadCustomerEndpoint: readCustomerEp, UpdateCustomerEndpoint: updateCustomerEp,
		ReadCustomerByEmailEndpoint: readCustomerByEmailEp}, nil
}

func encodeAddCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.AddCustomerReq{
		Name:          req.Name,
		CustomerType:  req.CustomerType,
		Mobile:        req.Mobile,
		Email:         req.Email,
		Gender:        req.Gender,
		CallbackPhone: req.CallbackPhone,
		Status:        req.Status,
	}, nil
}

func encodeReadCustomerByMobileRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.ReadCustomerByMobileReq{Mobile: req.Mobile}, nil
}

func encodeReadCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.UpdateCustomerReq{
		CustomerId:    req.CustomerId,
		Name:          req.Name,
		CustomerType:  req.CustomerType,
		Mobile:        req.Mobile,
		Email:         req.Email,
		Gender:        req.Gender,
		CallbackPhone: req.CallbackPhone,
		Status:        req.Status,
	}, nil
}

func encodeReadCustomerByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Customer)
	return &pb.ReadCustomerByEmailReq{Email: req.Email}, nil
}

func decodeCustomerResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadCustomerByMobileRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadCustomerByMobileResp)
	return svc.Customer{
		CustomerId:    resp.CustomerId,
		Name:          resp.Name,
		CustomerType:  resp.CustomerType,
		Mobile:        resp.Mobile,
		Email:         resp.Email,
		Gender:        resp.Gender,
		CallbackPhone: resp.CallbackPhone,
		Status:        resp.Status,
	}, nil
}

func decodeReadCustomerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadCustomerResp)
	var rsp svc.Customers

	for _, v := range resp.AllCustomer {
		itm := svc.Customer{
			CustomerId:    v.CustomerId,
			Name:          v.Name,
			CustomerType:  v.CustomerType,
			Mobile:        v.Mobile,
			Email:         v.Email,
			Gender:        v.Gender,
			CallbackPhone: v.CallbackPhone,
			Status:        v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadCustomerByEmailRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadCustomerByEmailResp)
	return svc.Customer{
		CustomerId:    resp.CustomerId,
		Name:          resp.Name,
		CustomerType:  resp.CustomerType,
		Mobile:        resp.Mobile,
		Email:         resp.Email,
		Gender:        resp.Gender,
		CallbackPhone: resp.CallbackPhone,
		Status:        resp.Status,
	}, nil
}

func makeClientAddCustomerEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddCustomer",
		encodeAddCustomerRequest,
		decodeCustomerResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddCustomer")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddCustomer",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadCustomerByMobileEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadCustomerByMobile",
		encodeReadCustomerByMobileRequest,
		decodeReadCustomerByMobileRespones,
		pb.ReadCustomerByMobileResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadCustomerByMobile")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadCustomerByMobile",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadCustomerEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadCustomer",
		encodeReadCustomerRequest,
		decodeReadCustomerResponse,
		pb.ReadCustomerResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadCustomer")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadCustomer",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateCustomer(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateCustomer",
		encodeUpdateCustomerRequest,
		decodeCustomerResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateCustomer")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateCustomer",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadCustomerByEmail(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadCustomerByEmail",
		encodeReadCustomerByEmailRequest,
		decodeReadCustomerByEmailRespones,
		pb.ReadCustomerByEmailResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadCustomerByEmail")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadCustomerByEmail",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
