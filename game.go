package main

import (
	"fmt"
	"sort"
)

var games []Game

type Game struct {
	homeTeam *Team
	awayTeam *Team
}

func NewGame(homeTeam, awayTeam *Team) Game {
	return Game{homeTeam, awayTeam}
}

func (g *Game) getInfo() string {
	return fmt.Sprintf("%s: %d - %s: %d", g.homeTeam.name, g.homeTeam.scores, g.awayTeam.name, g.awayTeam.scores)
}

func getSortedGames(games []Game) []Game {
	sort.SliceStable(games, func(i, j int) bool {
		scoreI := games[i].homeTeam.scores + games[i].awayTeam.scores
		scoreJ := games[j].homeTeam.scores + games[j].awayTeam.scores

		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return true
	})
	return games
}
