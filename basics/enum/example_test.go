package enum_test

import (
	"basics/enum"
	"encoding/json"
	"fmt"
	"log"
	"math/bits"
	"math/rand"
	"sync"
	"time"
)

func ExampleDifficulty() {
	// NOTE(jay): we use [1:] to get rid of the `\n` newline character at the
	// very start of this string. This is very common to see when using
	// multi-line strings in Go. We would have it look like
	// `Easy: %d
	//Medium: %d ...`
	// otherwise, which can be hard to read.
	fmt.Printf(`
Easy: %d
Medium: %d
Hard: %d
VeryHard: %d`[1:], enum.Easy, enum.Medium, enum.Hard, enum.VeryHard)
	// Output:
	// Easy: 0
	// Medium: 1
	// Hard: 2
	// VeryHard: 3
}

func ExampleStatus() {
	// NOTE(jay): You must run `go generate` or there will be no String method.
	type worker struct {
		s enum.Status
	}
	var wg sync.WaitGroup
	wg.Add(5)
	workers := make([]*worker, 5)
	for i := range workers {
		workers[i] = &worker{}
	}
	for _, w := range workers {
		go func(w *worker) {
			if w.s == enum.StatusPending {
				fmt.Println("worker is", w.s, "giving work")
			}
			w.s = enum.StatusActive
			// Act like a worker is doing work here...
			d := time.Duration(rand.Intn(255))
			time.Sleep(d * time.Millisecond)
			w.s = enum.StatusInactive
			fmt.Println("worker is done with it's work", w.s)
			// Cannot update workers slice here because it will result in data race
			if d > 127 {
				w.s = enum.StatusDeactivated
				fmt.Println("worker took too long", w.s)
			}
			wg.Done()
		}(w)
	}
	wg.Wait()

	fastest := make([]*worker, 0, len(workers))
	for _, w := range workers {
		if w.s != enum.StatusDeactivated {
			fastest = append(fastest, w)
		}
	}
	fmt.Println("The fastest remaining:", len(fastest))
	// XXX(jay): This may fail!!! Remember goroutines don't run one after the
	// other **and** we use `rand` here to get a random number, lots of unknowns.

	// Output:
	// worker is Pending giving work
	// worker is Pending giving work
	// worker is Pending giving work
	// worker is Pending giving work
	// worker is Pending giving work
	// worker is done with it's work Inactive
	// worker is done with it's work Inactive
	// worker is done with it's work Inactive
	// worker took too long Deactivated
	// worker is done with it's work Inactive
	// worker took too long Deactivated
	// worker is done with it's work Inactive
	// worker took too long Deactivated
	// The fastest remaining: 2
}

func ExampleState() {
	// NOTE(jay): You must run `go generate` or there will be no String method.
	if s := enum.PollAPI(true); s != 0 {
		fmt.Println("The current state is:", s)   // same as s.String()
		fmt.Printf("The current state is: %s", s) // can also have %d
	}
	if s := enum.PollAPI(false); s != 0 {
		fmt.Print("Never going to make it in here")
		fmt.Print("Never going to make it in here")
	}
	// Output:
	// The current state is: Progressing
	// The current state is: Progressing
}

func ExampleSport() {
	var s enum.Sport
	fmt.Println(s.String())
	// enum.Baseball.String() works too, but isn't preferred
	fmt.Println(enum.Baseball)
	s = enum.Boxing // == 3. Change this to see our cases print other messages.
	switch s {
	case enum.Baseball:
		fmt.Println("Time to play ball!")
	case enum.Boxing:
		fmt.Println("Round 1! Fight!")
	case enum.Soccer:
		fmt.Println("GOOOOOOOOOOAAAAAAAAAAAL!!!")
	case enum.Hockey:
		fmt.Println("Let's get on the ice!")
	case enum.Tennis:
		fmt.Println("Time for tennis!")
	default:
		fmt.Println("That sport hasn't been added yet! Maybe put in a PR?")
	}
	// Output:
	// Unknown
	// Baseball
	// Round 1! Fight!
}

