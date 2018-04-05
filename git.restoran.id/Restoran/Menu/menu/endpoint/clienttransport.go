package endpoint

import (
	"context"
	"time"

	pb "MiniProject/git.bluebird.id/Restoran/Menu/menu/grpc"
	svc "MiniProject/git.bluebird.id/Restoran/Menu/menu/server"

	util "MiniProject/git.bluebird.id/Restoran/Menu/util/grpc"
	disc "MiniProject/git.bluebird.id/Restoran/Menu/util/microservice"

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
	grpcName = "grpc.MenuService"
)

func NewGRPCMenuClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.MenuService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addMenuEp = retry
	}

	var readMenuByNamaMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadMenuByNamaMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readMenuByNamaMenuEp = retry
	}

	var readMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadMenuEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readMenuEp = retry
	}

	var updateMenuEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateMenu, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateMenuEp = retry
	}

	return MenuEndpoint{AddMenuEndpoint: addMenuEp, ReadMenuByNamaMenuEndpoint: readMenuByNamaMenuEp, ReadMenuEndpoint: readMenuEp, UpdateMenuEndpoint: updateMenuEp}, nil
}
func encodeAddMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Menu)
	return &pb.AddMenuReq{
		Namamenu: req.NamaMenu, Harga: req.Harga, Idkategorimenu: req.IDKategoriMenu, Createdby: req.CreatedBy,
	}, nil
}
func encodeReadMenuByNamaMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Menu)
	return &pb.ReadMenuByNamaMenuReq{Namamenu: req.NamaMenu}, nil
}

func encodeReadMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Menu)
	return &pb.UpdateMenurReq{
		Idmenu:         req.IDMenu,
		Namamenu:       req.NamaMenu,
		Harga:          req.Harga,
		Status:         req.Status,
		Idkategorimenu: req.IDKategoriMenu,
		UpdateBy:       req.UpdateBy,
	}, nil
}
func decodeMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadMenuByNamaMenuRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadMenuByNamaMenuResp)
	return svc.Menu{
		IDMenu:         resp.Idmenu,
		NamaMenu:       resp.Namamenu,
		Harga:          resp.Harga,
		Status:         resp.Status,
		IDKategoriMenu: resp.Idkategorimenu,
		CreatedBy:      resp.Createdby,
	}, nil
}

func decodeReadMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadMenuResp)
	var rsp svc.Menus

	for _, v := range resp.AllMenu {
		itm := svc.Menu{
			IDMenu:         v.Idmenu,
			NamaMenu:       v.Namamenu,
			Harga:          v.Harga,
			Status:         v.Status,
			IDKategoriMenu: v.Idkategorimenu,
			CreatedBy:      v.Createdby,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddMenu",
		encodeAddMenuRequest,
		decodeMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
func makeClientReadMenuByNamaMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadMenuByNamaMenu",
		encodeReadMenuByNamaMenuRequest,
		decodeReadMenuByNamaMenuRespones,
		pb.ReadMenuByNamaMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadMenuByNamaMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadMenuByNamaMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadMenuEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadMenu",
		encodeReadMenuRequest,
		decodeReadMenuResponse,
		pb.ReadMenuResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateMenu(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateMenu",
		encodeUpdateMenuRequest,
		decodeMenuResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateMenu")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateMenu",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
