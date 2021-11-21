package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const format = "%-31s|%3v |%3v |%3v |%3v |%3v\n"

// Team represents a team from the tournament.
type Team struct {
	Name string
	// Total amount of matches played.
	Matches int
	// Total points awarded to the team based on Wins and Draws
	Points int
	// Awards 3 points
	Wins int
	// Awards 1 point
	Draws int
	// Awards 0 points
	Losses int
}

// Update updates the current records of the season for a team. If the provided
// outcome is not "win", "draw", or "loss" an error is returned.
func (t *Team) Update(outcome string) error {
	switch outcome {
	case "win":
		t.Wins++
		t.Points += 3
	case "draw":
		t.Draws++
		t.Points++
	case "loss":
		t.Losses++
	default:
		return fmt.Errorf("no outcome matches: %s", outcome)
	}
	t.Matches++
	return nil
}

// String ...
func (t Team) String() string {
	return fmt.Sprintf(format, t.Name, t.Matches, t.Wins, t.Draws, t.Losses, t.Points)
}

// Tally takes a given input of teams delimited by semi-colons and writes to
// output a formatted table of the teams statistics in ascending value based
// off of points, matches played, and team names, in that order. If not
// properly delimited or an outcome is not win, loss, or draw an error is
// returned.
func Tally(input io.Reader, output io.Writer) error {
	teams := make(map[string]*Team)
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		if len(sc.Text()) == 0 || sc.Text()[0] == '#' {
			continue
		}

		line := strings.Split(sc.Text(), ";")
		if len(line) != 3 {
			return fmt.Errorf("incorrectly formatted input: %s", line)
		}

		team1, team2 := line[0], line[1]
		outcome1, outcome2 := line[2], flip(line[2])
		if err := update(teams, team1, outcome1); err != nil {
			return err
		} else if err := update(teams, team2, outcome2); err != nil {
			return err
		}
	}

	fmt.Fprintf(output, format, "Team", "MP", "W", "D", "L", "P")
	for _, t := range mapToSlice(teams) {
		fmt.Fprint(output, t.String())
	}

	return nil
}

// mapToSlice ...
func mapToSlice(teams map[string]*Team) []Team {
	slTeams := make([]Team, 0, len(teams)/2)
	for _, t := range teams {
		slTeams = append(slTeams, *t)
	}
	// We sort by greatest points then by team names, lexographically.
	sort.Slice(slTeams, func(i, j int) bool {
		if p1, p2 := slTeams[i].Points, slTeams[j].Points; p1 == p2 {
			return slTeams[i].Name < slTeams[j].Name
		} else {
			return p1 > p2
		}
	})
	return slTeams
}

// update creates a Team for the teams map if the team's name has not already
// been created and updates the values for that team.
func update(teams map[string]*Team, name string, outcome string) error {
	if _, ok := teams[name]; !ok {
		teams[name] = &Team{Name: name}
	}
	err := teams[name].Update(outcome)
	return err
}

// flip turns a win into a loss and a loss into a win. If any other value is
// given a draw is returned.
func flip(o string) string {
	switch o {
	case "win":
		return "loss"
	case "loss":
		return "win"
	default:
		return "draw"
	}
}
