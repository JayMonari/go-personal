package api

import (
	"example.xyz/bank/internal/db"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(s *db.Store) *Server {
	svc := &Server{store: s}
	router := gin.Default()

	router.POST("/accounts", svc.createAccount)
	router.GET("/accounts/:id", svc.getAccount)
	router.GET("/accounts", svc.listAccount)

	svc.router = router
	return svc
}

// Start runs a the server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
