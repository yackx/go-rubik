package rubik

import (
	"fmt"
)

// A naive solver using BFS.
//
// Because the resolution space is very large, this solver will only work on slightly
// scrambled cubes (3-4 moves away from a solved cube), even though cycles are detected
// and eliminated. Beyond that, it will only consume time and fail with insufficient memory.
//
// See smarter algorithms:
// https://en.wikipedia.org/wiki/Optimal_solutions_for_Rubik%27s_Cube
// "Algorithms for solving the Rubik's cube" - Harpreet Kaur (in this repo)

// Candidate solution, expressed as a succession of cubes and the corresponding moves
// to transition from the initial state to the current state (or solution).
type Vertex struct {
	Path  []Cube
	Moves []Move
}

// Create a new vertex.
func NewVertex(path []Cube, moves []Move) *Vertex {
	return &Vertex{path, moves}
}

// String representation of the solution.
func (vertex Vertex) String() string {
	s := "["
	for i, move := range vertex.Moves {
		s = s + string(move)
		if i != len(vertex.Moves)-1 {
			s = s + ", "
		}
	}
	s = s + "]"
	return s
}

// Check if the given cube in contained in this path (as vertex).
func (vertex Vertex) Contains(cube Cube) bool {
	for _, c := range vertex.Path {
		if c.Equals(cube) {
			return true
		}
	}
	return false
}

// Add a cube state and the corresponding move to this vertex.
func (vertex Vertex) Add(cube Cube, move Move) *Vertex {
	newPath := append(vertex.Path, cube)
	newMoves := append(vertex.Moves, move)
	return NewVertex(newPath, newMoves)
}

// Return the last cube state from this vertex.
func (vertex Vertex) LastCube() Cube {
	return vertex.Path[len(vertex.Path)-1]
}

// (Attempt to) solve the given cube, and return a list of moves.
func Solve(cube Cube) []Move {
	origin := NewVertex([]Cube{cube}, []Move{})
	q := make(chan *Vertex, 100000000)
	q <- origin

	for {
		vertex := <-q
		lastCube := vertex.LastCube()

		for _, move := range Moves {
			newCube := lastCube.Turn(move)
			if vertex.Contains(newCube) {
				continue
			}

			newVertex := vertex.Add(newCube, move)
			if newCube.IsSolved() {
				fmt.Printf("Found a solution: %s\n", newVertex.Moves)
				return newVertex.Moves
				panic("No")
			}
			q <- newVertex
		}
	}

	panic("unreachable statement reached")
}
