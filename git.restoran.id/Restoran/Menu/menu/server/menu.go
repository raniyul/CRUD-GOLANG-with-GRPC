package server

import (
	"context"
)

type menu struct {
	writer ReadWriter
}

func NewMenu(writer ReadWriter) MenuService {
	return &menu{writer: writer}
}

func (m *menu) AddMenuService(ctx context.Context, menu Menu) error {
	//fmt.Println("menu")
	err := m.writer.AddMenu(menu)
	if err != nil {
		return err
	}

	return nil
}
func (m *menu) ReadMenuService(ctx context.Context) (Menus, error) {
	men, err := m.writer.ReadMenu()
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *menu) UpdateMenuService(ctx context.Context, men Menu) error {
	err := m.writer.UpdateMenu(men)
	if err != nil {
		return err
	}
	return nil
}
func (m *menu) ReadMenuByNamaMenuService(ctx context.Context, namamenu string) (Menu, error) {
	men, err := m.writer.ReadMenuByNamaMenu(namamenu)
	if err != nil {
		return men, err
	}
	return men, nil
}
