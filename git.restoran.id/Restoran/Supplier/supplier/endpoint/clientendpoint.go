package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/server"
)

func (me SupplierEndpoint) AddSupplierService(ctx context.Context, supplier sv.Supplier) error {
	_, err := me.AddSupplierEndpoint(ctx, supplier)
	return err
}

func (me SupplierEndpoint) ReadByNamaSupplierService(ctx context.Context, namasupplier string) (sv.Supplier, error) {
	req := sv.Supplier{NamaSupplier: namasupplier}
	fmt.Println(req)
	resp, err := me.ReadByNamaSupplierEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Supplier)
	return cus, err
}

func (me SupplierEndpoint) ReadSupplierService(ctx context.Context) (sv.Suppliers, error) {
	resp, err := me.ReadSupplierEndpoint(ctx, nil)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Suppliers), err
}

func (me SupplierEndpoint) UpdateSupplierService(ctx context.Context, cus sv.Supplier) error {
	_, err := me.UpdateSupplierEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (me SupplierEndpoint) ReadByKeteranganService(ctx context.Context, keterangan string) (sv.Suppliers, error) {
	req := sv.Supplier{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := me.ReadByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Suppliers)
	return cus, err
}
