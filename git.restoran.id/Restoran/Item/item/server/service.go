package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "item.Restoran.id"
	OnAdd     Status = 1
)

type Item struct {
	IDItem     int32
	NamaItem   string
	Merk       string
	IDSupplier int32
	CreatedBy  string
	CreatedOn  string
	UpdateBy   string
	UpdateOn   string
	Status     int32
	Keterangan string
}
type Items []Item

type ReadWriter interface {
	AddItem(Item) error
	ReadItemByIDSupplier(int32) (Item, error)
	ReadItem() (Items, error)
	UpdateItem(Item) error
	ReadItemBystatus(int32) (Item, error)
	ReadByKeterangan(string) (Items, error)
}

type ItemService interface {
	AddItemService(context.Context, Item) error
	ReadItemByIDSupplierService(context.Context, int32) (Item, error)
	ReadItemService(context.Context) (Items, error)
	UpdateItemService(context.Context, Item) error
	ReadItemBystatusService(context.Context, int32) (Item, error)
	ReadByKeteranganService(context.Context, string) (Items, error)
}
