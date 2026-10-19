package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	timer := time.NewTimer(time.Second)
	countCorrectAnswer := 0
	for i := range questions {
		select {
		case res := <-answerCh:
			for j := range answers {
				if strings.ToLower(strings.TrimSpace(answers[j])) == strings.ToLower(strings.TrimSpace(res)) {
					if i == j {
						countCorrectAnswer++
						break
					}
				}		
			}
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(time.Second)
		case <- timer.C:
			timer.Reset(time.Second)		
		}
					
	}
	return countCorrectAnswer
}
			
	
