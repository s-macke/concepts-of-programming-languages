// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package stub

import (
	"context"

	"github.com/s-macke/concepts-of-programming-languages/src/distributed/idserv/core"
	"github.com/s-macke/concepts-of-programming-languages/src/distributed/idserv/remote/idserv"
)

// Server is used to implement idserv.IdServer
type Server struct{}

// NewUUID implements idserv.IdService interface
func (s *Server) NewUUID(c context.Context, r *idserv.IdRequest) (*idserv.IdReply, error) {
	service := core.IDServiceImpl{}
	return &idserv.IdReply{Uuid: service.NewUUID(r.GetClientId())}, nil
}
