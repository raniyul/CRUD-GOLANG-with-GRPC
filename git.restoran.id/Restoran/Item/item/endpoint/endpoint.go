package endpoint

import (
	"context"

	svc "MiniProject/git.restoran.id/Restoran/Item/item/server"
	kit "github.com/go-kit/kit/endpoint"
)

type ItemEndpoint struct {
	AddItemEndpoint              kit.Endpoint
	ReadItemByIDSupplierEndpoint kit.Endpoint
	ReadItemEndpoint             kit.Endpoint
	UpdateItemEndpoint           kit.Endpoint
	ReadItemBystatusEndpoint     kit.Endpoint
	ReadByKeteranganEndpoint     kit.Endpoint
}

func NewItemEndpoint(service svc.ItemService) ItemEndpoint {
	addItemEp := makeAddItemEndpoint(service)
	readItemByIDSupplierEp := makeReadItemByIDSupplierEndpoint(service)
	readItemEp := makeReadItemEndpoint(service)
	updateItemEp := makeUpdateItemEndpoint(service)
	readItemBystatusEp := makeReadItemBystatusEndpoint(service)
	readByKeteranganEp := makeReadByKeteranganEndpoint(service)
	return ItemEndpoint{AddItemEndpoint: addItemEp,
		ReadItemByIDSupplierEndpoint: readItemByIDSupplierEp,
		ReadItemEndpoint:             readItemEp,
		UpdateItemEndpoint:           updateItemEp,
		ReadItemBystatusEndpoint:     readItemBystatusEp,
		ReadByKeteranganEndpoint:     readByKeteranganEp,
	}
}

func makeAddItemEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Item)
		err := service.AddItemService(ctx, req)
		return nil, err
	}
}

func makeReadItemByIDSupplierEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Item)
		result, err := service.ReadItemByIDSupplierService(ctx, req.IDSupplier)
		/*return svc.Item{ItemId: result.ItemId, Name: result.Name,
		ItemType: result.ItemType, IDSupplier: result.IDSupplier, status: result.status,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadItemEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadItemService(ctx)
		return result, err
	}
}

func makeUpdateItemEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Item)
		err := service.UpdateItemService(ctx, req)
		return nil, err
	}
}

func makeReadItemBystatusEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Item)
		result, err := service.ReadItemBystatusService(ctx, req.Status)
		return result, err
	}
}

func makeReadByKeteranganEndpoint(service svc.ItemService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Item)
		result, err := service.ReadByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
