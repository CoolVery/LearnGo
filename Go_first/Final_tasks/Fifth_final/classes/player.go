package classes

import "slices"

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func calculateRating(assists, goals, misses int) float64 {
	switch misses == 0 {
	case true:
		return float64(goals + assists/2)
	case false:
		return float64((float64(goals + assists/2)) / float64(misses))
	}
	return 0.0
}

func NewPlayer(name string, goals, misses, assists int) Player {
	newRating := calculateRating(assists, goals, misses)
	newPlayer := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
		Rating:  newRating,
	}
	return newPlayer
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			switch {
			case a.Goals > b.Goals:
				return -1
			case a.Goals < b.Goals:
				return 1
			default:
				return 0
			}
		} else {
			switch {
			case a.Name > b.Name:
				return 1
			case a.Name < b.Name:
				return -1
			default:
				return 0
			}
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			switch {
			case int(a.Rating) >int(b.Rating):
				return -1
			case int(a.Rating) < int(b.Rating):
				return 1
			default:
				return 0
			}
		} else {
			switch {
			case a.Name > b.Name:
				return 1
			case a.Name < b.Name:
				return -1
			default:
				return 0
			}
			
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			if a.Misses == 0 {
				return -1
			} else if b.Misses == 0 {
				return 1
			}
			divRateFirst := a.Goals / a.Misses
			divRateSecond := b.Goals / b.Goals
			switch {
			case divRateFirst > divRateSecond:
				return -1
			case divRateFirst < divRateSecond:
				return 1
			default:
				return 0
			} 
		} else {
			switch {
			case a.Name > b.Name:
				return 1
			case a.Name < b.Name:
				return -1
			default:
				return 0
			}
		}
	})
	return players
}