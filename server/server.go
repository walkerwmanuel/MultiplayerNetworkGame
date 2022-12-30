package server

import (
	"fmt"
	"net/http"
	"webgame/game/deck"

	"github.com/gin-gonic/gin"
)

// Server - the struct to start gin web server
type Server struct {
	Port int
}

// Server.Start - starts the server
func (s *Server) Start() {

	if s.Port == 0 {
		s.Port = 8080
	}

	router := gin.Default()
	router.GET("/deck", getDeck)
	router.GET("/shuffled", getShuffledDeck)
	router.Run(fmt.Sprintf("localhost:%d", s.Port))
}

// getDeck responds with the list of all albums as JSON.
func getDeck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, deck.NewDeck())
}

func getShuffledDeck(c *gin.Context) {
	d := deck.NewDeck()
	d.Shuffle()
	c.IndentedJSON(http.StatusOK, d)
}
