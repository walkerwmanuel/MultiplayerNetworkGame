package server

import (
	"fmt"
	"net/http"
	"webgame/game/deck"
	"webgame/game/player"

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

	v1 := router.Group("/api/player")
	{
		v1.GET("", GetPlayer)
		v1.POST("", AddPlayer)
	}

	router.GET("/deck", getDeck)
	router.GET("/shuffled", getShuffledDeck)
	router.Run(fmt.Sprintf("localhost:%d", s.Port))

}

func GetPlayer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, player.NewPlayer(0, "Jimmy", "Password"))
}

func AddPlayer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, player.NewPlayer(0, "Jimmy", "Password"))
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

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
}
