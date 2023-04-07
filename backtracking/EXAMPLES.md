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

