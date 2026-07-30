package classes

import (
	"math/rand"
)

type Report struct {
	User 
	ReportID (int)
	Date (string)
}

func CreateReport(user User, reportDate string) Report {
	newReport := Report{
		User: user,
		Date: reportDate,
		ReportID: rand.Int(),
	}
	return newReport
}

func  GenerateUserReports(users []User, reportDate string) []Report {
	resultSlice := make([]Report, 0)
	for _, value := range users {
		newRaport := CreateReport(value, reportDate)
		resultSlice = append(resultSlice, newRaport)
	}
	return resultSlice
}