package lasagna

// OvenTime returns the amount of time lasagna should be in the oven
const OvenTime = 40

// RemainingOvenTime takes in the actual minutes the lasagna has been in the
// oven and returns how many minutes left until it is done
func RemainingOvenTime(t int) int {
	return OvenTime - t
}

// PreparationTime returns the amount of time it takes with the number of
// layers added to the lasagna
func PreparationTime(layers int) int {
	return 2 * layers
}

// ElapsedTime returns how long you spent working on the lasagna, it takes in
// the amount of layers in the lasagna and the time in minutes it has been in
// the oven
func ElapsedTime(layers, minutes int) int {
	return PreparationTime(layers) + minutes
}
