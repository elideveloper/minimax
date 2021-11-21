package tictactoe

import (
	"fmt"

	"github.com/elideveloper/minimax/minimax"
)

const (
	boardMaxSize = 5
)

var (
	BotID  int = 9
	UserID int = 3
)

type Move struct {
	X uint
	Y uint
}

type Board struct {
	size   uint
	winNum uint
	matrix [boardMaxSize][boardMaxSize]int
}

func NewBoard(size, numToWin uint) Board {
	b := Board{
		size:   size,
		winNum: numToWin, // for simple games the same as size
	}

	return b
}

func (b *Board) SetMove(m Move, playerID int) Board {
	// TODO possible to add validation for emptiness

	nb := *b
	if nb.matrix[m.X][m.Y] != 0 {
		panic("cannot move!")
	}
	nb.matrix[m.X][m.Y] = playerID

	return nb
}

func (b *Board) Print() {
	for i := range b.matrix {
		for j := range b.matrix[i] {
			fmt.Print(b.matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func (b *Board) IsWin(playerID int) bool {
	var numToWin, diagRightNumToWin, diagLeftNumToWin, i, j uint
	var diagRUpNumToWin, diagLUpNumToWin, diagRDownNumToWin, diagLDownNumToWin uint

	for i = 0; i < b.size; i++ {
		// check rows
		numToWin = 0
		for j = 0; j < b.size; j++ {
			if b.matrix[j][i] == playerID {
				numToWin++
				if numToWin >= b.winNum {
					return true
				}
			} else {
				numToWin = 0
			}
		}

		// check columns
		numToWin = 0
		for j = 0; j < b.size; j++ {
			if b.matrix[i][j] == playerID {
				numToWin++
				if numToWin >= b.winNum {
					return true
				}
			} else {
				numToWin = 0
			}
		}

		// count for left-to-right diagonal
		if b.matrix[i][i] == playerID {
			diagRightNumToWin++
			if diagRightNumToWin >= b.winNum {
				return true
			}
		} else {
			diagRightNumToWin = 0
		}

		// count for right-to-left diagonal
		if b.matrix[b.size-1-i][i] == playerID {
			diagLeftNumToWin++
			if diagLeftNumToWin >= b.winNum {
				return true
			}
		} else {
			diagLeftNumToWin = 0
		}
	}

	for i = 0; i < b.size-1; i++ {

		if b.matrix[i+1][i] == playerID {
			diagRUpNumToWin++
			if diagRUpNumToWin >= b.winNum {
				return true
			}
		} else {
			diagRUpNumToWin = 0
		}

		if b.matrix[i][i+1] == playerID {
			diagRDownNumToWin++
			if diagRDownNumToWin >= b.winNum {
				return true
			}
		} else {
			diagRDownNumToWin = 0
		}

		if b.matrix[b.size-2-i][i] == playerID {
			diagLUpNumToWin++
			if diagLUpNumToWin >= b.winNum {
				return true
			}
		} else {
			diagLUpNumToWin = 0
		}

		if b.matrix[b.size-1-i][i+1] == playerID {
			diagLDownNumToWin++
			if diagLDownNumToWin >= b.winNum {
				return true
			}
		} else {
			diagLDownNumToWin = 0
		}
	}

	return false
}

func (b *Board) IsEnd() bool {
	var i, j uint
	for i = 0; i < b.size; i++ {
		for j = 0; j < b.size; j++ {
			if b.matrix[j][i] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) GetAllPossibleMoves() []Move {
	possMoves := []Move{}
	var i, j uint
	for i = 0; i < b.size; i++ {
		for j = 0; j < b.size; j++ {
			if b.matrix[i][j] == 0 {
				possMoves = append(possMoves, Move{
					X: i,
					Y: j,
				})
			}
		}
	}

	return possMoves
}

func (b Board) Eval() int {
	if b.IsWin(BotID) {
		return 1000
	}

	if b.IsWin(UserID) {
		return -1000
	}

	return 0
}

func (b Board) GetChildren(isMaximizer bool) []minimax.State {
	if b.Eval() != 0 {
		// endgame
		return []minimax.State{}
	}

	allPossibleMoves := b.GetAllPossibleMoves()
	children := make([]minimax.State, len(allPossibleMoves), len(allPossibleMoves))
	plID := BotID
	if !isMaximizer {
		plID = UserID
	}
	for i := range allPossibleMoves {
		children[i] = b.SetMove(allPossibleMoves[i], plID)
	}
	return children
}
