package ticket_window

import (
	"cli-app/pkg/ticket"
	"math/rand"
)

type TicketWindow struct {
	ID      int
	Name    string
	IsOpen  bool
	Tickets []ticket.Ticket
}

func (t *TicketWindow) Create(id int, name string) {
	t.ID = id
	t.Name = name
	t.IsOpen = true
	t.Tickets = []ticket.Ticket{}
}

func (t *TicketWindow) Close() {
	t.IsOpen = false
}

func (t *TicketWindow) AssignTickets(tickets []ticket.Ticket) {
	t.Tickets = append(t.Tickets, tickets...)
}

func GenerateTicketWindows() []TicketWindow {
	ticketWindows := []TicketWindow{}
	count := rand.Intn(5)

	for i := 0; i < count; i++ {
		ticketWindow := TicketWindow{}
		ticketWindow.Create(i, "Window "+string(i))
		ticketWindows = append(ticketWindows, ticketWindow)
	}

	return ticketWindows
}

func FilterOpenWindows(windows []TicketWindow) []TicketWindow {
	openWindows := []TicketWindow{}

	for _, window := range windows {
		if window.IsOpen {
			openWindows = append(openWindows, window)
		}
	}

	return openWindows
}
