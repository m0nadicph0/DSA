# More exampes on backtracking

## 
The N Queen problem is a classic problem in computer science that involves placing N chess queens on an NxN chessboard so that no two queens threaten each other. In other words, no two queens can share the same row, column, or diagonal.

Here are the steps to solve the N Queen problem using backtracking:

Define a function that takes as input the current state of the board, the current column, and the number of queens to place. The current state of the board should be represented as an array of integers, where each integer represents the row of the queen in the corresponding column. If a queen has not yet been placed in a column, the corresponding integer should be set to -1.

If the current column is equal to the number of queens, the board is complete and we have found a solution. Return true to indicate success.

Loop over all possible rows in the current column. For each row, check if placing a queen in that row would result in a valid board state. To do this, check if the current row is already occupied by another queen, if any queen in a previous column threatens the current position, and if any queen in a diagonal threatens the current position.

If the current row is valid, place a queen in that row and recurse on the next column.

If the recursion on the next column returns true, a solution has been found, so return true to indicate success.

If the recursion on the next column returns false, backtrack by removing the queen from the current position and trying the next row.

If none of the possible rows in the current column resulted in a valid board state, return false to indicate failure.

```go
func NQueens(n int) [][]int {
    board := make([]int, n)
    for i := 0; i < n; i++ {
        board[i] = -1
    }
    var solutions [][]int
    solveNQueens(board, 0, &solutions)
    return solutions
}

func solveNQueens(board []int, col int, solutions *[][]int) {
    if col == len(board) {
        // Base case: all queens placed successfully
        *solutions = append(*solutions, append([]int{}, board...))
        return
    }
    for row := 0; row < len(board); row++ {
        if isValid(board, row, col) {
            board[col] = row
            solveNQueens(board, col+1, solutions)
            board[col] = -1
        }
    }
}

func isValid(board []int, row, col int) bool {
    for i := 0; i < col; i++ {
        if board[i] == row {
            // Check if the row is already occupied by a queen
            return false
        }
        if abs(board[i]-row) == abs(i-col) {
            // Check if the diagonal is threatened by a queen
            return false
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

```

And finally the `main` function is as follows:

```go

func main() {
    n := 4
    solutions := NQueens(n)
    fmt.Printf("Solutions for %d queens:\n", n)
    for _, solution := range solutions {
        for _, row := range solution {
            for i := 0; i < n; i++ {
                if i == row {
                    fmt.Print("Q ")
                } else {
                    fmt.Print(". ")
                }
            }
            fmt.Println()
        }
        fmt.Println()
    }
}

```
The NQueens function initializes the board and calls the solveNQueens function to solve the problem. The solveNQueens function uses backtracking to recursively place queens on the board, and stores all valid solutions in the solutions slice.

The isValid function checks if a given position is a valid placement for a queen by checking if the row and diagonal are free of any other queens.

Note that this implementation stores all valid solutions in a slice,


This main function sets n to 4, calls the NQueens function to find all solutions for 4 queens, and then prints each solution in a human-readable format. The output should look like this:
```
Solutions for 4 queens:
. Q . .
. . . Q
Q . . .
. . Q .
    
. . Q .
Q . . .
. . . Q
. Q . .

```

---

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

---
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

```go
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

## Rat in a Maze problem

The **"Rat in a Maze"** problem is a classic example of a problem that can be solved using backtracking. The problem is as follows: imagine a maze represented as a 2D array of 0s and 1s, where 0s represent walls and 1s represent open spaces. There is a rat starting at the top-left corner of the maze, and it wants to reach the bottom-right corner. The rat can only move down or right, and it cannot move through walls. The goal is to find a path for the rat from the starting position to the goal position, if one exists.

Here are the steps to solve this problem using backtracking:

1. Start at the top-left corner of the maze. If this position is the goal position, then we are done. Otherwise, mark the position as visited.
2. Generate all possible moves from the current position (i.e., move down or right one square). If a move is invalid (i.e., it would take the rat through a wall or off the edge of the maze), discard it.
3. For each valid move, recursively call the backtracking function with the new position as the current position.
4. If the function returns true (i.e., a path to the goal position was found), then add the current position to the path and return true.
5. If none of the moves lead to a solution, then backtrack by unmarking the current position as visited and returning false.
6. Repeat steps 2-5 until a solution is found or all possible paths have been explored.

```go
type Point struct {
    row, col int
}

func RatInAMaze(maze [][]int) ([]Point, bool) {
    path := []Point{}
    visited := make(map[Point]bool)
    if solveRatInAMaze(maze, Point{0, 0}, visited, &path) {
        return path, true
    }
    return nil, false
}

func solveRatInAMaze(maze [][]int, curr Point, visited map[Point]bool, path *[]Point) bool {
    if curr.row == len(maze)-1 && curr.col == len(maze[0])-1 {
        // We have reached the goal
        *path = append(*path, curr)
        return true
    }
    visited[curr] = true
    moves := []Point{{1, 0}, {0, 1}} // Down, right
    for _, move := range moves {
        next := Point{curr.row + move.row, curr.col + move.col}
        if next.row >= 0 && next.row < len(maze) && next.col >= 0 && next.col < len(maze[0]) && maze[next.row][next.col] == 1 && !visited[next] {
            if solveRatInAMaze(maze, next, visited, path) {
                *path = append(*path, curr)
                return true
            }
        }
    }
    visited[curr] = false // backtrack
    return false
}
```

This implementation uses a recursive backtracking function called solveRatInAMaze to explore the state space of the maze. The function takes the current position, a set of visited positions, and a pointer to the current path as arguments. It returns true if a path to the goal position was found, and false otherwise.

The RatInAMaze function is a wrapper function that initializes the visited set and calls solveRatInAMaze with the starting position. If a path is found, it returns the path and true. If no path is found, it returns

Finally the `main` function is as follows:

```go
func main() {
    maze := [][]int{
        {1, 0, 0, 0},
        {1, 1, 0, 1},
        {0, 1, 0, 0},
        {1, 1, 1, 1},
    }
    path, found := RatInAMaze(maze)
    if found {
        fmt.Println("Path found:")
        for i := len(path) - 1; i >= 0; i-- {
            fmt.Printf("(%d, %d)\n", path[i].row, path[i].col)
        }
    } else {
        fmt.Println("No path found.")
    }
}
```

In this example, we define a 2D array maze that represents a 4x4 maze. We then call the RatInAMaze function with the maze as an argument, and store the resulting path and a boolean flag indicating whether a path was found.

If a path was found, we print out the path in reverse order (since we built it up from the goal position backwards), with each position represented as a tuple of row and column coordinates. If no path was found, we print a message indicating that.

---
## 
