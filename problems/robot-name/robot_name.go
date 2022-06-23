package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// names is a set of all of the distributed names.
var names = map[string]bool{}

const (
	maxNamesSize = alphabetLen * alphabetLen * 1000
	alphabetLen  = 26
)

// Robot is a simple struct with just a name. It will be given a name when its
// Name() method is called and can be reset with the Reset() method.
type Robot struct {
	name string
}

// Name returns the name of the robot or creates a randomly generated one if
// not previously initialized. It will return an error if there are no more
// unique names to be handed out.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(names) == maxNamesSize {
		return "", errors.New("No more available names for production.")
	}
	rand.Seed(time.Now().UnixNano())
	for {
		r.name = fmt.Sprintf("%c%c%03d",
			'A'+rune(rand.Intn(26)),
			'A'+rune(rand.Intn(26)),
			rand.Intn(1000))
		if !names[r.name] {
			names[r.name] = true
			return r.name, nil
		}
	}
}

// Reset removes the previously generated name for the Robot.
func (r *Robot) Reset() {
	r.name = ""
}
