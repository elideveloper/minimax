package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/elideveloper/minimax/minimax"
	"github.com/elideveloper/minimax/tictactoe"
)

func main() {
	for {
		level := 0
		fmt.Println("Начало игры")
		board := tictactoe.NewBoard(5, 4)

		for !board.IsEnd() {
			if level == 0 {
				board = board.SetMove(tictactoe.Move{X: 2, Y: 2}, tictactoe.BotID)
			} else {
				// bot is considered as maximizer player
				board = minimax.FindBestUsingMinimax(board, true).(*tictactoe.Board)
			}

			fmt.Println("\nМой ход!")
			board.Print()
			if board.IsWin(tictactoe.BotID) {
				fmt.Println("Я выиграл!!!")
				break
			}
			if board.IsEnd() {
				fmt.Println("Ничья.")
				break
			}
			time.Sleep(time.Second * 2)

			fmt.Println()
			fmt.Println("Возможные ходы: ")
			possibMoves := board.GetAllPossibleMoves()
			for i, m := range possibMoves {
				buffBoard := board.SetMove(m, tictactoe.UserID)
				fmt.Printf("номер хода '%d'\n", i)
				buffBoard.Print()
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Введите номер выбранного хода: ")

			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Println("Ваш ход: ", text)

			if text == "q\n" {
				break
			}

			moveIndex, err := strconv.Atoi(text[:len(text)-1])
			if err != nil {
				panic(err)
			}

			board = board.SetMove(possibMoves[moveIndex], tictactoe.UserID)
			fmt.Println("Доска после вашего хода")
			board.Print()
			if board.IsWin(tictactoe.UserID) {
				fmt.Println("Вы выиграли.")
				break
			}

			level++
			time.Sleep(time.Second * 1)
		}

		time.Sleep(time.Second * 4)
		fmt.Println()
	}
}
