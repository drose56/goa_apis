package addapi

import (
	add "add/gen/add"
	"context"
	"log"
)

// add service example implementation.
// The example methods log the requests and return zero values.
type addsrvc struct {
	logger *log.Logger
}

// NewAdd returns the add service implementation.
func NewAdd(logger *log.Logger) add.Service {
	return &addsrvc{logger}
}

// Addnums implements addnums.
func (s *addsrvc) Addnums(ctx context.Context, p *add.AddnumsPayload) (res int, err error) {
	return p.A + p.B, nil
}
