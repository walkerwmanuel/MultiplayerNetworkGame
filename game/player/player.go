package player

type Player struct {
	Username string
	Credits  int
}

// NewPlayer - returns a pointer to a new player struct
func NewPlayer(n string) *Player {
	// sets p to an empty Player struct
	p := Player{}
	// sets p's Username to n
	p.Username = n
	p.Credits = 200

	return &p
}
