package endpoint

import (
	"context"
	"fmt"

	pb "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/grpc"
	scv "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcSupplierServer struct {
	addSupplier        grpctransport.Handler
	readSupplier       grpctransport.Handler
	updateSupplier     grpctransport.Handler
	readByNamaSupplier grpctransport.Handler
	readByKeterangan   grpctransport.Handler
}

func NewGRPCSupplierServer(endpoints SupplierEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.SupplierServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcSupplierServer{
		addSupplier: grpctransport.NewServer(endpoints.AddSupplierEndpoint,
			decodeAddSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddSupplier", logger)))...),
		readSupplier: grpctransport.NewServer(endpoints.ReadSupplierEndpoint,
			decodeReadSupplierRequest,
			encodeReadSupplierResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadSupplier", logger)))...),
		updateSupplier: grpctransport.NewServer(endpoints.UpdateSupplierEndpoint,
			decodeUpdateSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateSupplier", logger)))...),
		readByNamaSupplier: grpctransport.NewServer(endpoints.ReadByNamaSupplierEndpoint,
			decodeReadByNamaSupplierRequest,
			encodeReadByNamaSupplierResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadByNamaSupplier", logger)))...),
		readByKeterangan: grpctransport.NewServer(endpoints.ReadByKeteranganEndpoint,
			decodeReadByKeteranganRequest,
			encodeReadByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadByKeterangan", logger)))...),
	}
}

func decodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddSupplierReq)
	return scv.Supplier{NamaSupplier: req.GetNamaSupplier(), Alamat: req.GetAlamat(), Telepon: req.GetTelepon(), Email: req.GetEmail(), CreatedBy: req.GetCreatedBy(), Keterangan: req.GetKeterangan()}, nil
}

func decodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadByNamaSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadByNamaSupplierReq)
	return scv.Supplier{NamaSupplier: req.NamaSupplier}, nil
}

func decodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadByKeteranganReq)
	return scv.Supplier{Keterangan: req.Keterangan}, nil
}

func decodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateSupplierReq)
	return scv.Supplier{
		IDSupplier:   req.IDSupplier,
		NamaSupplier: req.NamaSupplier,
		Alamat:       req.Alamat,
		Telepon:      req.Telepon,
		Email:        req.Email,
		Status:       req.Status,
		UpdateBy:     req.UpdateBy,
	}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadByNamaSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Supplier)
	return &pb.ReadByNamaSupplierResp{
		IDSupplier:   resp.IDSupplier,
		NamaSupplier: resp.NamaSupplier,
		Alamat:       resp.Alamat,
		Telepon:      resp.Telepon,
		Email:        resp.Email,
		Status:       resp.Status,
		CreatedBy:    resp.CreatedBy,
	}, nil
}

func encodeReadByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Suppliers)

	rsp := &pb.ReadByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadByNamaSupplierResp{
			IDSupplier:   v.IDSupplier,
			NamaSupplier: v.NamaSupplier,
			Alamat:       v.Alamat,
			Telepon:      v.Telepon,
			Email:        v.Email,
			Status:       v.Status,
			CreatedBy:    v.CreatedBy,
			Keterangan:   v.Keterangan,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func encodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Suppliers)
	fmt.Println("encode read")
	rsp := &pb.ReadSupplierResp{}

	for _, v := range resp {
		itm := &pb.ReadByNamaSupplierResp{
			IDSupplier:   v.IDSupplier,
			NamaSupplier: v.NamaSupplier,
			Alamat:       v.Alamat,
			Telepon:      v.Telepon,
			Email:        v.Email,
			Status:       v.Status,
			CreatedBy:    v.CreatedBy,
			Keterangan:   v.Keterangan,
		}
		rsp.AllSupplier = append(rsp.AllSupplier, itm)
	}
	return rsp, nil
}

func (s *grpcSupplierServer) AddSupplier(ctx oldcontext.Context, supplier *pb.AddSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addSupplier.ServeGRPC(ctx, supplier)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcSupplierServer) ReadSupplier(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadSupplierResp, error) {
	_, resp, err := s.readSupplier.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadSupplierResp), nil
}

func (s *grpcSupplierServer) UpdateSupplier(ctx oldcontext.Context, cus *pb.UpdateSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateSupplier.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcSupplierServer) ReadByNamaSupplier(ctx oldcontext.Context, namasupplier *pb.ReadByNamaSupplierReq) (*pb.ReadByNamaSupplierResp, error) {
	_, resp, err := s.readByNamaSupplier.ServeGRPC(ctx, namasupplier)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadByNamaSupplierResp), nil
}

func (s *grpcSupplierServer) ReadByKeterangan(ctx oldcontext.Context, namasupplier *pb.ReadByKeteranganReq) (*pb.ReadByKeteranganResp, error) {
	_, resp, err := s.readByKeterangan.ServeGRPC(ctx, namasupplier)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadByKeteranganResp), nil
}
