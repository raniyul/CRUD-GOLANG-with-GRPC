package endpoint

import (
	"context"
	"time"

	pb "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/grpc"
	svc "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/server"

	util "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/grpc"
	disc "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/microservice"

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
	grpcName = "grpc.PelangganService"
)

func NewGRPCPelangganClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.PelangganService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addPelangganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddPelangganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addPelangganEp = retry
	}

	var readByNamaPelangganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadByNamaPelangganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readByNamaPelangganEp = retry
	}

	var readPelangganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadPelangganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readPelangganEp = retry
	}

	var updatePelangganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdatePelanggan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updatePelangganEp = retry
	}

	var readByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readByKeteranganEp = retry
	}

	return PelangganEndpoint{AddPelangganEndpoint: addPelangganEp, ReadByNamaPelangganEndpoint: readByNamaPelangganEp, ReadPelangganEndpoint: readPelangganEp, UpdatePelangganEndpoint: updatePelangganEp, ReadByKeteranganEndpoint: readByKeteranganEp}, nil
}
func encodeAddPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Pelanggan)
	return &pb.AddPelangganReq{
		NamaPelanggan: req.NamaPelanggan,
		Telepon:       req.Telepon,
		Email:         req.Email,
		Status:        req.Status,
		CreatedBy:     req.CreatedBy,
		Keterangan:    req.Keterangan,
	}, nil
}
func encodeReadByNamaPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Pelanggan)
	return &pb.ReadByNamaPelangganReq{NamaPelanggan: req.NamaPelanggan}, nil
}

func encodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Pelanggan)
	return &pb.ReadByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdatePelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Pelanggan)
	return &pb.UpdatePelangganReq{
		IDPelanggan:   req.IDPelanggan,
		NamaPelanggan: req.NamaPelanggan,
		Telepon:       req.Telepon,
		Email:         req.Email,
		Status:        req.Status,
		UpdateBy:      req.UpdateBy,
	}, nil
}
func decodePelangganResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadByNamaPelangganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadByNamaPelangganResp)
	return svc.Pelanggan{
		IDPelanggan:   resp.IDPelanggan,
		NamaPelanggan: resp.NamaPelanggan,
		Telepon:       resp.Telepon,
		Email:         resp.Email,
		Status:        resp.Status,
		CreatedBy:     resp.CreatedBy,
	}, nil
}

func decodeReadByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadByKeteranganResp)
	var rsp svc.Pelanggans

	for _, v := range resp.AllKeterangan {
		itm := svc.Pelanggan{
			IDPelanggan:   v.IDPelanggan,
			NamaPelanggan: v.NamaPelanggan,
			Telepon:       v.Telepon,
			Email:         v.Email,
			Status:        v.Status,
			CreatedBy:     v.CreatedBy,
			Keterangan:    v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadPelangganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadPelangganResp)
	var rsp svc.Pelanggans

	for _, v := range resp.AllPelanggan {
		itm := svc.Pelanggan{
			IDPelanggan:   v.IDPelanggan,
			NamaPelanggan: v.NamaPelanggan,
			Telepon:       v.Telepon,
			Email:         v.Email,
			Status:        v.Status,
			CreatedBy:     v.CreatedBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddPelangganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddPelanggan",
		encodeAddPelangganRequest,
		decodePelangganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddPelanggan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddPelanggan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadByNamaPelangganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadByNamaPelanggan",
		encodeReadByNamaPelangganRequest,
		decodeReadByNamaPelangganRespones,
		pb.ReadByNamaPelangganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadByNamaPelanggan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadByNamaPelanggan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadPelangganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadPelanggan",
		encodeReadPelangganRequest,
		decodeReadPelangganResponse,
		pb.ReadPelangganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadPelanggan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadPelanggan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdatePelanggan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdatePelanggan",
		encodeUpdatePelangganRequest,
		decodePelangganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdatePelanggan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdatePelanggan",
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
