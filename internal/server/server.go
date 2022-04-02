package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *Config
	router *gin.Engine
}

func New() *Server {
	return &Server{
		config: NewConfig(),
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	s.setupRouter()
	return s.router.Run(s.config.Host + ":" + s.config.Port)
}
