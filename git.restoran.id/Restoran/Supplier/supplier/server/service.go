package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "supplier.bluebird.id"
	OnAdd     Status = 1
)

type Supplier struct {
	IDSupplier   int32
	NamaSupplier string
	Alamat       string
	Telepon      string
	Email        string
	Status       int32
	CreatedBy    string
	CreatedOn    string
	UpdateBy     string
	UpdateOn     string
	Keterangan   string
}
type Suppliers []Supplier

type ReadWriter interface {
	AddSupplier(Supplier) error
	ReadByNamaSupplier(string) (Supplier, error)
	ReadSupplier() (Suppliers, error)
	UpdateSupplier(Supplier) error
	ReadByKeterangan(string) (Suppliers, error)
}

type SupplierService interface {
	AddSupplierService(context.Context, Supplier) error
	ReadByNamaSupplierService(context.Context, string) (Supplier, error)
	ReadSupplierService(context.Context) (Suppliers, error)
	UpdateSupplierService(context.Context, Supplier) error
	ReadByKeteranganService(context.Context, string) (Suppliers, error)
}
