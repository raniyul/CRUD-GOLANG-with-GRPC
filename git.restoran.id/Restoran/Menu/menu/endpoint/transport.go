package endpoint

import (
	"context"

	pb "MiniProject/git.bluebird.id/Restoran/Menu/menu/grpc"
	scv "MiniProject/git.bluebird.id/Restoran/Menu/menu/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcMenuServer struct {
	addMenu            grpctransport.Handler
	readMenu           grpctransport.Handler
	updateMenu         grpctransport.Handler
	readMenuByNamaMenu grpctransport.Handler
}

func NewGRPCMenuServer(endpoints MenuEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.MenuServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcMenuServer{
		addMenu: grpctransport.NewServer(endpoints.AddMenuEndpoint,
			decodeAddMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddMenu", logger)))...),
		readMenu: grpctransport.NewServer(endpoints.ReadMenuEndpoint,
			decodeReadMenuRequest,
			encodeReadMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadMenu", logger)))...),
		updateMenu: grpctransport.NewServer(endpoints.UpdateMenuEndpoint,
			decodeUpdateMenuRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateMenu", logger)))...),
		readMenuByNamaMenu: grpctransport.NewServer(endpoints.ReadMenuByNamaMenuEndpoint,
			decodeReadMenuByNamaMenuRequest,
			encodeReadMenuByNamaMenuResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadMenuByNamaMenu", logger)))...),
	}
}

func decodeAddMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddMenuReq)
	return scv.Menu{NamaMenu: req.GetNamamenu(), Harga: req.GetHarga(), IDKategoriMenu: req.GetIdkategorimenu(), CreatedBy: req.GetCreatedby()}, nil
}

func decodeReadMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadMenuByNamaMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadMenuByNamaMenuReq)
	return scv.Menu{NamaMenu: req.Namamenu}, nil
}

func decodeUpdateMenuRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateMenurReq)
	return scv.Menu{IDMenu: req.Idmenu, NamaMenu: req.Namamenu, Harga: req.Harga, IDKategoriMenu: req.Idkategorimenu, Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadMenuByNamaMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Menu)
	return &pb.ReadMenuByNamaMenuResp{Idmenu: resp.IDMenu, Namamenu: resp.NamaMenu, Harga: resp.Harga, Status: resp.Status, Idkategorimenu: resp.IDKategoriMenu, Createdby: resp.CreatedBy}, nil
}
func encodeReadMenuResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Menus)

	rsp := &pb.ReadMenuResp{}

	for _, v := range resp {
		itm := &pb.ReadMenuByNamaMenuResp{
			Idmenu:         v.IDMenu,
			Namamenu:       v.NamaMenu,
			Harga:          v.Harga,
			Status:         v.Status,
			Idkategorimenu: v.IDKategoriMenu,
			Createdby:      v.CreatedBy,
		}
		rsp.AllMenu = append(rsp.AllMenu, itm)
	}
	return rsp, nil
}

func (s *grpcMenuServer) AddMenu(ctx oldcontext.Context, menu *pb.AddMenuReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addMenu.ServeGRPC(ctx, menu)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcMenuServer) ReadMenu(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadMenuResp, error) {
	_, resp, err := s.readMenu.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadMenuResp), nil
}

func (s *grpcMenuServer) UpdateMenu(ctx oldcontext.Context, cus *pb.UpdateMenurReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateMenu.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcMenuServer) ReadMenuByNamaMenu(ctx oldcontext.Context, namamenu *pb.ReadMenuByNamaMenuReq) (*pb.ReadMenuByNamaMenuResp, error) {
	_, resp, err := s.readMenuByNamaMenu.ServeGRPC(ctx, namamenu)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadMenuByNamaMenuResp), nil
}
