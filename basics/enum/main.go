package enum

type Difficulty int

const (
	Easy     Difficulty = iota // iota == 0
	Medium                     // iota == 1
	Hard                       // iota == 2
	VeryHard                   // iota == 3
)

// Default value problem

// PollAPI acts like we call out to an API that informs us of the current State
// of a request we made. Let's pretend that if we ask about some request it
// doesn't know, it gives back the `int` default value (0).
func PollAPI(found bool) State {
	if !found {
		return 0
	}
	return Progressing
}

//go:generate stringer -type=State
type State int

const (
	Start       State = iota + 1 // iota == 1
	Progressing                  // iota == 2
	Success                      // iota == 3
	Fail                         // iota == 4
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Sport
type Sport int

const (
	Unknown Sport = iota
	Baseball
	Soccer
	Boxing
	Tennis
	Hockey
)

// Using name in front

// os.Mode
//go:generate go run golang.org/x/tools/cmd/stringer -type=Role -trimprefix=Role
type Role uint8

const (
	RoleUnknown Role = iota // iota == 0 -- Good for unknown to be 0!
	RoleGuest
	RoleMember
	RoleModerator
	RoleAdmin
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Status -trimprefix=Status
type Status uint8

const (
	StatusPending Status = iota
	StatusActive
	StatusInactive
	StatusDeactivated
)

// bit by bit

//go:generate go run golang.org/x/tools/cmd/stringer -type=Direction -trimprefix=Direction
type Direction uint8

const (
	DirectionNorth     Direction = 1 << iota // 1 << 1 == 0b00000001 or 1
	DirectionNorthWest                       // 1 << 1 == 0b00000010 or 2
	DirectionWest                            // 1 << 1 == 0b00000100 or 4
	DirectionSouthWest                       // 1 << 1 == 0b00001000 or 8
	DirectionSouth                           // 1 << 1 == 0b00010000 or 16
	DirectionSouthEast                       // 1 << 1 == 0b00100000 or 32
	DirectionEast                            // 1 << 1 == 0b01000000 or 64
	DirectionNorthEast                       // 1 << 1 == 0b10000000 or 128
)

// Multiple flags

//go:generate go run github.com/dmarkham/enumer -type=Day -trimprefix=Day -json -text -yaml -sql
type Day uint8

const (
	DaySunday    Day = 1 << iota // 1 << 0 == 0b00000001 or 1
	DayMonday                    // 1 << 1 == 0b00000010 or 2
	DayTuesday                   // 1 << 2 == 0b00000100 or 4
	DayWednesday                 // 1 << 3 == 0b00001000 or 8
	DayThursday                  // 1 << 5 == 0b00010000 or 16
	DayFriday                    // 1 << 5 == 0b00100000 or 32
	DaySaturday                  // 1 << 6 == 0b01000000 or 64

	// DayAll is a utility enum constant that adds up all of the days.
	// 0b01111111 or 127
	DayAll = DayMonday | DayTuesday | DayWednesday | DayThursday |
		DayFriday | DaySaturday | DaySunday

	// 0b00111110 or 62
	DayWeekdays = DayMonday | DayTuesday | DayWednesday | DayThursday | DayFriday
	DayWeekend  = DaySaturday | DaySunday // 0b01000001 65
)

// Incorrect forms -- May see in the wild

type StrWeekday string

// Look at how we have to repeat the type over and over and over and over....
const (
	StrMonday    StrWeekday = "monday"
	StrTuesday   StrWeekday = "tuesday"
	StrWednesday StrWeekday = "wednesday"
	StrThursday  StrWeekday = "thursday"
	StrFriday    StrWeekday = "friday"
	StrSaturday  StrWeekday = "saturday"
	StrSunday    StrWeekday = "sunday"
)

type StructWeekday struct{ slug string }

func (w StructWeekday) String() string { return w.slug }

// **NONE** of these are constant, they can all change **AND** look at all of
// the repetition of `StructWeekday{...}`
var (
	StructMonday    = StructWeekday{"monday"}
	StructTuesday   = StructWeekday{"tuesday"}
	StructWednesday = StructWeekday{"wednesday"}
	StructThursday  = StructWeekday{"thursday"}
	StructFriday    = StructWeekday{"friday"}
	StructSaturday  = StructWeekday{"saturday"}
	StructSunday    = StructWeekday{"sunday"}
)
