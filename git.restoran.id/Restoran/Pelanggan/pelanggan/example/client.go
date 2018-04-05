package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/endpoint"
	opt "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/grpc"
	util "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCPelangganClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}
	//Add Pelanggan
	//client.AddPelangganService(context.Background(), svc.Pelanggan{NamaPelanggan: "Sujono", Telepon: "08967876893", Email: "sujonoo@gmail.com", Status: 1, CreatedBy: "Sujo", Keterangan: "Masuk"})
	//namaPelanggan, _ := client.ReadByNamaPelangganService(context.Background(), "Siti Astuti")
	//fmt.Println("Pelanggan based on NamaPelanggan:", namaPelanggan)
	//menn, _ := client.ReadPelangganService(context.Background())
	//fmt.Println("all Pelanggan:", menn)
	//client.UpdatePelangganService(context.Background(), svc.Pelanggan{NamaPelanggan: "Siti Hariyati", Telepon: "083892837029", Email: "sitiastuti@gmail.com", Status: 1, UpdateBy: "Joko", IDPelanggan: 1})

	keterangan, _ := client.ReadByKeteranganService(context.Background(), "M%")
	fmt.Println("Pelanggan based on Keterangan:", keterangan)
}
