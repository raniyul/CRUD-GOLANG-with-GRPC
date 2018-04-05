package server

import (
	"context"
)

type item struct {
	writer ReadWriter
}

func NewItem(writer ReadWriter) ItemService {
	return &item{writer: writer}
}

//Methode pada interface itemService di service.go
func (c *item) AddItemService(ctx context.Context, item Item) error {
	//fmt.Println("item")
	err := c.writer.AddItem(item)
	if err != nil {
		return err
	}

	return nil
}

func (c *item) ReadItemByIDSupplierService(ctx context.Context, sup int32) (Item, error) {
	cus, err := c.writer.ReadItemByIDSupplier(sup)
	//fmt.Println(cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *item) ReadItemService(ctx context.Context) (Items, error) {
	cus, err := c.writer.ReadItem()
	//fmt.Println("item", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *item) UpdateItemService(ctx context.Context, cus Item) error {
	err := c.writer.UpdateItem(cus)
	if err != nil {
		return err
	}
	return nil
}

func (c *item) ReadItemBystatusService(ctx context.Context, status int32) (Item, error) {
	cus, err := c.writer.ReadItemBystatus(status)
	//fmt.Println("item:", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (m *item) ReadByKeteranganService(ctx context.Context, keterangan string) (Items, error) {
	men, err := m.writer.ReadByKeterangan(keterangan)
	if err != nil {
		return men, err
	}
	return men, nil
}
