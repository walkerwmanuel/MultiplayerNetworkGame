package server

import (
	"webblackjack/game/deck"
)

type Server struct {
	Port int
}

func server() {
	// getAlbums responds with the list of all albums as JSON.
func getDeck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, deck.NewDeck())

	router := gin.Default()
    router.GET("/deck", getDeck)

    router.Run("localhost:8080")
}
}
