package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	teams := map[string]int{}
	for i, c := range competitions {
		if results[i] == HOME_TEAM_WON {
			teams[c[0]] += 3
		} else {
			teams[c[1]] += 3
		}
	}
	max := 0
	winner := ""
	for t, pts := range teams {
		if pts > max {
			max = pts
			winner = t
		}
	}
	return winner
}

// Test Case 1
// {
//   "competitions": [
//     ["HTML", "C#"],
//     ["C#", "Python"],
//     ["Python", "HTML"]
//   ],
//   "results": [0, 0, 1]
// }
