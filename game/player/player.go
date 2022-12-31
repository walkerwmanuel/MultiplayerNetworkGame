package player

type Player struct {
	Id       int
	Username string
	Password string
	Credits  int
}

// NewPlayer - returns a pointer to a new player struct
func NewPlayer(i int, n string, m string) *Player {
	// sets p to an empty Player struct
	p := Player{}
	// sets p's Username to n
	p.Id = i
	p.Username = n
	p.Password = m
	p.Credits = 200

	return &p
}
