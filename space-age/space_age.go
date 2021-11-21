package space

// Planet is a string of one of the eight in the Solar System.
type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// Each value is given seconds
const (
	EarthYear   = 31557600
	MercuryYear = EarthYear * 0.2408467
	VenusYear   = EarthYear * 0.61519726
	MarsYear    = EarthYear * 1.8808158
	JupiterYear = EarthYear * 11.862615
	SaturnYear  = EarthYear * 29.447498
	UranusYear  = EarthYear * 84.016846
	NeptuneYear = EarthYear * 164.79132
)

// Age takes in seconds and a planet and gives the age of a person on that
// planet if the planet is not recognized 0 is returned.
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case Earth:
		return seconds / EarthYear
	case Mercury:
		return seconds / MercuryYear
	case Venus:
		return seconds / VenusYear
	case Mars:
		return seconds / MarsYear
	case Jupiter:
		return seconds / JupiterYear
	case Saturn:
		return seconds / SaturnYear
	case Uranus:
		return seconds / UranusYear
	case Neptune:
		return seconds / NeptuneYear
	default:
		return 0
	}
}
