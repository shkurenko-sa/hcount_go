package server

import "github.com/shkurenko-sa/hcount_go/internal/users"

func (s *Server) setupRouter() {
	test := s.router.Group("/user")
	{
		test.GET("/test", users.Test)
	}
}
