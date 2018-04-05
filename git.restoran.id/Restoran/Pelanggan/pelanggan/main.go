package main

import (
	"fmt"

	ep "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/endpoint"
	pb "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/grpc"
	svc "MiniProject/git.bluebird.id/Restoran/Pelanggan/pelanggan/server"
	run "MiniProject/git.bluebird.id/mini/util/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cfg "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/config"
	util "MiniProject/git.bluebird.id/Restoran/Pelanggan/util/microservice"
)

func main() {
	//import dari utility
	//logginh

	logger := util.Logger()

	ok := cfg.AppConfig.LoadConfig()
	if !ok {
		logger.Log(util.LogError, "failed to load configuration")
		return
	}

	discHost := cfg.GetA("discoveryhost", "127.0.0.1:2181")
	ip := cfg.Get("serviceip", "127.0.0.1")
	port := cfg.Get("serviceport", "7001")
	address := fmt.Sprintf("%s:%s", ip, port)

	registrar, err := util.ServiceRegistry(discHost, svc.ServiceID, address, logger)
	if err != nil {
		logger.Log(util.LogError, "cannot find registrar")
		return
	}
	registrar.Register()
	defer registrar.Deregister()

	tracerHost := cfg.Get("tracerhost", "127.0.0.1:9999")
	tracer := util.Tracer(tracerHost)

	var server pb.PelangganServiceServer
	{
		//chHost := cfg.Get("chhost", "127.0.0.1:6379")
		//cacher := svc.NewRedisCache(chHost)

		//gmapKey := cfg.Get("gmapkey", "AIzaSyD9tm3UVfxRWeaOy_MQ7tsCj1fVCLfG8Bo")
		//locator := svc.NewLocator(gmapKey)

		dbHost := cfg.Get(cfg.DBhost, "127.0.0.1:3306")
		dbName := cfg.Get(cfg.DBname, "Restoran")
		dbUser := cfg.Get(cfg.DBuid, "root")
		dbPwd := cfg.Get(cfg.DBpwd, "root")

		dbReadWriter := svc.NewDBReadWriter(dbHost, dbName, dbUser, dbPwd)
		service := svc.NewPelanggan(dbReadWriter)
		endpoint := ep.NewPelangganEndpoint(service)
		fmt.Println(endpoint)
		server = ep.NewGRPCPelangganServer(endpoint, tracer, logger)
	}
	grpcServer := grpc.NewServer(run.Recovery(logger)...)
	pb.RegisterPelangganServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	exit := make(chan bool, 1)
	go run.Serve(address, grpcServer, exit, logger)

	<-exit
}
