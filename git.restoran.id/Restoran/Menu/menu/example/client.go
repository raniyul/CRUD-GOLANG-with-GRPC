package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.bluebird.id/Restoran/Menu/menu/endpoint"
	opt "MiniProject/git.bluebird.id/Restoran/Menu/util/grpc"
	util "MiniProject/git.bluebird.id/Restoran/Menu/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCMenuClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}
	//Add Menu
	//client.AddCustomerService(context.Background(), svc.Customer{Name: "baba", CustomerType: 1, Mobile: "087889", Email: "jkl@gmail", Gender: "F", CallbackPhone: "087889"})
	//client.AddMenuService(context.Background(), svc.Menu{NamaMenu: "Soto", Harga: 10000, IDKategoriMenu: 1, CreatedBy: "Jojo"})
	//namaMenu, _ := client.ReadMenuByNamaMenuService(context.Background(), "Bakso")
	//fmt.Println("Menu based on NamaMenu:", namaMenu)
	menn, _ := client.ReadMenuService(context.Background())
	fmt.Println("all Menu:", menn)
	//client.UpdateMenuService(context.Background(), svc.Menu{NamaMenu: "Bakso", Harga: 9000, IDKategoriMenu: 5, UpdateBy: "Joko", Status: 1, IDMenu: 1})

}
