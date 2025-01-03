package main

import "testing"

func Test_CalcGrid(t *testing.T) {
	grid := calcGrid(fixtureState())
	for _, r := range grid {
		t.Log(r)
	}
	t.Fail()
}

/*
 */
func fixtureState() GameState {
	return GameState{
		Game: Game{ID: "1", Ruleset: Ruleset{Name: "", Version: "", Settings: RulesetSettings{}}, Map: "", Source: "", Timeout: 0},
		Turn: 0,
		Board: Board{
			Height: 10,
			Width:  10,
			Food: []Coord{
				{0, 2},
				{2, 2},
			},
			Hazards: []Coord{
				{5, 8},
				{7, 8},
			},
			Snakes: []Battlesnake{
				{
					ID:     "1",
					Name:   "1",
					Health: 5,
					Body: []Coord{
						{X: 5, Y: 5},
						{X: 5, Y: 4},
						{X: 5, Y: 3},
						{X: 5, Y: 2},
						{X: 5, Y: 1},
					},
					Head:   Coord{X: 5, Y: 5},
					Length: 5,
				},
			},
		},
		You: Battlesnake{
			ID:     "0",
			Name:   "0",
			Health: 5,
			Body: []Coord{
				{X: 6, Y: 5},
				{X: 6, Y: 4},
				{X: 6, Y: 3},
				{X: 6, Y: 2},
				{X: 6, Y: 1},
			},
			Head:   Coord{X: 6, Y: 5},
			Length: 5,
		},
	}
}
