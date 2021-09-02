package main

/*
- install gRPCCURL:
https://github.com/fullstorydev/grpcurl
- grpcurl --plaintext localhost:9092 list => get list of services
- grpcurl --plaintext localhost:9092 list Currency => get list of methods supported bu service Currency
- grpcurl --plaintext localhost:9092 describe Currency.GetRate
- grpcurl --plaintext localhost:9092 describe .RateRequest
- grpcurl --plaintext -d '{"base":"GBP", "destination":"USD"}' localhost:9092 Currency.GetRate
*/

import (
	"net"
	"os"

	protos "Go_Learning/22_grpc/protos/currency"
	currency "Go_Learning/22_grpc/server"

	"github.com/hashicorp/go-hclog"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := currency.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
