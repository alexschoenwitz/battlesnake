package main

import (
	"fmt"
	"log"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "Hi alex. this was automatic after pushing",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	//for range state.Board.Width {
	//}
	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if myNeck.X < myHead.X || myHead.X == 0 { // Neck is left of head, don't move left
		isMoveSafe["left"] = false
	}
	if myNeck.X > myHead.X || myHead.X == boardWidth-1 { // Neck is right of head, don't move right
		isMoveSafe["right"] = false
	}
	if myNeck.Y < myHead.Y || myHead.Y == 0 { // Neck is below head, don't move down
		isMoveSafe["down"] = false
	}
	if myNeck.Y > myHead.Y || myHead.Y == boardHeight-1 { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// mybody := state.You.Body

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	// opponents := state.Board.Snakes

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	//
	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

type Case struct {
	Valid    bool
	Walkable bool
	isCase
}

type isCase interface {
	Valid() bool
	String() string
}

func (c Case) String() string {
	return c.String()
}

type SnakePart struct {
	Snake int
	Index int
}

func (p SnakePart) String() string {
	return fmt.Sprintf("%d_%d", p.Snake, p.Index)
}

func calcGrid(state GameState) [][]Case {
	board := make([][]Case, state.Board.Height)
	for i := range state.Board.Height {
		board[i] = make([]Case, state.Board.Width)
	}

	for s, snake := range state.Board.Snakes {
		for i, coord := range snake.Body {
			board[coord.X][coord.Y] = Case{
				Valid:    true,
				Snake:    s,
				Index:    i,
				Walkable: false,
			}
		}
	}
	for _, hazard := range state.Board.Hazards {
		board[hazard.X][hazard.Y] = Case{
			Valid:    true,
			Snake:    s,
			Index:    i,
			Walkable: false,
		}
	}
	return board
}

func main() {
	RunServer()
}
