// Package space will calculate how old someone would be on various planets
// given an age in seconds.
package space

type Planet string

const earthYearInSeconds = float64(31557600)

// Age calculates a persons age on a planet in the Solar System given an input
// of planet name and age in Earth seconds
func Age(ageInSeconds float64, planetName Planet) float64 {
	ageOnPlanet := ageInSeconds / earthYearInSeconds
	orbitalPeriodInEarthYears := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return ageOnPlanet / orbitalPeriodInEarthYears[planetName]
}
