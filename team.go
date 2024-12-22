package main

type Team struct {
	Name   string
	Scores int
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
