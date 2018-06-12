// Package space will calculate how old someone would be on various planets
// given an age in seconds.
package space

const (
	// Orbital period in Earth years
	mercuryEarthYear = 0.2408467
	venusEarthYear   = 0.61519726
	marsEarthYear    = 1.8808158
	jupiterEarthYear = 11.862615
	saturnEarthYear  = 29.447498
	uranusEarthYear  = 84.016846
	neptuneEarthYear = 164.79132
	// Determine number of seconds in each planet's "year"
	earthYearSeconds   = 31557600
	mercuryYearSeconds = mercuryEarthYear * earthYearSeconds
	venusYearSeconds   = venusEarthYear * earthYearSeconds
	marsYearSeconds    = marsEarthYear * earthYearSeconds
	jupiterYearSeconds = jupiterEarthYear * earthYearSeconds
	saturnYearSeconds  = saturnEarthYear * earthYearSeconds
	uranusYearSeconds  = uranusEarthYear * earthYearSeconds
	neptuneYearSeconds = neptuneEarthYear * earthYearSeconds
)

// Age calculates a persons age on a planet in the Solar System given an input
// of planet name and age in Earth seconds
func Age(ageInSeconds float64, planetName string) float64 {
	var ageOnPlanet float64
	switch {
	case planetName == "Mercury":
		ageOnPlanet = ageInSeconds / mercuryYearSeconds
	case planetName == "Venus":
		ageOnPlanet = ageInSeconds / venusYearSeconds
	case planetName == "Earth":
		ageOnPlanet = ageInSeconds / earthYearSeconds
	case planetName == "Mars":
		ageOnPlanet = ageInSeconds / marsYearSeconds
	case planetName == "Jupiter":
		ageOnPlanet = ageInSeconds / jupiterYearSeconds
	case planetName == "Saturn":
		ageOnPlanet = ageInSeconds / saturnYearSeconds
	case planetName == "Uranus":
		ageOnPlanet = ageInSeconds / uranusYearSeconds
	case planetName == "Neptune":
		ageOnPlanet = ageInSeconds / neptuneYearSeconds
	}
	return ageOnPlanet
}
