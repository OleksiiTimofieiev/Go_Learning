package server

import (
	protos "Go_Learning/22_grpc/protos/currency"
	"context"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base: ", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
