package server

import (
	"github/thallesrangel/api-book-go/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("server is running port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
