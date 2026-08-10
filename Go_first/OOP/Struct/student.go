package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (student Student) IsExcellentStudent() bool {
	if student.scoreForOneTask*float64(student.solvedProblems) > student.passingScore {
		return true
	}
	return false
}
