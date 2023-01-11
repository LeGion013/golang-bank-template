package api

import (
	db "github.com/LeGion013/banktemplate/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.POST("/accounts/:id", server.getAccount)

	server.router = router
	return server

}

// Start the HTTP server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
