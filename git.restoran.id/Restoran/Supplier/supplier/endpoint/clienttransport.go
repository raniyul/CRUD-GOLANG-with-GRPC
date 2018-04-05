package endpoint

import (
	"context"
	"time"

	pb "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/grpc"
	svc "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/server"

	util "MiniProject/git.bluebird.id/Restoran/Supplier/util/grpc"
	disc "MiniProject/git.bluebird.id/Restoran/Supplier/util/microservice"

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
	grpcName = "grpc.SupplierService"
)

func NewGRPCSupplierClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.SupplierService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addSupplierEp = retry
	}

	var readByNamaSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadByNamaSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readByNamaSupplierEp = retry
	}

	var readSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readSupplierEp = retry
	}

	var updateSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateSupplier, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateSupplierEp = retry
	}

	var readByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readByKeteranganEp = retry
	}

	return SupplierEndpoint{AddSupplierEndpoint: addSupplierEp, ReadByNamaSupplierEndpoint: readByNamaSupplierEp, ReadSupplierEndpoint: readSupplierEp, UpdateSupplierEndpoint: updateSupplierEp, ReadByKeteranganEndpoint: readByKeteranganEp}, nil
}
func encodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.AddSupplierReq{
		NamaSupplier: req.NamaSupplier,
		Alamat:       req.Alamat,
		Telepon:      req.Telepon,
		Email:        req.Email,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		Keterangan:   req.Keterangan,
	}, nil
}
func encodeReadByNamaSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.ReadByNamaSupplierReq{NamaSupplier: req.NamaSupplier}, nil
}

func encodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.ReadByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.UpdateSupplierReq{
		IDSupplier:   req.IDSupplier,
		NamaSupplier: req.NamaSupplier,
		Alamat:       req.Alamat,
		Telepon:      req.Telepon,
		Email:        req.Email,
		Status:       req.Status,
		UpdateBy:     req.UpdateBy,
	}, nil
}
func decodeSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadByNamaSupplierRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadByNamaSupplierResp)
	return svc.Supplier{
		IDSupplier:   resp.IDSupplier,
		NamaSupplier: resp.NamaSupplier,
		Alamat:       resp.Alamat,
		Telepon:      resp.Telepon,
		Email:        resp.Email,
		Status:       resp.Status,
		CreatedBy:    resp.CreatedBy,
	}, nil
}

func decodeReadByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadByKeteranganResp)
	var rsp svc.Suppliers

	for _, v := range resp.AllKeterangan {
		itm := svc.Supplier{
			IDSupplier:   v.IDSupplier,
			NamaSupplier: v.NamaSupplier,
			Alamat:       v.Alamat,
			Telepon:      v.Telepon,
			Email:        v.Email,
			Status:       v.Status,
			CreatedBy:    v.CreatedBy,
			Keterangan:   v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadSupplierResp)
	var rsp svc.Suppliers

	for _, v := range resp.AllSupplier {
		itm := svc.Supplier{
			IDSupplier:   v.IDSupplier,
			NamaSupplier: v.NamaSupplier,
			Alamat:       v.Alamat,
			Telepon:      v.Telepon,
			Email:        v.Email,
			Status:       v.Status,
			CreatedBy:    v.CreatedBy,
			Keterangan:   v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddSupplier",
		encodeAddSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadByNamaSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadByNamaSupplier",
		encodeReadByNamaSupplierRequest,
		decodeReadByNamaSupplierRespones,
		pb.ReadByNamaSupplierResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadByNamaSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadByNamaSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadSupplier",
		encodeReadSupplierRequest,
		decodeReadSupplierResponse,
		pb.ReadSupplierResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateSupplier(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateSupplier",
		encodeUpdateSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadByKeterangan",
		encodeReadByKeteranganRequest,
		decodeReadByKeteranganRespones,
		pb.ReadByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
