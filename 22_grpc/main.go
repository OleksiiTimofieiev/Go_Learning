package main

import (
	"github.com/hashicorp/go-hclog"
	grpc "google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := currency.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":9090")
	if err != {
		log.Error("Unable to listen", "error", err)
		os.exit(1)
	}

	gs.Serve()
}
