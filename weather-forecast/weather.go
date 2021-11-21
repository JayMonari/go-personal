// Package weather implements a simple logging package for the CurrentCondition
// and CurrentLocation and can create logs from the Log function.
package weather

var (
	// CurrentCondition gives the weather condition from the last Log statement
	CurrentCondition string
	// CurrentCondition gives the location of forecast from the last Log statement
	CurrentLocation  string
)

// Log returns the location followed by the current condition.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
