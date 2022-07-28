package main

const (
	n = "NORTH"
	s = "SOUTH"
	e = "EAST"
	w = "WEST"
)

// https://www.codewars.com/kata/550f22f4d758534c1100025a/train/go
func DirReduc(dirs []string) []string {
	for skipped := false; !skipped; {
		reduc := make([]string, 0, len(dirs))
		wantLast := true
		for i := 0; i < len(dirs)-1; i++ {
			one, two := dirs[i], dirs[i+1]
			if (one == n && two == s) || (one == s && two == n) ||
				(one == e && two == w) || (one == w && two == e) {
				if i++; i == len(dirs)-1 {
					wantLast = false
				}
				continue
			}
			reduc = append(reduc, dirs[i])
		}

		if wantLast && len(dirs) > 0 {
			reduc = append(reduc, dirs[len(dirs)-1])
		}
		skipped = len(dirs) == len(reduc)
		dirs = reduc
	}
	return dirs
}
