package server

import (
	"context"
)

type pelanggan struct {
	writer ReadWriter
}

func NewPelanggan(writer ReadWriter) PelangganService {
	return &pelanggan{writer: writer}
}

func (m *pelanggan) AddPelangganService(ctx context.Context, pelanggan Pelanggan) error {
	//fmt.Println("pelanggan")
	err := m.writer.AddPelanggan(pelanggan)
	if err != nil {
		return err
	}

	return nil
}
func (m *pelanggan) ReadPelangganService(ctx context.Context) (Pelanggans, error) {
	men, err := m.writer.ReadPelanggan()
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *pelanggan) UpdatePelangganService(ctx context.Context, pelanggan Pelanggan) error {
	err := m.writer.UpdatePelanggan(pelanggan)
	if err != nil {
		return err
	}
	return nil
}
func (m *pelanggan) ReadByNamaPelangganService(ctx context.Context, namaPelanggan string) (Pelanggan, error) {
	men, err := m.writer.ReadByNamaPelanggan(namaPelanggan)
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *pelanggan) ReadByKeteranganService(ctx context.Context, keterangan string) (Pelanggans, error) {
	men, err := m.writer.ReadByKeterangan(keterangan)
	if err != nil {
		return men, err
	}
	return men, nil
}
