package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "pelanggan.bluebird.id"
	OnAdd     Status = 1
)

type Pelanggan struct {
	IDPelanggan   int32
	NamaPelanggan string
	Telepon       string
	Email         string
	Status        int32
	CreatedBy     string
	CreatedOn     string
	UpdateBy      string
	UpdateOn      string
	Keterangan    string
}
type Pelanggans []Pelanggan

type ReadWriter interface {
	AddPelanggan(Pelanggan) error
	ReadByNamaPelanggan(string) (Pelanggan, error)
	ReadPelanggan() (Pelanggans, error)
	UpdatePelanggan(Pelanggan) error
	ReadByKeterangan(string) (Pelanggans, error)
}

type PelangganService interface {
	AddPelangganService(context.Context, Pelanggan) error
	ReadByNamaPelangganService(context.Context, string) (Pelanggan, error)
	ReadPelangganService(context.Context) (Pelanggans, error)
	UpdatePelangganService(context.Context, Pelanggan) error
	ReadByKeteranganService(context.Context, string) (Pelanggans, error)
}
