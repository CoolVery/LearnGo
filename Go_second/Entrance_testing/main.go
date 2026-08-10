package main

import (
	"github.com/CoolVery/LearnGo.git/classes"
	. "fmt"
)

func main() {
	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04
`
	user := "Паша Попов"
	tasks := classes.GetTasks(chatHistory, &user, nil)
	Println(tasks)
}