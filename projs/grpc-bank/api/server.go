package api

import (
	"example.xyz/bank/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(s db.Store) *Server {
	svr := &Server{store: s}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/accounts", svr.createAccount)
	router.GET("/accounts/:id", svr.getAccount)
	router.GET("/accounts", svr.listAccount)

	router.POST("/transfers", svr.createTransfer)

	router.POST("/users", svr.createUser)

	svr.router = router
	return svr
}

// Start runs a the server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