func ExampleRole() {
	// Let's act like we're routing user traffic and depending on their `Role` we
	// will guide them to a certain page.
	var r enum.Role
	for ; r < 6; r++ {
		switch r {
		case enum.RoleUnknown:
			fmt.Println(r, "should be directed to the login page.")
		case enum.RoleGuest:
			fmt.Println(r, "should be allowed to chat with restrictions.")
		case enum.RoleMember:
			fmt.Println(r, "should be directed to their home page.")
		case enum.RoleModerator:
			fmt.Println(r, "should be routed to the moderator overview page.")
		case enum.RoleAdmin:
			fmt.Println(r, "should be routed to the Administration domain.")
		default:
			fmt.Println("This is not a role!", r)
		}
	}
	// Output:
	// Unknown should be directed to the login page.
	// Guest should be allowed to chat with restrictions.
	// Member should be directed to their home page.
	// Moderator should be routed to the moderator overview page.
	// Admin should be routed to the Administration domain.
	// This is not a role! Role(5)
}

func ExampleDirection() {
	type compass struct{ d enum.Direction }
	c := compass{d: enum.DirectionNorth}
	rotate := func(turnLeft bool) {
		if turnLeft {
			c.d = enum.Direction(bits.RotateLeft8(uint8(c.d), 1))
		} else {
			c.d = enum.Direction(bits.RotateLeft8(uint8(c.d), -1))
		}
		fmt.Println("Compass now facing:", c.d)
	}
	for i := 0; i < 8; i++ {
		rotate(true)
	}
	// Output:
	// Compass now facing: NorthWest
	// Compass now facing: West
	// Compass now facing: SouthWest
	// Compass now facing: South
	// Compass now facing: SouthEast
	// Compass now facing: East
	// Compass now facing: NorthEast
	// Compass now facing: North
}

func ExampleDay() {
	for d := enum.DaySunday; d <= enum.DaySaturday; d <<= 1 {
		switch {
		case d&enum.DayWeekend != 0:
			fmt.Println(d, "-- IT'S TIME TO PARTYYY!!")
		case d&enum.DayWeekdays != 0:
			fmt.Println(d, "-- Time for work...")
		}
	}
	// Output:
	// Sunday -- IT'S TIME TO PARTYYY!!
	// Monday -- Time for work...
	// Tuesday -- Time for work...
	// Wednesday -- Time for work...
	// Thursday -- Time for work...
	// Friday -- Time for work...
	// Saturday -- IT'S TIME TO PARTYYY!!
}

func ExampleDay_second() {
	for i := enum.Day(0); i < enum.DayAll; i++ {
		switch enum.DayAll & i {
		case enum.DayMonday, enum.DayWednesday, enum.DayFriday:
			fmt.Println("Time to workout 💪 it's", i)
		case enum.DayTuesday, enum.DayThursday:
			fmt.Println("Time to rest 🛌💤 on", i)
			// XXX(jay): If you uncomment this you're going to get 127 lines of
			// output... You've been warned -- but it will show why we mask.
			// default:
			// 	fmt.Println("Not a true `enum.Day`", i)
		}
	}
	// Output:
	// Time to workout 💪 it's Monday
	// Time to rest 🛌💤 on Tuesday
	// Time to workout 💪 it's Wednesday
	// Time to rest 🛌💤 on Thursday
	// Time to workout 💪 it's Friday
}

func ExampleDayString() {
	d, err := enum.DayString("Wednesday")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("We got the day: %d -- %q\n", d, d)
	}
	d, err = enum.DayString("wednesday")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("We got the day: %d -- %q\n", d, d)
	}
	d, err = enum.DayString("chewsday")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("We got the day: %d -- %q\n", d, d)
	}
	// Output:
	// We got the day: 8 -- "Wednesday"
	// We got the day: 8 -- "Wednesday"
	// chewsday does not belong to Day values
}

func ExampleDayValues() {
	for _, d := range enum.DayValues() {
		fmt.Println("The day is", uint8(d), d)
	}
	// Output:
	// The day is 0 Unknown
	// The day is 1 Sunday
	// The day is 2 Monday
	// The day is 4 Tuesday
	// The day is 8 Wednesday
	// The day is 16 Thursday
	// The day is 32 Friday
	// The day is 64 Saturday
}

func ExampleDayStrings() {
	for _, d := range enum.DayStrings() {
		fmt.Printf("Only the string %q\n", d)
	}
	// Output:
	// Only the string "Unknown"
	// Only the string "Sunday"
	// Only the string "Monday"
	// Only the string "Tuesday"
	// Only the string "Wednesday"
	// Only the string "Thursday"
	// Only the string "Friday"
	// Only the string "Saturday"
}

