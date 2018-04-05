package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.bluebird.id/Restoran/Supplier/supplier/endpoint"
	opt "MiniProject/git.bluebird.id/Restoran/Supplier/util/grpc"
	util "MiniProject/git.bluebird.id/Restoran/Supplier/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCSupplierClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}
	//Add Supplier
	//client.AddSupplierService(context.Background(), svc.Supplier{NamaSupplier: "Siti Astuti", Alamat: "Kuningan Timur", Telepon: "083892837029", Email: "sitiastuti@gmail.com", Status: 1, CreatedBy: "Jono", Keterangan: "Libur"})
	//namaSupplier, _ := client.ReadByNamaSupplierService(context.Background(), "Siti Astuti")
	//fmt.Println("Supplier based on NamaSupplier:", namaSupplier)
	//menn, _ := client.ReadSupplierService(context.Background())
	//fmt.Println("all suppliers:", menn)

	//client.UpdateSupplierService(context.Background(), svc.Supplier{NamaSupplier: "Siti Hariyati", Alamat: "Depok", Telepon: "083892837029", Email: "sitiastuti@gmail.com", Status: 1, UpdateBy: "Joko", IDSupplier: 1})

	keterangan, _ := client.ReadByKeteranganService(context.Background(), "M%")
	fmt.Println("Supplier based on Keterangan:", keterangan)
}
