package main

import (
  "strings"
  "strconv"
	"container/heap"
	"fmt"
	"os"
)

// Puzzle represents the state of the puzzle.
type Puzzle struct {
	board   []int
	empty   int
	moves   int
	priority int
}

// State is used to implement the heap.Interface for the priority queue.
type State []*Puzzle

func (s State) Len() int           { return len(s) }
func (s State) Less(i, j int) bool { return s[i].priority < s[j].priority }
func (s State) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *State) Push(x interface{}) {
	*s = append(*s, x.(*Puzzle))
}

func (s *State) Pop() interface{} {
	old := *s
	n := len(old)
	puzzle := old[n-1]
	*s = old[0 : n-1]
	return puzzle
}

// Manhattan distance heuristic.
func manhattanDistance(board []int) int {
	distance := 0
	for i, num := range board {
		if num != 0 {
			// Calculate the distance to the goal position.
			rowGoal, colGoal := num/5, num%5
			rowCurrent, colCurrent := i/5, i%5
			distance += abs(rowGoal-rowCurrent) + abs(colGoal-colCurrent)
		}
	}
	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AStar solves the sliding puzzle using the A* algorithm.
func AStar(initial []int, rows, cols int) int {
	initialPuzzle := &Puzzle{board: initial, empty: findEmpty(initial), moves: 0, priority: 0}
	priorityQueue := make(State, 1)
	priorityQueue[0] = initialPuzzle

	heap.Init(&priorityQueue)

	visited := make(map[string]bool)

	for priorityQueue.Len() > 0 {
		currentPuzzle := heap.Pop(&priorityQueue).(*Puzzle)

		if isGoalState(currentPuzzle.board) {
			return currentPuzzle.moves
		}

		visited[getStateKey(currentPuzzle.board)] = true

		for _, neighbor := range getNeighbors(currentPuzzle, rows, cols) {
			if !visited[getStateKey(neighbor.board)] {
				neighbor.priority = neighbor.moves + manhattanDistance(neighbor.board)
				heap.Push(&priorityQueue, neighbor)
			}
		}
	}

	return -1 // No solution found
}

func findEmpty(board []int) int {
	for i, num := range board {
		if num == 0 {
			return i
		}
	}
	return -1
}

func isGoalState(board []int) bool {
	for i, num := range board {
		if num != i+1 && num != 0 {
			return false
		}
	}
	return true
}

func getStateKey(board []int) string {
	key := ""
	for _, num := range board {
		key += fmt.Sprintf("%d", num)
	}
	return key
}

func getNeighbors(puzzle *Puzzle, rows, cols int) []*Puzzle {
	neighbors := make([]*Puzzle, 0)

	emptyRow, emptyCol := puzzle.empty/cols, puzzle.empty%cols

	// Possible moves: up, down, left, right
	for _, move := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newRow, newCol := emptyRow+move[0], emptyCol+move[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			newBoard := make([]int, len(puzzle.board))
			copy(newBoard, puzzle.board)
			newEmpty := newRow*cols + newCol

			newBoard[puzzle.empty], newBoard[newEmpty] = newBoard[newEmpty], newBoard[puzzle.empty]

			neighbors = append(neighbors, &Puzzle{
				board:   newBoard,
				empty:   newEmpty,
				moves:   puzzle.moves + 1,
				priority: 0,
			})
		}
	}

	return neighbors
}

func main() {
  var initialState []int
  file, _ := os.ReadFile(os.Args[1])
  input := strings.Split(string(file), "|")
  for _, num := range input {
    num = strings.ReplaceAll(num, " ", "")
    if num == "X" { initialState = append(initialState, 0) }
    x, err := strconv.Atoi(num)
    if err != nil { continue }
    initialState = append(initialState, x)
  }

	rows, cols := 5, 5
//	initialState = []int{1, 7, 3, 0, 5, 6, 4, 8, 2}
  fmt.Println(initialState)
	moves := AStar(initialState, rows, cols)

	if moves == -1 {
		fmt.Println("No solution found.")
	} else {
		fmt.Printf("Minimum number of moves: %d\n", moves)
	}
}

