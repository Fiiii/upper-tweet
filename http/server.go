package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router *gin.Engine
}

type InitRouterFunction func(group *gin.RouterGroup)

func BuildServer(router *gin.Engine) *Server {
	return &Server{router: router}
}

func (s *Server) Run() {
	if err := s.router.Run(":8080"); err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}