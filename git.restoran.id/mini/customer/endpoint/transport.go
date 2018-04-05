package endpoint

import (
	"context"

	scv "MiniProject/git.bluebird.id/mini/customer/server"

	pb "MiniProject/git.bluebird.id/mini/customer/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcCustomerServer struct {
	addCustomer          grpctransport.Handler
	readCustomerByMobile grpctransport.Handler
	readCustomer         grpctransport.Handler
	updateCustomer       grpctransport.Handler
	readCustomerByEmail  grpctransport.Handler
}

func NewGRPCCustomerServer(endpoints CustomerEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.CustomerServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcCustomerServer{
		addCustomer: grpctransport.NewServer(endpoints.AddCustomerEndpoint,
			decodeAddCustomerRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddCustomer", logger)))...),
		readCustomerByMobile: grpctransport.NewServer(endpoints.ReadCustomerByMobileEndpoint,
			decodeReadCustomerByMobileRequest,
			encodeReadCustomerByMobileResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadCustomerByMobile", logger)))...),
		readCustomer: grpctransport.NewServer(endpoints.ReadCustomerEndpoint,
			decodeReadCustomerRequest,
			encodeReadCustomerResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadCustomer", logger)))...),
		updateCustomer: grpctransport.NewServer(endpoints.UpdateCustomerEndpoint,
			decodeUpdateCustomerRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateCustomer", logger)))...),
		readCustomerByEmail: grpctransport.NewServer(endpoints.ReadCustomerByEmailEndpoint,
			decodeReadCustomerByEmailRequest,
			encodeReadCustomerByEmailResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadCustomerByEmail", logger)))...),
	}
}

func decodeAddCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddCustomerReq)
	return scv.Customer{Name: req.GetName(), CustomerType: req.GetCustomerType(),
		Mobile: req.GetMobile(), Email: req.GetEmail(), Gender: req.GetGender(),
		CallbackPhone: req.GetCallbackPhone(), Status: req.GetStatus()}, nil
}

func decodeReadCustomerByMobileRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadCustomerByMobileReq)
	return scv.Customer{Mobile: req.Mobile}, nil
}

func decodeReadCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateCustomerReq)
	return scv.Customer{CustomerId: req.CustomerId, Name: req.Name, CustomerType: req.CustomerType,
		Mobile: req.Mobile, Email: req.Email, Gender: req.Gender, CallbackPhone: req.CallbackPhone,
		Status: req.Status}, nil
}

func decodeReadCustomerByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadCustomerByEmailReq)
	return scv.Customer{Email: req.Email}, nil

}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadCustomerByMobileResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Customer)
	return &pb.ReadCustomerByMobileResp{CustomerId: resp.CustomerId, Name: resp.Name, CustomerType: resp.CustomerType,
		Mobile: resp.Mobile, Email: resp.Email, Gender: resp.Gender, CallbackPhone: resp.CallbackPhone,
		Status: resp.Status}, nil
}

func encodeReadCustomerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Customers)

	rsp := &pb.ReadCustomerResp{}

	for _, v := range resp {
		itm := &pb.ReadCustomerByMobileResp{
			CustomerId:    v.CustomerId,
			Name:          v.Name,
			CustomerType:  v.CustomerType,
			Mobile:        v.Mobile,
			Email:         v.Email,
			Gender:        v.Gender,
			CallbackPhone: v.CallbackPhone,
			Status:        v.Status,
		}
		rsp.AllCustomer = append(rsp.AllCustomer, itm)
	}
	return rsp, nil
}

func encodeReadCustomerByEmailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Customer)
	return &pb.ReadCustomerByEmailResp{CustomerId: resp.CustomerId, Name: resp.Name, CustomerType: resp.CustomerType,
		Mobile: resp.Mobile, Email: resp.Email, Gender: resp.Gender, CallbackPhone: resp.CallbackPhone,
		Status: resp.Status}, nil
}

func (s *grpcCustomerServer) AddCustomer(ctx oldcontext.Context, customer *pb.AddCustomerReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addCustomer.ServeGRPC(ctx, customer)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcCustomerServer) ReadCustomerByMobile(ctx oldcontext.Context, mobile *pb.ReadCustomerByMobileReq) (*pb.ReadCustomerByMobileResp, error) {
	_, resp, err := s.readCustomerByMobile.ServeGRPC(ctx, mobile)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadCustomerByMobileResp), nil
}

func (s *grpcCustomerServer) ReadCustomer(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadCustomerResp, error) {
	_, resp, err := s.readCustomer.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadCustomerResp), nil
}

func (s *grpcCustomerServer) UpdateCustomer(ctx oldcontext.Context, cus *pb.UpdateCustomerReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateCustomer.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcCustomerServer) ReadCustomerByEmail(ctx oldcontext.Context, email *pb.ReadCustomerByEmailReq) (*pb.ReadCustomerByEmailResp, error) {
	_, resp, err := s.readCustomerByEmail.ServeGRPC(ctx, email)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadCustomerByEmailResp), nil
}
