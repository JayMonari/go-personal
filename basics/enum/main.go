package enum

//go:generate go run github.com/dmarkham/enumer -type=Difficulty -json -text -yaml -sql
type Difficulty int

const (
	Easy     Difficulty = iota // iota == 0
	Medium                     // iota == 1
	Hard                       // iota == 2
	VeryHard                   // iota == 3
)

// Default value problem

//go:generate stringer -type=State
type State int

const (
	Start       State = iota + 1 // iota == 1
	Progressing                  // iota == 2
	Success                      // iota == 3
	Fail                         // iota == 4
)

type TODO int

const (
	Unknown TODO = iota
)

// Using name in front

//go:generate stringer -type=Role -trimprefix=Role
type Role uint8

const (
	RoleUnknown Role = iota // iota == 0 -- Good for unknown to be 0!
	RoleGuest
	RoleMember
	RoleModerator
	RoleAdmin
)

//go:generate stringer -type=Status
type Status uint8

const (
	StatusPending Status = iota + 1
	StatusActive
	StatusInactive
	StatusDeleted
)

// Multiple flags

//go:generate stringer -type=Direction
type Direction uint8

const (
	DirectionNorth Direction = 1 << iota // 1 << 1 == 0b00000001 or 1
	DirectionEast                        // 1 << 2 == 0b00000010 or 2
	DirectionWest                        // 1 << 3 == 0b00000100 or 4
	DirectionSouth                       // 1 << 4 == 0b00001000 or 8
)

// Multiple ways

//go:generate stringer -type=Weekday
type Weekday uint8

const (
	Monday    Weekday = 1 << iota // 1 << 0 == 0b00000001 or 1
	Tuesday                       // 1 << 1 == 0b00000010 or 2
	Wednesday                     // 1 << 2 == 0b00000100 or 4
	Thursday                      // 1 << 3 == 0b00001000 or 8
	Friday                        // 1 << 5 == 0b00010000 or 16
	Saturday                      // 1 << 5 == 0b00100000 or 32
	Sunday                        // 1 << 6 == 0b01000000 or 64

	// AllDays is a utility enum constant that adds up all of the days.
	AllDays  = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
	WeekDays = Monday | Tuesday | Wednesday | Thursday | Friday
	WeekEnd  = Saturday | Sunday
)

// Odd forms -- May see in the wild

type StrWeekday string

const (
	StrMonday    = "monday"
	StrTuesday   = "tuesday"
	StrWednesday = "wednesday"
	StrThursday  = "thursday"
	StrFriday    = "friday"
	StrSaturday  = "saturday"
	StrSunday    = "sunday"
)

type StructWeekday struct{ slug string }

var (
	SafeMonday    = StructWeekday{"monday"}
	SafeTuesday   = StructWeekday{"tuesday"}
	SafeWednesday = StructWeekday{"wednesday"}
	SafeThursday  = StructWeekday{"thursday"}
	SafeFriday    = StructWeekday{"friday"}
	SafeSaturday  = StructWeekday{"saturday"}
	SafeSunday    = StructWeekday{"sunday"}
)
