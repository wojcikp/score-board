package main

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
