package server

import "github.com/shkurenko-sa/hcount_go/internal/receiver"

func (s *Server) setupRouter() {
	test := s.router.Group("/")
	{
		test.GET("", receiver.Test)
	}

	receive := s.router.Group("/receiver")
	{
		receive.Any("/request")
	}
}
