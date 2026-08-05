package classes

import (
	"errors"
	"slices"
	"fmt"
)

type Company struct {
	workers []Worker
}

var sliceCheckPosition = []string{
		"директор",
		"зам. директора",
		"начальник цеха",
		"мастер",
		"рабочий",
	}

var mapPosition = map[string]int {
	sliceCheckPosition[0]: 4,
	sliceCheckPosition[1]: 3,
	sliceCheckPosition[2]: 2,
	sliceCheckPosition[3]: 1,
	sliceCheckPosition[4]: 0,
}

func (company *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	
	if !slices.Contains(sliceCheckPosition, position) {
		return errors.New("Такой должности нет")
	}

	company.workers = append(company.workers, Worker{Name: name, Position: position, Salary: salary, Experience: experience})
	return nil
}

func (company *Company) SortWorkers() ([]string, error)  {
		slices.SortFunc(company.workers, func(a, b Worker) int {
		tempSalaryFirst := a.Salary * a.Experience
		tempSalarySecond :=  b.Salary * b.Experience
		tempTopPositionFirst := mapPosition[a.Position]
		tempTopPositionSecond := mapPosition[b.Position]
		if tempSalaryFirst != tempSalarySecond {
			if tempSalaryFirst > tempSalarySecond {
				return -1
			}
			return 0
		}
		if tempTopPositionFirst > tempTopPositionSecond {
			return -1
		}
		return 0
	})	
	resultSlice := make([]string, len(company.workers)) 
	for i := 0; i < len(company.workers); i++ {
		worker := company.workers[i]
		resultSlice[i] = fmt.Sprintf("%s - %d - %s", worker.Name, worker.Salary * worker.Experience, worker.Position)
	}

	return resultSlice, nil
}

