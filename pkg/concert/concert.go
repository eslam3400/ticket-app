package concert

import (
	"cli-app/pkg/util"
	"math/rand"
)

type Concert struct {
	ID       int
	Name     string
	Location string
}

func (c Concert) create(name string, location string) Concert {
	return Concert{
		ID:       rand.Int(),
		Name:     name,
		Location: location,
	}
}

func GenerateConcerts() []Concert {
	concerts := []Concert{}
	count := rand.Intn(10)

	for i := 0; i < count; i++ {
		concert := Concert{}
		concert.create(util.GenerateRandomText(), util.GenerateRandomText())
	}

	return concerts
}
