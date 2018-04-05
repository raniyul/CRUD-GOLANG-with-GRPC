package endpoint

import (
	"context"

	pb "MiniProject/git.restoran.id/Restoran/Item/item/grpc"
	scv "MiniProject/git.restoran.id/Restoran/Item/item/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcItemServer struct {
	addItem              grpctransport.Handler
	readItemByIDSupplier grpctransport.Handler
	readItem             grpctransport.Handler
	updateItem           grpctransport.Handler
	readItemBystatus     grpctransport.Handler
	readByKeterangan     grpctransport.Handler
}

func NewGRPCItemServer(endpoints ItemEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.ItemServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcItemServer{
		addItem: grpctransport.NewServer(endpoints.AddItemEndpoint,
			decodeAddItemRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddItem", logger)))...),
		readItemByIDSupplier: grpctransport.NewServer(endpoints.ReadItemByIDSupplierEndpoint,
			decodeReadItemByIDSupplierRequest,
			encodeReadItemByIDSupplierResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadItemByIDSupplier", logger)))...),
		readItem: grpctransport.NewServer(endpoints.ReadItemEndpoint,
			decodeReadItemRequest,
			encodeReadItemResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadItem", logger)))...),
		updateItem: grpctransport.NewServer(endpoints.UpdateItemEndpoint,
			decodeUpdateItemRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateItem", logger)))...),
		readItemBystatus: grpctransport.NewServer(endpoints.ReadItemBystatusEndpoint,
			decodeReadItemBystatusRequest,
			encodeReadItemBystatusResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadItemBystatus", logger)))...),
		readByKeterangan: grpctransport.NewServer(endpoints.ReadByKeteranganEndpoint,
			decodeReadByKeteranganRequest,
			encodeReadByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadByKeterangan", logger)))...),
	}
}

func decodeAddItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddItemReq)
	return scv.Item{NamaItem: req.GetNamaItem(), Merk: req.GetMerk(),
		IDSupplier: req.GetIDSupplier(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn()}, nil
}

func decodeReadItemByIDSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadItemByIDSupplierReq)
	return scv.Item{IDSupplier: req.IDSupplier}, nil
}

func decodeReadItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateItemReq)
	return scv.Item{IDItem: req.IDItem, NamaItem: req.NamaItem, Merk: req.Merk,
		IDSupplier: req.IDSupplier, Status: req.Status, UpdateBy: req.UpdateBy, UpdateOn: req.UpdateOn}, nil
}

func decodeReadItemBystatusRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadItemBystatusReq)
	return scv.Item{Status: req.Status}, nil

}

func decodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadByKeteranganReq)
	return scv.Item{Keterangan: req.Keterangan}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadItemByIDSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Item)
	return &pb.ReadItemByIDSupplierResp{
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

func encodeReadItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Items)

	rsp := &pb.ReadItemResp{}

	for _, v := range resp {
		itm := &pb.ReadItemByIDSupplierResp{
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
		rsp.AllItem = append(rsp.AllItem, itm)
	}
	return rsp, nil
}

func encodeReadItemBystatusResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Item)
	return &pb.ReadItemBystatusResp{
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

func encodeReadByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Items)

	rsp := &pb.ReadByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadItemByIDSupplierResp{
			IDItem:     v.IDItem,
			NamaItem:   v.NamaItem,
			Merk:       v.Merk,
			IDSupplier: v.IDSupplier,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func (s *grpcItemServer) AddItem(ctx oldcontext.Context, item *pb.AddItemReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addItem.ServeGRPC(ctx, item)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcItemServer) ReadItemByIDSupplier(ctx oldcontext.Context, sup *pb.ReadItemByIDSupplierReq) (*pb.ReadItemByIDSupplierResp, error) {
	_, resp, err := s.readItemByIDSupplier.ServeGRPC(ctx, sup)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadItemByIDSupplierResp), nil
}

func (s *grpcItemServer) ReadItem(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadItemResp, error) {
	_, resp, err := s.readItem.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadItemResp), nil
}

func (s *grpcItemServer) UpdateItem(ctx oldcontext.Context, cus *pb.UpdateItemReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateItem.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcItemServer) ReadItemBystatus(ctx oldcontext.Context, status *pb.ReadItemBystatusReq) (*pb.ReadItemBystatusResp, error) {
	_, resp, err := s.readItemBystatus.ServeGRPC(ctx, status)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadItemBystatusResp), nil
}

func (s *grpcItemServer) ReadByKeterangan(ctx oldcontext.Context, namaitem *pb.ReadByKeteranganReq) (*pb.ReadByKeteranganResp, error) {
	_, resp, err := s.readByKeterangan.ServeGRPC(ctx, namaitem)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadByKeteranganResp), nil
}
