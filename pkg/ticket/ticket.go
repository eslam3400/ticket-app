package ticket

import (
	"cli-app/pkg/concert"
	"math/rand"
)

type Ticket struct {
	ID      int
	Concert concert.Concert
	Owner   string
}

func (t Ticket) create(concert concert.Concert) Ticket {
	return Ticket{
		ID:      rand.Int(),
		Concert: concert,
	}
}

func (t Ticket) sell(owner string) {
	t.Owner = owner
}

func GenerateTickets() {
	concert := concert.Concert{}
	ticket := Ticket{}
	ticket.create(concert)
	ticket.sell("John")
}
