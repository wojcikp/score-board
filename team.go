package main

type Team struct {
	name   string
	scores int
}

func NewTeam(name string) *Team {
	return &Team{name, 0}
}

func (t *Team) addOnePoint() {
	t.scores++
}

func (t *Team) removeOnePoint() {
	if t.scores > 0 {
		t.scores--
	}
}
