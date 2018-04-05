package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/server"
	kit "github.com/go-kit/kit/endpoint"
)

type SupplierEndpoint struct {
	AddSupplierEndpoint        kit.Endpoint
	ReadSupplierEndpoint       kit.Endpoint
	UpdateSupplierEndpoint     kit.Endpoint
	ReadByNamaSupplierEndpoint kit.Endpoint
	ReadByKeteranganEndpoint   kit.Endpoint
}

func NewSupplierEndpoint(service svc.SupplierService) SupplierEndpoint {
	addSupplierEp := makeAddSupplierEndpoint(service)
	readSupplierEp := makeReadSupplierEndpoint(service)
	updateSupplierEp := makeUpdateSupplierEndpoint(service)
	readByNamaSupplierEp := makeReadByNamaSupplierEndpoint(service)
	readByKeteranganEp := makeReadByKeteranganEndpoint(service)
	return SupplierEndpoint{AddSupplierEndpoint: addSupplierEp,
		ReadSupplierEndpoint:       readSupplierEp,
		UpdateSupplierEndpoint:     updateSupplierEp,
		ReadByNamaSupplierEndpoint: readByNamaSupplierEp,
		ReadByKeteranganEndpoint:   readByKeteranganEp,
	}
}

func makeAddSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.AddSupplierService(ctx, req)
		return nil, err
	}
}

func makeReadSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadSupplierService(ctx)
		return result, err
	}
}

func makeUpdateSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.UpdateSupplierService(ctx, req)
		return nil, err
	}
}
func makeReadByNamaSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		result, err := service.ReadByNamaSupplierService(ctx, req.NamaSupplier)
		return result, err
	}
}

func makeReadByKeteranganEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		result, err := service.ReadByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
