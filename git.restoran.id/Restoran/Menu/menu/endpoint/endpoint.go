package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/Restoran/Menu/menu/server"
	kit "github.com/go-kit/kit/endpoint"
)

type MenuEndpoint struct {
	AddMenuEndpoint            kit.Endpoint
	ReadMenuEndpoint           kit.Endpoint
	UpdateMenuEndpoint         kit.Endpoint
	ReadMenuByNamaMenuEndpoint kit.Endpoint
}

func NewMenuEndpoint(service svc.MenuService) MenuEndpoint {
	addMenuEp := makeAddMenuEndpoint(service)
	readMenuEp := makeReadMenuEndpoint(service)
	updateMenuEp := makeUpdateMenuEndpoint(service)
	readMenuByNamaMenuEp := makeReadMenuByNamaMenuEndpoint(service)
	return MenuEndpoint{AddMenuEndpoint: addMenuEp,
		ReadMenuEndpoint:           readMenuEp,
		UpdateMenuEndpoint:         updateMenuEp,
		ReadMenuByNamaMenuEndpoint: readMenuByNamaMenuEp,
	}
}

func makeAddMenuEndpoint(service svc.MenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Menu)
		err := service.AddMenuService(ctx, req)
		return nil, err
	}
}

func makeReadMenuEndpoint(service svc.MenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadMenuService(ctx)
		return result, err
	}
}

func makeUpdateMenuEndpoint(service svc.MenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Menu)
		err := service.UpdateMenuService(ctx, req)
		return nil, err
	}
}
func makeReadMenuByNamaMenuEndpoint(service svc.MenuService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Menu)
		result, err := service.ReadMenuByNamaMenuService(ctx, req.NamaMenu)
		return result, err
	}
}
