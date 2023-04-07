# More exampes on backtracking
## The Knight's tour problem

The Knight's tour problem is a classic puzzle that involves finding a sequence of moves for a knight on a 
chessboard such that the knight visits every square exactly once. This problem can be solved using backtracking, 
which involves trying out all possible moves and undoing them if they don't lead to a solution.

The steps to solve the Knight's tour problem using backtracking are as follows:

1. Initialize the chessboard with all squares set to zero.
2. Place the knight on any square of the chessboard.
3. For each possible move that the knight can make, check if the new square has not been visited yet. If it hasn't, mark the new square with the current move number and recursively try to find a solution starting from the new square.
4. If all possible moves have been tried and none of them lead to a solution, backtrack to the previous square and undo the move.
5. Repeat steps 3-4 until a solution is found or all possible moves have been tried.

Here is the code in Go that solves the Knight's tour problem using backtracking:

```go
package main

import "fmt"

const N = 8

func main() {
    // initialize the chessboard
    board := make([][]int, N)
    for i := range board {
        board[i] = make([]int, N)
    }

    // start the knight's tour from square (0, 0)
    if knightTour(board, 0, 0, 1) {
        printSolution(board)
    } else {
        fmt.Println("No solution exists")
    }
}

func knightTour(board [][]int, x, y, move int) bool {
    if move == N*N+1 {
        // all squares have been visited
        return true
    }

    // possible moves for the knight
    moves := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

    // try all possible moves
    for _, m := range moves {
        nextX, nextY := x+m[0], y+m[1]
        if isValidMove(board, nextX, nextY) {
            board[nextX][nextY] = move
            if knightTour(board, nextX, nextY, move+1) {
                return true
            }
            board[nextX][nextY] = 0
        }
    }

    return false
}

func isValidMove(board [][]int, x, y int) bool {
    // check if the move is within the bounds of the board and the square has not been visited yet
    return x >= 0 && x < N && y >= 0 && y < N && board[x][y] == 0
}

func printSolution(board [][]int) {
    for _, row := range board {
        for _, col := range row {
            fmt.Printf("%2d ", col)
        }
        fmt.Println()
    }
}
```

In this code, the `knightTour` function takes the chessboard, the current position of the knight, 
and the current move number as parameters. It tries all possible moves and returns true if a 
solution is found, and false otherwise.

The isValidMove function checks if the move is within the bounds of the board and the square has not been visited yet.

The printSolution function simply prints the chessboard with the sequence of moves that the knight made.

When the program is run, it will output the sequence of moves that the knight made on the chessboard. 
If no solution exists, it will print "No solution"

## 15 Puzzle Problem

The 15 Puzzle is a popular sliding puzzle game consisting of 15 numbered square tiles in a 4x4 grid, with one empty space where tiles can be moved. The goal of the game is to arrange the tiles in ascending order from left to right, top to bottom, with the empty space in the bottom-right corner.

Backtracking is a recursive algorithm used to solve combinatorial problems, where we explore all possible solutions and backtrack if we reach an invalid solution. In the context of the 15 Puzzle, we can use backtracking to explore all possible moves and their consequences until we find a solution.

Here are the steps to solve the 15 Puzzle using backtracking:

1. Define the initial state of the puzzle, which is a 4x4 grid with the tiles arranged in any order, and one empty space.
2. Define the goal state of the puzzle, which is a 4x4 grid with the tiles arranged in ascending order, and one empty space in the bottom-right corner.
3. Define the possible moves that can be made in the puzzle, which are moving a tile up, down, left or right into the empty space.
4. Use recursion to explore all possible moves from the initial state, and their consequences.
5. If a move leads to the goal state, return that state as the solution.
6. If a move leads to an invalid state (e.g. two tiles are in the same position), backtrack and try a different move.
7. If all possible moves have been explored and none lead to the goal state, backtrack to the previous state and try a different move.
8. Repeat steps 4-7 until a solution is found.

