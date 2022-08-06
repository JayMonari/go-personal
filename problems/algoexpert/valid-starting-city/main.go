package main

func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	nCities := len(distances)
	for i := 0; i < nCities; i++ {
		milesLeft := 0
		for j := i; j < i+nCities; j++ {
			if milesLeft < 0 {
				continue
			}

			j := j % nCities
			milesLeft += fuel[j]*mpg - distances[j]
		}
		if milesLeft >= 0 {
			return i
		}
	}
	return -1
}

func ValidStartingCityOpt(distances []int, fuel []int, mpg int) int {
	nCities := len(distances)
	milesLeft := 0
	startCityIdx := 0
	milesLeftAtStartCity := 0
	for i := 1; i < nCities; i++ {
		milesLeft += fuel[i-1]*mpg - distances[i-1]
		if milesLeft < milesLeftAtStartCity {
			milesLeftAtStartCity = milesLeft
			startCityIdx = i
		}
	}
	return startCityIdx
}

// Test Case 1
// {
//   "distances": [5, 25, 15, 10, 15],
//   "fuel": [1, 2, 1, 0, 3],
//   "mpg": 10
// }
// Test Case 2
// {
//   "distances": [10, 20, 10, 15, 5, 15, 25],
//   "fuel": [0, 2, 1, 0, 0, 1, 1],
//   "mpg": 20
// }
// Test Case 3
// {
//   "distances": [30, 25, 5, 100, 40],
//   "fuel": [3, 2, 1, 0, 4],
//   "mpg": 20
// }
// Test Case 4
// {
//   "distances": [1, 3, 10, 6, 7, 7, 2, 4],
//   "fuel": [1, 1, 1, 1, 1, 1, 1, 1],
//   "mpg": 5
// }
// Test Case 5
// {
//   "distances": [5, 2, 3],
//   "fuel": [1, 0, 1],
//   "mpg": 5
// }
// Test Case 6
// {
//   "distances": [4, 6],
//   "fuel": [1, 9],
//   "mpg": 1
// }
// Test Case 7
// {
//   "distances": [30, 40, 10, 10, 17, 13, 50, 30, 10, 40],
//   "fuel": [1, 2, 0, 1, 1, 0, 3, 1, 0, 1],
//   "mpg": 25
// }
// Test Case 8
// {
//   "distances": [30, 40, 10, 10, 17, 13, 50, 30, 10, 40],
//   "fuel": [10, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//   "mpg": 25
// }
// Test Case 9
// {
//   "distances": [15, 20, 25, 30, 35, 5],
//   "fuel": [0, 3, 0, 0, 1, 1],
//   "mpg": 26
// }
// Test Case 10
// {
//   "distances": [10, 10, 10, 10],
//   "fuel": [1, 2, 3, 4],
//   "mpg": 4
// }
