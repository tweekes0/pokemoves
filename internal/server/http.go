package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tweekes0/pokemonmoves-backend/internal/models"
)

type httpServer struct {
	*models.DBConn
	*gin.Engine
}

func NewHttpServer(db *models.DBConn) *httpServer {
	srv := &httpServer{
		db,
		gin.Default(),
	}
	
	srv.SetupRoutes()	
	return srv 
} 

func (s *httpServer) SetupRoutes() {
	s.GET("/", s.indexHandler())
	s.GET("/pokemon", s.getAllPokemon())
	s.GET("/pokemon/:id", s.validateID(), s.getPokemon())
}