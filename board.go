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

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{}
}

func initBoard() {
	board := NewScoreBoard()
	for {
		clearConsole()
		scanner := bufio.NewScanner(os.Stdin)

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
			scanner.Scan()
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
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

	}
}

func (b ScoreBoard) startNewGame() {
	homeTeamName, awayTeamName := getTeamNamesInput()
	homeTeam := NewTeam(homeTeamName)
	awayTeam := NewTeam(awayTeamName)
	game := NewGame(homeTeam, awayTeam)
	for {
		clearConsole()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("%s: %d vs. %s: %d\n", homeTeam.Name, homeTeam.Scores, awayTeam.Name, awayTeam.Scores)
		fmt.Println("Options: ")
		fmt.Println("-------------------------")
		fmt.Printf("1. %s scored!\n", homeTeam.Name)
		fmt.Printf("2. %s scored!\n", awayTeam.Name)
		fmt.Printf("3. Remove one point from the %s team\n", homeTeam.Name)
		fmt.Printf("4. Remove one point from the %s team\n", awayTeam.Name)
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
			games = append(games, game)
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

func getTeamNamesInput() (string, string) {
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
				homeTeamName = strings.TrimSpace(scanner.Text())
			} else {
				break
			}
		}

		for {
			if awayTeamName == "" {
				fmt.Println("Enter a name of the Away team and press Enter:")
				scanner.Scan()
				awayTeamName = strings.TrimSpace(scanner.Text())
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
	scanner := bufio.NewScanner(os.Stdin)

	sortedGames := make([]Game, len(games))
	copy(sortedGames, games)
	sortGames(sortedGames)

	fmt.Println("Summary of games:")
	if len(sortedGames) == 0 {
		fmt.Println("No games played yet")
		fmt.Println("Press Enter to close the games summary view...")
		scanner.Scan()
		return
	}
	for i, game := range sortedGames {
		fmt.Printf("%d. %s\n", i+1, game.getInfo())
	}

	fmt.Println("Press Enter to close the games summary view...")
	scanner.Scan()
}

func (b ScoreBoard) updateGameScore() {
	for {
		clearConsole()
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Choose a match to edit scores: ")
		if len(games) == 0 {
			fmt.Println("No games played yet")
			fmt.Println("Press Enter to close the update game score view...")
			scanner.Scan()
			return
		}
		for i, game := range games {
			fmt.Println(i+1, game.getInfo())
		}

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(games) {
			fmt.Printf("Invalid input. Please enter a number between 1 and %d.\n", len(games))
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

		editGame := games[choice-1]
		fmt.Println("Match", editGame.getInfo(), "score editing...")

		fmt.Println("What is the new", editGame.HomeTeam.Name, "score?")
		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		newHomeTeamScore, err := strconv.Atoi(input)
		if err != nil || newHomeTeamScore < 0 {
			fmt.Println("Invalid input. Please enter a number greater than 0 or equal.")
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

		fmt.Println("What is the new", editGame.AwayTeam.Name, "score?")
		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)
		newAwayTeamScore, err := strconv.Atoi(input)
		if err != nil || newAwayTeamScore < 0 {
			fmt.Println("Invalid input. Please enter a number greater than 0 or equal.")
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
			continue
		}

		editGame.HomeTeam.Scores = newHomeTeamScore
		editGame.AwayTeam.Scores = newAwayTeamScore

		fmt.Println("Editing operation successful. New scores:", editGame.getInfo())
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		return
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

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error clearing console: %v\n", err)
	}
}
