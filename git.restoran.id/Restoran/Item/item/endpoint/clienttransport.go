package endpoint

import (
	"context"
	"time"

	pb "MiniProject/git.restoran.id/Restoran/Item/item/grpc"
	svc "MiniProject/git.restoran.id/Restoran/Item/item/server"

	util "MiniProject/git.restoran.id/Restoran/Item/util/grpc"
	disc "MiniProject/git.restoran.id/Restoran/Item/util/microservice"

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
	grpcName = "grpc.ItemService"
)

func NewGRPCItemClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.ItemService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addItemEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddItemEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer) //retryMax menjalankan fungsi circuit braker
		addItemEp = retry
	}

	var readItemByIDSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadItemByIDSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readItemByIDSupplierEp = retry
	}

	var readItemEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadItemEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readItemEp = retry
	}

	var updateItemEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateItem, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateItemEp = retry
	}

	var readItemBystatusEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadItemBystatus, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readItemBystatusEp = retry
	}

	var readByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readByKeteranganEp = retry
	}

	return ItemEndpoint{AddItemEndpoint: addItemEp, ReadItemByIDSupplierEndpoint: readItemByIDSupplierEp,
		ReadItemEndpoint: readItemEp, UpdateItemEndpoint: updateItemEp,
		ReadItemBystatusEndpoint: readItemBystatusEp,
		ReadByKeteranganEndpoint: readByKeteranganEp}, nil
}

func encodeAddItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Item)
	return &pb.AddItemReq{
		NamaItem:   req.NamaItem,
		Merk:       req.Merk,
		IDSupplier: req.IDSupplier,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadItemByIDSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Item)
	return &pb.ReadItemByIDSupplierReq{IDSupplier: req.IDSupplier}, nil
}

func encodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Item)
	return &pb.ReadByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Item)
	return &pb.UpdateItemReq{
		IDItem:     req.IDItem,
		NamaItem:   req.NamaItem,
		Merk:       req.Merk,
		IDSupplier: req.IDSupplier,
		Status:     req.Status,
		UpdateBy:   req.UpdateBy,
		UpdateOn:   req.CreatedOn,
	}, nil
}

func encodeReadItemBystatusRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Item)
	return &pb.ReadItemBystatusReq{Status: req.Status}, nil
}

func decodeItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadItemByIDSupplierRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadItemByIDSupplierResp)
	return svc.Item{
		IDItem:     resp.IDItem,
		NamaItem:   resp.NamaItem,
		Merk:       resp.Merk,
		IDSupplier: resp.IDSupplier,
		Status:     resp.Status,
		CreatedBy:  resp.CreatedBy,
		CreatedOn:  resp.CreatedOn,
		UpdateBy:   resp.UpdateBy,
		UpdateOn:   resp.UpdateOn,
	}, nil
}

func decodeReadItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadItemResp)
	var rsp svc.Items

	for _, v := range resp.AllItem {
		itm := svc.Item{
			IDItem:     v.IDItem,
			NamaItem:   v.NamaItem,
			Merk:       v.Merk,
			IDSupplier: v.IDSupplier,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			CreatedOn:  v.CreatedOn,
			UpdateBy:   v.UpdateBy,
			UpdateOn:   v.UpdateOn,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadItemBystatusRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadItemBystatusResp)
	return svc.Item{
		IDItem:     resp.IDItem,
		NamaItem:   resp.NamaItem,
		Merk:       resp.Merk,
		IDSupplier: resp.IDSupplier,
		Status:     resp.Status,
		CreatedBy:  resp.CreatedBy,
		CreatedOn:  resp.CreatedOn,
		UpdateBy:   resp.UpdateBy,
		UpdateOn:   resp.UpdateOn,
	}, nil
}

func decodeReadByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadByKeteranganResp)
	var rsp svc.Items

	for _, v := range resp.AllKeterangan {
		itm := svc.Item{
			IDItem:     v.IDItem,
			NamaItem:   v.NamaItem,
			Merk:       v.Merk,
			IDSupplier: v.IDSupplier,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddItemEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddItem",
		encodeAddItemRequest,
		decodeItemResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddItem")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddItem",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadItemByIDSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadItemByIDSupplier",
		encodeReadItemByIDSupplierRequest,
		decodeReadItemByIDSupplierRespones,
		pb.ReadItemByIDSupplierResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadItemByIDSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadItemByIDSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadItemEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadItem",
		encodeReadItemRequest,
		decodeReadItemResponse,
		pb.ReadItemResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadItem")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadItem",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateItem(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateItem",
		encodeUpdateItemRequest,
		decodeItemResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateItem")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateItem",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadItemBystatus(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadItemBystatus",
		encodeReadItemBystatusRequest,
		decodeReadItemBystatusRespones,
		pb.ReadItemBystatusResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadItemBystatus")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadItemBystatus",
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
