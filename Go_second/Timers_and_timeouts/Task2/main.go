package main

import "time"

func QuizRunner(questions, answers []string, answerCh chan string) int {
	timer := time.NewTimer(time.Second)
	countCorrectAnswer := 0
	for _, q := range questions {
		select {
		case res := <-answerCh:
			for _, a := range answers {
				if res == a {
					if string([]rune(res)[1:]) == string([]rune(q)[1:]) {
						countCorrectAnswer++

						if !timer.Stop() {
							<-timer.C
						}
						timer.Reset(time.Second)
					}
					
				}
			}
			

		case <- timer.C:
			timer.Reset(time.Second)		
		}
	}
	return countCorrectAnswer
}
