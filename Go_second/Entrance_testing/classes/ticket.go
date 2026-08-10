package classes

import (
	"slices"
	"strings"
	"time"
)

type Ticket struct {
    Ticket string
    User   string
    Status string
    Date   time.Time
}

func createTicketSlice(text string) []Ticket {
	stringCopy := strings.TrimSpace(text)
	allTicketSplit := strings.Split(stringCopy, "\n")
	ticketSlice := []Ticket {}
	targetStatus := []string{"Готово", "В работе", "Не будет сделано"}
	for _, ticket := range(allTicketSplit) {
		ticketSplit := strings.SplitN(ticket, "-", 2)
		if len(ticketSplit) < 2 {
			continue
		}
		ticketInfo := strings.Split(ticketSplit[1], "_")
		if len(ticketInfo) != 4 {
			continue
		}
		if !slices.Contains(targetStatus, ticketInfo[2]) {
			continue
		}
		timeStamp, err := time.Parse("2006-01-02", ticketInfo[3])
		if err != nil {
			continue
		}
		newTicket := Ticket{
			Ticket: ticketSplit[0] + "-" + ticketInfo[0],
			User: ticketInfo[1],
			Status: ticketInfo[2],
			Date: timeStamp,
		}
		ticketSlice = append(ticketSlice, newTicket)
	}
	return ticketSlice
}

func getTaskForUser(ticketSlice []Ticket, user *string) []Ticket {
	resultSlice := []Ticket{}
	for _, ticket := range(ticketSlice) {
		if ticket.User == *user {
			resultSlice = append(resultSlice, ticket)
		}
	}
	return resultSlice
}

func getTaskForStatus(ticketSlice []Ticket, status *string) []Ticket {
	resultSlice := []Ticket{}
	for _, ticket := range(ticketSlice) {
		if ticket.Status == *status {
			resultSlice = append(resultSlice, ticket)
		}
	}
	return resultSlice
}

func getTaskForUserAndStatus(ticketSlice []Ticket, user, status *string) []Ticket {
	resultSlice := []Ticket{}
	for _, ticket := range(ticketSlice) {
		if ticket.User == *user && ticket.Status == *status {
			resultSlice = append(resultSlice, ticket)
		}
	}
	return resultSlice
}

func GetTasks(text string, user *string, status *string) []Ticket {
	ticketSlice := createTicketSlice(text)
	if user == nil && status == nil {
		return ticketSlice
	} else if user == nil {
		resultSlice := getTaskForStatus(ticketSlice, status)
		return resultSlice
	} else if status == nil {
		resultSlice := getTaskForUser(ticketSlice, user)
		return resultSlice
	} else {
		resultSlice := getTaskForUserAndStatus(ticketSlice, user, status)
		return resultSlice
	}

	
}