package robot

import (
	"fmt"
)

var _ fmt.Stringer = Dir(16)

type Dir uint8

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) String() string {
	switch d {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	default:
		panic("direction not valid")
	}
}

var Step1Robot struct {
	X, Y int
	Dir
}

func Right() { Step1Robot.Dir = (Step1Robot.Dir + 1) % 4 }
func Left()  { Step1Robot.Dir = (Step1Robot.Dir - 1) % 4 }
func Advance() {
	switch Step1Robot.Dir {
	case 0:
		Step1Robot.Y++
	case 1:
		Step1Robot.X++
	case 2:
		Step1Robot.Y--
	case 3:
		Step1Robot.X--
	default:
		panic("direction not valid")
	}
}

type Command byte

const (
	CmdRight = 'R'
	CmdLeft  = 'L'
	CmdAccel = 'A'
)

type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

type Action uint8

const (
	complete Action = iota
	advance
	turnLeft
	turnRight
	invalidCmd
)

func (r *Step2Robot) TurnRight() { r.Dir = (r.Dir + 1) % 4 }
func (r *Step2Robot) TurnLeft()  { r.Dir = (r.Dir - 1) % 4 }
func (r *Step2Robot) Advance(rect Rect) bool {
	switch r.Dir {
	case N:
		if r.Pos.Northing < rect.Max.Northing {
			r.Pos.Northing++
			return true
		}
	case E:
		if r.Pos.Easting < rect.Max.Easting {
			r.Pos.Easting++
			return true
		}
	case S:
		if r.Pos.Northing > rect.Min.Northing {
			r.Pos.Northing--
			return true
		}
	case W:
		if r.Pos.Easting > rect.Min.Easting {
			r.Pos.Easting--
			return true
		}
	}
	return false
}

func StartRobot(cmdCh chan Command, axnCh chan Action) {
	for cmd := range cmdCh {
		switch cmd {
		case CmdLeft:
			axnCh <- turnLeft
		case CmdRight:
			axnCh <- turnRight
		case CmdAccel:
			axnCh <- advance
		}
	}
	axnCh <- complete
	close(axnCh)
}

func Room(rect Rect, r Step2Robot, axnCh chan Action, report chan Step2Robot) {
	for axn := range axnCh {
		switch axn {
		case complete:
			report <- r
			return
		case advance:
			r.Advance(rect)
		case turnLeft:
			r.TurnLeft()
		case turnRight:
			r.TurnRight()
		}
	}
}

type Action3 struct {
	name string
	Action
}

type Step3Robot struct {
	Name string
	Step2Robot
}

func (r *Step2Robot) IsInsideRoom(rect Rect) bool {
	return r.Pos.Northing >= rect.Min.Northing && r.Pos.Northing <= rect.Max.Northing &&
		r.Pos.Easting >= rect.Min.Easting && r.Pos.Easting <= rect.Max.Easting
}

func StartRobot3(name, script string, axnCh chan Action3, log chan string) {
	if name == "" {
		return
	}
	for _, b := range []byte(script) {
		switch b {
		case CmdAccel:
			axnCh <- Action3{name: name, Action: advance}
		case CmdLeft:
			axnCh <- Action3{name: name, Action: turnLeft}
		case CmdRight:
			axnCh <- Action3{name: name, Action: turnRight}
		default:
			axnCh <- Action3{name: name, Action: invalidCmd}
			return
		}
	}
	axnCh <- Action3{name: name, Action: complete}
}

func Room3(rect Rect, rr []Step3Robot, axnCh chan Action3, rep chan []Step3Robot, log chan string) {
	defer close(rep)

	posRecords := map[Pos]struct{}{}
	idRecords := map[string]int{}
	for i, r := range rr {
		if r.Name == "" {
			log <- "A robot without a name"
			return
		} else if idRecords[r.Name] > 0 {
			log <- "Duplicate robot names"
			return
		} else if _, ok := posRecords[r.Pos]; ok {
			log <- "Robots placed at the same place"
			return
		} else if !r.IsInsideRoom(rect) {
			log <- "A robot placed outside of the room"
			return
		}
		idRecords[r.Name] = i + 1
		posRecords[r.Pos] = struct{}{}
	}

	for a := range axnCh {
		id := idRecords[a.name] - 1
		if id < 0 {
			log <- "An action from unknown robot"
			return
		}
		switch a.Action {
		case advance:
			r := rr[id].Step2Robot
			lastPos := r.Pos
			if !r.Advance(rect) {
				log <- "A robot is attempting to advance into a wall"
			} else if _, ok := posRecords[r.Pos]; ok {
				log <- "A robot is attempting to advance into another robot"
			} else {
				delete(posRecords, lastPos)
				posRecords[r.Pos] = struct{}{}
				rr[id].Step2Robot = r
			}
		case turnLeft:
			rr[id].Step2Robot.TurnLeft()
		case turnRight:
			rr[id].Step2Robot.TurnRight()
		case complete:
			if delete(idRecords, a.name); len(idRecords) == 0 {
				rep <- rr
				return
			}
		case invalidCmd:
			log <- "An undefined command was found in the script"
			return
		}
	}
}