func ExampleDay_IsADay() {
	notADay := enum.Day(0)
	if notADay.IsADay() && notADay != enum.DayUnknown {
		fmt.Println("Well I'll be a monkey's uncle 🐒")
	}
	notADay = enum.Day(255)
	if notADay.IsADay() {
		fmt.Println("Well I'll be a monkey's uncle 🐒")
	}
	notADay = enum.DayWeekend
	if notADay.IsADay() {
		fmt.Println("Still won't work, it's not a Day, it's a mask for them.")
	}
	notADay = enum.DayFriday
	if notADay.IsADay() {
		fmt.Println("We validated that", notADay, "is finally, truly, a Day.")
	}
	// Output:
	// We validated that Friday is finally, truly, a Day.
}

func ExampleDay_MarshalJSON() {
	// NOTE(jay): In order to understand why this works we need to understand
	// about the `json` package. In order to `Marshal` types, they must satisfy
	// the `json.Marshaler` interface, which has one method (like any good
	// interface) `MarshalJSON`. The `Marshal` method calls each values
	// `MarshalJSON` method if it has one -- `v.MarshalJSON`. Check out
	// [enum.Day.MarshalJSON] for it's implementation.
	type routine struct {
		Sets      uint8    `json:"sets"`
		Reps      uint8    `json:"reps"`
		Exercises []string `json:"exercises"`
		Day       enum.Day `json:"day"`
	}
	r, err := json.Marshal(routine{
		Sets:      3,
		Reps:      15,
		Exercises: []string{"squats", "calf raises", "leg curls", "hip thrusters"},
		Day:       enum.DayTuesday,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(r))
	// Output:
	// {"sets":3,"reps":15,"exercises":["squats","calf raises","leg curls","hip thrusters"],"day":"Tuesday"}
}

func ExampleDay_UnmarshalJSON() {
	// This mimics some JSON Object we get back from an API or a file.
	resp := `{"sets":5,"reps":21,"exercises":["squats","lunges","leg curls","hip thrusters"],"day":"Thursday"}`
	type routine struct {
		Sets      uint8    `json:"sets"`
		Reps      uint8    `json:"reps"`
		Exercises []string `json:"exercises"`
		Day       enum.Day `json:"day"`
	}
	r := routine{}
	if err := json.Unmarshal([]byte(resp), &r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\nday isn't a string: %d\n", r, r.Day)
	// Output:
	// {Sets:5 Reps:21 Exercises:[squats lunges leg curls hip thrusters] Day:Thursday}
	// day isn't a string: 16
}

func ExampleStrWeekday() {
	// XXX(jay): We cannot manipulate any `StrWeekday` like we did in the
	// previous example **AND** we can make inappropriate values....
	fmt.Println(enum.StrWeekday("This isn't a weekday"))

	// No way to **ENUMERATE** our Enum now.... Have to go through every single
	// one.
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StrMonday)
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StrTuesday)
	// ... Till the end because I'm not typing all of this
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StrSaturday)
	// Output:
	// This isn't a weekday
	// Tedious and not enumerable hence why it's not an enum -- monday
	// Tedious and not enumerable hence why it's not an enum -- tuesday
	// Tedious and not enumerable hence why it's not an enum -- saturday
}

func ExampleStructWeekday() {
	// XXX(jay): **All** `StructWeekday` values can be changed
	enum.StructMonday = enum.StructWednesday
	enum.StructTuesday = enum.StructWednesday
	enum.StructWednesday = enum.StructFriday
	enum.StructThursday = enum.StructSaturday
	enum.StructFriday = enum.StructSaturday
	enum.StructSaturday = enum.StructSunday
	enum.StructSunday = enum.StructWeekday{}
	// No way to **ENUMERATE** our Enum now.... Have to go through every single
	// one.
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StructMonday)
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StructTuesday)
	// ... Till the end because I'm not typing all of this
	fmt.Println("Tedious and not enumerable hence why it's not an enum --",
		enum.StructSaturday)
	// Output:
	// Tedious and not enumerable hence why it's not an enum -- wednesday
	// Tedious and not enumerable hence why it's not an enum -- wednesday
	// Tedious and not enumerable hence why it's not an enum -- sunday
}
