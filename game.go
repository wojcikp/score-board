package main

import (
	"fmt"
	"sort"
)

var games []Game

type Game struct {
	HomeTeam *Team
	AwayTeam *Team
}

func (g Game) String() string {
	return fmt.Sprintf("%s vs %s", g.HomeTeam, g.AwayTeam)
}

func NewGame(homeTeam, awayTeam *Team) Game {
	return Game{homeTeam, awayTeam}
}

func (g *Game) getInfo() string {
	return fmt.Sprintf("%s: %d - %s: %d", g.HomeTeam.Name, g.HomeTeam.Scores, g.AwayTeam.Name, g.AwayTeam.Scores)
}

func sortGames(games []Game) {
	sort.SliceStable(games, func(i, j int) bool {
		scoreI := games[i].HomeTeam.Scores + games[i].AwayTeam.Scores
		scoreJ := games[j].HomeTeam.Scores + games[j].AwayTeam.Scores

		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return true
	})
}
