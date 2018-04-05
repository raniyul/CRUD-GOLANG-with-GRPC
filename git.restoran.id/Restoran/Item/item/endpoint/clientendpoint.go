package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.restoran.id/Restoran/Item/item/server"
)

func (ce ItemEndpoint) AddItemService(ctx context.Context, item sv.Item) error {
	_, err := ce.AddItemEndpoint(ctx, item)
	return err
}

func (ce ItemEndpoint) ReadItemByIDSupplierService(ctx context.Context, sup int32) (sv.Item, error) {
	req := sv.Item{IDSupplier: sup}
	fmt.Println(req)
	resp, err := ce.ReadItemByIDSupplierEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Item)
	return cus, err
}

func (ce ItemEndpoint) ReadItemService(ctx context.Context) (sv.Items, error) {
	resp, err := ce.ReadItemEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Items), err
}

func (ce ItemEndpoint) UpdateItemService(ctx context.Context, cus sv.Item) error {
	_, err := ce.UpdateItemEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (ce ItemEndpoint) ReadItemBystatusService(ctx context.Context, status int32) (sv.Item, error) {
	req := sv.Item{Status: status}
	fmt.Println(req)
	resp, err := ce.ReadItemBystatusEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Item)
	return cus, err
}

func (me ItemEndpoint) ReadByKeteranganService(ctx context.Context, keterangan string) (sv.Items, error) {
	req := sv.Item{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := me.ReadByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Items)
	return cus, err
}
