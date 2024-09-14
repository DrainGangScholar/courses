package api

import (
	db "main/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func testHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "TEST",
	})
}

func NewServer(db *db.Store) *Server {
	server := &Server{store: db}
	router := gin.Default()
	server.router = router

	router.GET("/test")

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
