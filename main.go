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
	initBoard()
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

func initBoard() {
	board := NewScoreBoard()
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
			board.startNewGame()
		case 2:
			board.getSummaryOfGames()
		case 3:
			board.updateGameScore()
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
	homeTeamName, awayTeamName := getTeamNames()
	homeTeam := NewTeam(homeTeamName)
	awayTeam := NewTeam(awayTeamName)
	game := NewGame(*homeTeam, *awayTeam)
	for {
		clearConsole()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("%s: %d vs. %s: %d\n", homeTeam.name, homeTeam.scores, awayTeam.name, awayTeam.scores)
		fmt.Println("Options: ")
		fmt.Printf("1. %s scored!\n", homeTeam.name)
		fmt.Printf("2. %s scored!\n", awayTeam.name)
		fmt.Printf("3. Remove one point from the %s team\n", homeTeam.name)
		fmt.Printf("4. Remove one point from the %s team\n", awayTeam.name)
		fmt.Println("5. Finish game")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

		switch choice {
		case 1:
			homeTeam.addOnePoint()
			continue
		case 2:
			awayTeam.addOnePoint()
			continue
		case 3:
			homeTeam.removeOnePoint()
			continue
		case 4:
			awayTeam.removeOnePoint()
			continue
		case 5:
			fmt.Println("Finishing game...")
			games = append(games, *game)
			time.Sleep(3 * time.Second)
			return
		default:
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}
	}
}

func getTeamNames() (string, string) {
	var homeTeamName, awayTeamName string
	for {
		clearConsole()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Start new game")
		fmt.Println("-------------------------")

		for {
			if homeTeamName == "" {
		fmt.Println("Enter a name of the Home team and press Enter:")
		scanner.Scan()
				homeTeamName = scanner.Text()
			} else {
				break
			}
		}

		for {
			if awayTeamName == "" {
		fmt.Println("Enter a name of the Away team and press Enter:")
		scanner.Scan()
				awayTeamName = scanner.Text()
			} else {
				break
			}
		}

		fmt.Println("The game is about to start")
		fmt.Println("Home team: ", homeTeamName)
		fmt.Println("Away team: ", awayTeamName)
		fmt.Println("Choose an option:")
		fmt.Println("1. Team names are correct, let's start the game!")
		fmt.Println("2. Team names are incorrect, let's correct them.")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 2 {
			fmt.Println("Invalid input. Please enter a number 1 or 2.")
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

		switch choice {
		case 1:
			return homeTeamName, awayTeamName
		case 2:
			fmt.Println("Correcting team names...")
			homeTeamName = ""
			awayTeamName = ""
			time.Sleep(3 * time.Second)
			continue
		default:
			fmt.Println("Invalid input. Please enter a number 1 or 2.")
			time.Sleep(3 * time.Second)
			continue
		}
	}
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

func (t *Team) addOnePoint() {
	t.scores++
}

func (t *Team) removeOnePoint() {
	if t.scores > 0 {
		t.scores--
	}
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
