package main

import (
	log "github.com/sirupsen/logrus"
	"hexarch/internal/adapters/app/api"
	"hexarch/internal/adapters/core/arithmetic"
	"hexarch/internal/adapters/framework/left/grpc"
	"hexarch/internal/adapters/framework/right/db"
	"hexarch/internal/ports"
	"hexarch/utils"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)

	var err error

	// ports
	var dbAdapter ports.DBPort
	var gRPCAdapter ports.GRPCPort
	var appAdapter ports.APIPort
	var coreAdapter ports.ArithmeticPort

	// adapters
	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DB_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("db ping failure: %v", err)
	}
	defer dbAdapter.CloseDBConnection()

	coreAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, coreAdapter)
	gRPCAdapter = grpc.NewAdapter(appAdapter)
	log.Infof("Starting gRPC server...")
	gRPCAdapter.Run()
}
