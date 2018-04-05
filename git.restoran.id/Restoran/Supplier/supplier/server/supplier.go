package server

import (
	"context"
)

type supplier struct {
	writer ReadWriter
}

func NewSupplier(writer ReadWriter) SupplierService {
	return &supplier{writer: writer}
}

func (m *supplier) AddSupplierService(ctx context.Context, supplier Supplier) error {
	//fmt.Println("supplier")
	err := m.writer.AddSupplier(supplier)
	if err != nil {
		return err
	}

	return nil
}
func (m *supplier) ReadSupplierService(ctx context.Context) (Suppliers, error) {
	men, err := m.writer.ReadSupplier()
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *supplier) UpdateSupplierService(ctx context.Context, supplier Supplier) error {
	err := m.writer.UpdateSupplier(supplier)
	if err != nil {
		return err
	}
	return nil
}
func (m *supplier) ReadByNamaSupplierService(ctx context.Context, namaSupplier string) (Supplier, error) {
	men, err := m.writer.ReadByNamaSupplier(namaSupplier)
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *supplier) ReadByKeteranganService(ctx context.Context, keterangan string) (Suppliers, error) {
	men, err := m.writer.ReadByKeterangan(keterangan)
	if err != nil {
		return men, err
	}
	return men, nil
}
