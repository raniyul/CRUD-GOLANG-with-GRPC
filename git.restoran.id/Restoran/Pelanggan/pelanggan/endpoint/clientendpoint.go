package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/server"
)

func (me PelangganEndpoint) AddPelangganService(ctx context.Context, pelanggan sv.Pelanggan) error {
	_, err := me.AddPelangganEndpoint(ctx, pelanggan)
	return err
}

func (me PelangganEndpoint) ReadByNamaPelangganService(ctx context.Context, namapelanggan string) (sv.Pelanggan, error) {
	req := sv.Pelanggan{NamaPelanggan: namapelanggan}
	fmt.Println(req)
	resp, err := me.ReadByNamaPelangganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Pelanggan)
	return cus, err
}

func (me PelangganEndpoint) ReadPelangganService(ctx context.Context) (sv.Pelanggans, error) {
	resp, err := me.ReadPelangganEndpoint(ctx, nil)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Pelanggans), err
}

func (me PelangganEndpoint) UpdatePelangganService(ctx context.Context, cus sv.Pelanggan) error {
	_, err := me.UpdatePelangganEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (me PelangganEndpoint) ReadByKeteranganService(ctx context.Context, keterangan string) (sv.Pelanggans, error) {
	req := sv.Pelanggan{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := me.ReadByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Pelanggans)
	return cus, err
}
