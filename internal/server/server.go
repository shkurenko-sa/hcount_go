package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config *Config
	router *gin.Engine
	db     *mongo.Database
}

func New() *Server {
	return &Server{
		config: NewConfig(),
		router: gin.Default(),
		db:     nil,
	}
}

func (s *Server) Start() error {
	s.setupRouter()
	s.MongoConnect()
	return s.router.Run(s.config.Host + ":" + s.config.Port)
}
