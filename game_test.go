package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSortedGames(t *testing.T) {
	testCases := []struct {
		name     string
		in, want []Game
	}{
		{
			name: "Test games with the same total score returned ordered by the most recently added to the system",
			in: []Game{
				NewGame(&Team{Name: "Mexico", Scores: 0}, &Team{Name: "Canada", Scores: 5}),
				NewGame(&Team{Name: "Spain", Scores: 10}, &Team{Name: "Brazil", Scores: 2}),
				NewGame(&Team{Name: "Germany", Scores: 2}, &Team{Name: "France", Scores: 2}),
				NewGame(&Team{Name: "Uruguay", Scores: 6}, &Team{Name: "Italy", Scores: 6}),
				NewGame(&Team{Name: "Argentina", Scores: 3}, &Team{Name: "Australia", Scores: 1}),
			},
			want: []Game{
				NewGame(&Team{Name: "Uruguay", Scores: 6}, &Team{Name: "Italy", Scores: 6}),
				NewGame(&Team{Name: "Spain", Scores: 10}, &Team{Name: "Brazil", Scores: 2}),
				NewGame(&Team{Name: "Mexico", Scores: 0}, &Team{Name: "Canada", Scores: 5}),
				NewGame(&Team{Name: "Argentina", Scores: 3}, &Team{Name: "Australia", Scores: 1}),
				NewGame(&Team{Name: "Germany", Scores: 2}, &Team{Name: "France", Scores: 2}),
			},
		},
		{
			name: "Test all games with the same score returned ordered by the most recently added to the system",
			in: []Game{
				NewGame(&Team{Name: "Mexico", Scores: 0}, &Team{Name: "Canada", Scores: 0}),
				NewGame(&Team{Name: "Spain", Scores: 0}, &Team{Name: "Brazil", Scores: 0}),
				NewGame(&Team{Name: "Germany", Scores: 0}, &Team{Name: "France", Scores: 0}),
				NewGame(&Team{Name: "Uruguay", Scores: 0}, &Team{Name: "Italy", Scores: 0}),
				NewGame(&Team{Name: "Argentina", Scores: 0}, &Team{Name: "Australia", Scores: 0}),
			},
			want: []Game{
				NewGame(&Team{Name: "Argentina", Scores: 0}, &Team{Name: "Australia", Scores: 0}),
				NewGame(&Team{Name: "Uruguay", Scores: 0}, &Team{Name: "Italy", Scores: 0}),
				NewGame(&Team{Name: "Germany", Scores: 0}, &Team{Name: "France", Scores: 0}),
				NewGame(&Team{Name: "Spain", Scores: 0}, &Team{Name: "Brazil", Scores: 0}),
				NewGame(&Team{Name: "Mexico", Scores: 0}, &Team{Name: "Canada", Scores: 0}),
			},
		},
		{
			name: "Test basic sorting",
			in: []Game{
				NewGame(&Team{Name: "Mexico", Scores: 10}, &Team{Name: "Canada", Scores: 5}),
				NewGame(&Team{Name: "Spain", Scores: 8}, &Team{Name: "Brazil", Scores: 6}),
				NewGame(&Team{Name: "Germany", Scores: 15}, &Team{Name: "France", Scores: 3}),
			},
			want: []Game{
				NewGame(&Team{Name: "Germany", Scores: 15}, &Team{Name: "France", Scores: 3}),
				NewGame(&Team{Name: "Mexico", Scores: 10}, &Team{Name: "Canada", Scores: 5}),
				NewGame(&Team{Name: "Spain", Scores: 8}, &Team{Name: "Brazil", Scores: 6}),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sortGames(tc.in)
			if diff := cmp.Diff(tc.in, tc.want); diff != "" {
				t.Errorf("ERROR: \nGot:\n%+v\nWant:\n%+v\nDiff:\n%s", tc.in, tc.want, diff)
			}
		})
	}
}
