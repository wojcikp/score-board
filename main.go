package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type ScoreBoard struct{}

type Team struct {
	name   string
	scores int
}

type Game struct {
	homeTeam Team
	awayTeam Team
}

var games []Game

func main() {
	board := NewScoreBoard()
	board.promptHelloScreen()
}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{}
}

func NewTeam(name string) *Team {
	return &Team{name, 0}
}

func NewGame(homeTeam, awayTeam Team) *Game {
	return &Game{homeTeam, awayTeam}
}

func (b ScoreBoard) promptHelloScreen() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		clearConsole()

		fmt.Println("Football World Cup Score Board")
		fmt.Println("-------------------------")
		fmt.Println("Options:")
		fmt.Println("1. Start new game")
		fmt.Println("2. Get a summary of games")
		fmt.Println("3. Update game score")
		fmt.Println("4. Exit")
		fmt.Println("-------------------------")
		fmt.Println("Make a choice and press Enter")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Invalid input. Please enter a number 1, 2, 3 or 4.")
			fmt.Println("Press Enter to continue...")
			continue
		}

		switch choice {
		case 1:
			b.startNewGame()
		case 2:
			b.getSummaryOfGames()
		case 3:
			b.updateGameScore()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid input. Please enter a number 1, 2, 3 or 4.")
			time.Sleep(3 * time.Second)
			continue
		}

	}
}


func (b ScoreBoard) startNewGame() {
	clearConsole()
	fmt.Println("Start new game")
	time.Sleep(2 * time.Second)
	return
}

func (b ScoreBoard) getSummaryOfGames() {
	clearConsole()
	fmt.Println("Summary of games:")
	time.Sleep(2 * time.Second)
	return
}

func (b ScoreBoard) updateGameScore() {
	clearConsole()
	fmt.Println("Update game score")
	time.Sleep(2 * time.Second)
	return
}

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
