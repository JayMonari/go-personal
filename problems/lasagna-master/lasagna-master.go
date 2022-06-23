package lasagna

// PreparationTime accepts a slice of ingredients for each layer and the
// average amount of time in minutes it takes to cook each layer. It returns
// the time it takes to cook all of the layers. If the value for avgMins is 0,
// then a default value of 2 minutes per layer is used.
func PreparationTime(layers []string, avgMins int) int {
	if avgMins == 0 {
		avgMins = 2
	}
	return len(layers) * avgMins
}

// Quantities accepts a slice of ingredients for each layer and returns the
// quantity of noodles and sauce needed to make the lasagna.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, ingd := range layers {
		if ingd == "noodles" {
			noodles += 50
		} else if ingd == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient takes in two slices and returns a new slice of the last
// item from the first slice appended to the second slice.
func AddSecretIngredient(frSlice, mySlice []string) []string {
	sl := make([]string, len(mySlice))
	copy(sl, mySlice)
	sl = append(sl, frSlice[len(frSlice)-1])
	return sl
}

// ScaleRecipe returns a new slice of amounts needed for the desired number of
// portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	p := float64(portions) / 2
	sl := make([]float64, len(amounts))
	copy(sl, amounts)
	for i, amt := range sl {
		sl[i] = amt * p
	}
	return sl
}
