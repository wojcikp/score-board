package main

import "fmt"

type Team struct {
	Name   string
	Scores int
}

func (t Team) String() string {
	return fmt.Sprintf("%s: %d", t.Name, t.Scores)
}

func NewTeam(name string) *Team {
	return &Team{name, 0}
}

func (t *Team) addOnePoint() {
	t.Scores++
}

func (t *Team) removeOnePoint() {
	if t.Scores > 0 {
		t.Scores--
	}
}
