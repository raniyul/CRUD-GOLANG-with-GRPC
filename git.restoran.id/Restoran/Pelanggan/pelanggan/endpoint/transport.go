package endpoint

import (
	"context"

	pb "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/grpc"
	scv "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcPelangganServer struct {
	addPelanggan        grpctransport.Handler
	readPelanggan       grpctransport.Handler
	updatePelanggan     grpctransport.Handler
	readByNamaPelanggan grpctransport.Handler
	readByKeterangan    grpctransport.Handler
}

func NewGRPCPelangganServer(endpoints PelangganEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.PelangganServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcPelangganServer{
		addPelanggan: grpctransport.NewServer(endpoints.AddPelangganEndpoint,
			decodeAddPelangganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddPelanggan", logger)))...),
		readPelanggan: grpctransport.NewServer(endpoints.ReadPelangganEndpoint,
			decodeReadPelangganRequest,
			encodeReadPelangganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadPelanggan", logger)))...),
		updatePelanggan: grpctransport.NewServer(endpoints.UpdatePelangganEndpoint,
			decodeUpdatePelangganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdatePelanggan", logger)))...),
		readByNamaPelanggan: grpctransport.NewServer(endpoints.ReadByNamaPelangganEndpoint,
			decodeReadByNamaPelangganRequest,
			encodeReadByNamaPelangganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadByNamaPelanggan", logger)))...),
		readByKeterangan: grpctransport.NewServer(endpoints.ReadByKeteranganEndpoint,
			decodeReadByKeteranganRequest,
			encodeReadByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadByKeterangan", logger)))...),
	}
}

func decodeAddPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddPelangganReq)
	return scv.Pelanggan{NamaPelanggan: req.GetNamaPelanggan(), Telepon: req.GetTelepon(), Email: req.GetEmail(), CreatedBy: req.GetCreatedBy(), Keterangan: req.GetKeterangan()}, nil
}

func decodeReadPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadByNamaPelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadByNamaPelangganReq)
	return scv.Pelanggan{NamaPelanggan: req.NamaPelanggan}, nil
}

func decodeReadByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadByKeteranganReq)
	return scv.Pelanggan{Keterangan: req.Keterangan}, nil
}

func decodeUpdatePelangganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdatePelangganReq)
	return scv.Pelanggan{
		IDPelanggan:   req.IDPelanggan,
		NamaPelanggan: req.NamaPelanggan,
		Telepon:       req.Telepon,
		Email:         req.Email,
		Status:        req.Status,
		UpdateBy:      req.UpdateBy,
	}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadByNamaPelangganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Pelanggan)
	return &pb.ReadByNamaPelangganResp{
		IDPelanggan:   resp.IDPelanggan,
		NamaPelanggan: resp.NamaPelanggan,
		Telepon:       resp.Telepon,
		Email:         resp.Email,
		Status:        resp.Status,
		CreatedBy:     resp.CreatedBy,
	}, nil
}

func encodeReadByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Pelanggans)
	rsp := &pb.ReadByKeteranganResp{}
	for _, v := range resp {
		itm := &pb.ReadByNamaPelangganResp{
			IDPelanggan:   v.IDPelanggan,
			NamaPelanggan: v.NamaPelanggan,
			Telepon:       v.Telepon,
			Email:         v.Email,
			Status:        v.Status,
			CreatedBy:     v.CreatedBy,
			Keterangan:    v.Keterangan,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func encodeReadPelangganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Pelanggans)

	rsp := &pb.ReadPelangganResp{}

	for _, v := range resp {
		itm := &pb.ReadByNamaPelangganResp{
			IDPelanggan:   v.IDPelanggan,
			NamaPelanggan: v.NamaPelanggan,
			Telepon:       v.Telepon,
			Email:         v.Email,
			Status:        v.Status,
			CreatedBy:     v.CreatedBy,
			Keterangan:    v.Keterangan,
		}
		rsp.AllPelanggan = append(rsp.AllPelanggan, itm)
	}
	return rsp, nil
}

func (s *grpcPelangganServer) AddPelanggan(ctx oldcontext.Context, pelanggan *pb.AddPelangganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addPelanggan.ServeGRPC(ctx, pelanggan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcPelangganServer) ReadPelanggan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadPelangganResp, error) {
	_, resp, err := s.readPelanggan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadPelangganResp), nil
}

func (s *grpcPelangganServer) UpdatePelanggan(ctx oldcontext.Context, cus *pb.UpdatePelangganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updatePelanggan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcPelangganServer) ReadByNamaPelanggan(ctx oldcontext.Context, namapelanggan *pb.ReadByNamaPelangganReq) (*pb.ReadByNamaPelangganResp, error) {
	_, resp, err := s.readByNamaPelanggan.ServeGRPC(ctx, namapelanggan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadByNamaPelangganResp), nil
}

func (s *grpcPelangganServer) ReadByKeterangan(ctx oldcontext.Context, namapelanggan *pb.ReadByKeteranganReq) (*pb.ReadByKeteranganResp, error) {
	_, resp, err := s.readByKeterangan.ServeGRPC(ctx, namapelanggan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadByKeteranganResp), nil
}
