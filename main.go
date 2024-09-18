package main

import (
	"database/sql"
	"log"
	"main/api"
	db "main/db/sqlc"
	"main/gapi"
	"main/pb"
	"main/util"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var err error
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	runGrpcServer(config, store)
	//runGinServer(config,store)
}

func runGrpcServer(config util.Config, store *db.Store) {
	server := gapi.NewServer(store)

	grpcServer := grpc.NewServer()
	pb.RegisterCoursesServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener", err)
	}

	log.Printf("Starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err!=nil{
	}
}
func runGinServer(config util.Config, store *db.Store) {
	server := api.NewServer(store)
	err := server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
