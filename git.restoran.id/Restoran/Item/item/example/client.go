package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.restoran.id/Restoran/Item/item/endpoint"
	opt "MiniProject/git.restoran.id/Restoran/Item/util/grpc"
	util "MiniProject/git.restoran.id/Restoran/Item/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCItemClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Item
	//client.AddItemService(context.Background(), svc.Item{NamaItem: "Garam", Merk: "Segitiga", IDSupplier: 22, Status: 0, CreatedBy: "Adam"})

	//Get Item By IDSupplier No
	//cusIDSupplier, _ := client.ReadItemByIDSupplierService(context.Background(), 12)
	//fmt.Println("item based on IDSupplier:", cusIDSupplier)

	//List Item
	//cuss, _ := client.ReadItemService(context.Background())
	//fmt.Println("all items:", cuss)

	//Update Item
	//client.UpdateItemService(context.Background(), svc.Item{NamaItem: "Gula", Merk: "Gula-Gula", IDSupplier: 10, Status: 1, UpdateBy: "Yusnita", IDItem: 1})

	//Get Item By status
	//cusstatus, _ := client.ReadItemBystatusService(context.Background(), 1)
	//fmt.Println("item based on status:", cusstatus)
	keterangan, _ := client.ReadByKeteranganService(context.Background(), "M%")
	fmt.Println("Supplier based on Keterangan:", keterangan)
}