To optimize the backtracking algorithm, we can use heuristics to guide the search towards the goal state. For example, we can use the Manhattan distance heuristic to calculate the total distance between each tile and its correct position, and prioritize moves that reduce the total distance.

Overall, backtracking is a powerful algorithm for solving combinatorial problems like the 15 Puzzle, but it can be computationally expensive for larger problem sizes.

```go
package main

import (
    "fmt"
)

type Board [4][4]int

type Node struct {
    board    Board
    moves    []string
    emptyRow int
    emptyCol int
}

func (b Board) String() string {
    s := ""
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            s += fmt.Sprintf("%2d ", b[i][j])
        }
        s += "\n"
    }
    return s
}

func (n *Node) Move(dir string) (*Node, bool) {
    switch dir {
    case "up":
        if n.emptyRow == 0 {
            return nil, false
        }
        newBoard := n.board
        newBoard[n.emptyRow][n.emptyCol] = newBoard[n.emptyRow-1][n.emptyCol]
        newBoard[n.emptyRow-1][n.emptyCol] = 0
        return &Node{newBoard, append(n.moves, dir), n.emptyRow - 1, n.emptyCol}, true
    case "down":
        if n.emptyRow == 3 {
            return nil, false
        }
        newBoard := n.board
        newBoard[n.emptyRow][n.emptyCol] = newBoard[n.emptyRow+1][n.emptyCol]
        newBoard[n.emptyRow+1][n.emptyCol] = 0
        return &Node{newBoard, append(n.moves, dir), n.emptyRow + 1, n.emptyCol}, true
    case "left":
        if n.emptyCol == 0 {
            return nil, false
        }
        newBoard := n.board
        newBoard[n.emptyRow][n.emptyCol] = newBoard[n.emptyRow][n.emptyCol-1]
        newBoard[n.emptyRow][n.emptyCol-1] = 0
        return &Node{newBoard, append(n.moves, dir), n.emptyRow, n.emptyCol - 1}, true
    case "right":
        if n.emptyCol == 3 {
            return nil, false
        }
        newBoard := n.board
        newBoard[n.emptyRow][n.emptyCol] = newBoard[n.emptyRow][n.emptyCol+1]
        newBoard[n.emptyRow][n.emptyCol+1] = 0
        return &Node{newBoard, append(n.moves, dir), n.emptyRow, n.emptyCol + 1}, true
    }
    return nil, false
}

func (n *Node) IsGoal() bool {
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if i == 3 && j == 3 {
                return true
            }
            if n.board[i][j] != i*4+j+1 {
                return false
            }
        }
    }
    return true
}

func Solve(start Board) ([]string, bool) {
    queue := []*Node{{start, []string{}, 3, 3}}
    visited := make(map[Board]bool)
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        if curr.IsGoal() {
            return curr.moves, true
        }
        visited[curr.board] = true
        for _, dir := range []string{"up", "down", "left", "right"} {
            next, ok := curr.Move(dir)
            if ok && !visited[next.board] {
                queue = append(queue, next)
            }
        }
    }
    return nil, false
}
    
```
This function uses a breadth-first search algorithm to explore the state space of the puzzle. It starts by adding the starting board to a queue of nodes to visit. Then, it repeatedly dequeues a node from the front of the queue, checks if it is the goal state, and if not, generates all possible moves from that state and adds them to the back of the queue (if they have not already been visited). The algorithm continues until the goal state is found or the queue is empty (in which case the puzzle is unsolvable).

```
func main() {
    start := Board{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 0, 12},
        {13, 14, 11, 15},
    }
    fmt.Println("Starting board:")
    fmt.Println(start)
    moves, ok := Solve(start)
    if ok {
        fmt.Println("Solution found in", len(moves), "moves:")
        for _, move := range moves {
            fmt.Println(move)
        }
    } else {
        fmt.Println("Puzzle is unsolvable")
    }
}
```
The `main` function creates a starting board, calls the Solve function to find a solution (if possible), and prints out the solution moves (if found) or a message indicating that the puzzle is unsolvable.



