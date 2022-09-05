package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"log"
	"os"

	gRPC "hex/internal/adapters/framework/left/grpc"
)

func main() {

	// ports
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("Fail DB init")
	}
	defer dbaseAdapter.CloseDbConnection()



	core := arithmetic.NewAdapter()
	appAdapter := api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter := gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}