package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/server"
	kit "github.com/go-kit/kit/endpoint"
)

type PelangganEndpoint struct {
	AddPelangganEndpoint        kit.Endpoint
	ReadPelangganEndpoint       kit.Endpoint
	UpdatePelangganEndpoint     kit.Endpoint
	ReadByNamaPelangganEndpoint kit.Endpoint
	ReadByKeteranganEndpoint    kit.Endpoint
}

func NewPelangganEndpoint(service svc.PelangganService) PelangganEndpoint {
	addPelangganEp := makeAddPelangganEndpoint(service)
	readPelangganEp := makeReadPelangganEndpoint(service)
	updatePelangganEp := makeUpdatePelangganEndpoint(service)
	readByNamaPelangganEp := makeReadByNamaPelangganEndpoint(service)
	readByKeteranganEp := makeReadByKeteranganEndpoint(service)
	return PelangganEndpoint{AddPelangganEndpoint: addPelangganEp,
		ReadPelangganEndpoint:       readPelangganEp,
		UpdatePelangganEndpoint:     updatePelangganEp,
		ReadByNamaPelangganEndpoint: readByNamaPelangganEp,
		ReadByKeteranganEndpoint:    readByKeteranganEp,
	}
}

func makeAddPelangganEndpoint(service svc.PelangganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Pelanggan)
		err := service.AddPelangganService(ctx, req)
		return nil, err
	}
}

func makeReadPelangganEndpoint(service svc.PelangganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadPelangganService(ctx)
		return result, err
	}
}

func makeUpdatePelangganEndpoint(service svc.PelangganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Pelanggan)
		err := service.UpdatePelangganService(ctx, req)
		return nil, err
	}
}
func makeReadByNamaPelangganEndpoint(service svc.PelangganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Pelanggan)
		result, err := service.ReadByNamaPelangganService(ctx, req.NamaPelanggan)
		return result, err
	}
}

func makeReadByKeteranganEndpoint(service svc.PelangganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Pelanggan)
		result, err := service.ReadByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
