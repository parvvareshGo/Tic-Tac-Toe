# Tic-Tac-Toe Game in Go

A simple **Tic-Tac-Toe** game implemented in **Go** where you play against the computer. The game runs in the console, with the player using `X` and the computer using `O`.

## Features

* Console-based Tic-Tac-Toe game
* Player vs Computer
* Randomized computer moves
* Checks for win or draw
* Simple and easy-to-understand code

## How to Play

1. Clone the repository or download the source code.
2. Open a terminal in the project directory.
3. Build or run the game:

```bash
go run main.go
```

4. You are `X` and the computer is `O`.
5. Enter the **row** and **column** numbers (0-2) to place your `X`.
6. The computer will automatically place its `O`.
7. The game continues until either the player or computer wins, or the board is full (draw).

### Example

```
-------------
|   |   |   |
-------------
|   |   |   |
-------------
|   |   |   |
-------------
Enter row (0-2): 1
Enter col (0-2): 1
Computer chose row 0, col 2
```

## How it Works

* The board is a 3x3 grid represented by a 2D array.
* `playerMove` function handles user input.
* `computerMove` function picks a random empty cell.
* `checkWin` checks for winning combinations in rows, columns, and diagonals.
* `checkDraw` checks if the board is full without a winner.

## Requirements

* Go 1.18+

## License

This project is licensed under the MIT License.

