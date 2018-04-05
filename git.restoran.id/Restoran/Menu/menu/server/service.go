package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "menu.bluebird.id"
	OnAdd     Status = 1
)

type Menu struct {
	IDMenu         int32
	NamaMenu       string
	Harga          int32
	IDKategoriMenu int32
	Status         int32
	CreatedBy      string
	CreatedOn      string
	UpdateBy       string
	UpdateOn       string
}
type Menus []Menu

type ReadWriter interface {
	AddMenu(Menu) error
	ReadMenuByNamaMenu(string) (Menu, error)
	ReadMenu() (Menus, error)
	UpdateMenu(Menu) error
}

type MenuService interface {
	AddMenuService(context.Context, Menu) error
	ReadMenuByNamaMenuService(context.Context, string) (Menu, error)
	ReadMenuService(context.Context) (Menus, error)
	UpdateMenuService(context.Context, Menu) error
}
