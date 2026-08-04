package classes

import (
	"errors"
	"slices"
)

type Company struct {
	workers []Worker
}

func (company *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	sliceCheckPosition := []string{
		"директор",
		"зам. директора",
		"начальник цеха",
		"мастер",
		"рабочий",
	}
	if !slices.Contains(sliceCheckPosition, position) {
		return errors.New("Такой должности нет")
	}

	company.workers = append(company.workers, Worker{name: name, position: position, salary: salary, experience: experience})
	return nil
}