package space

type Planet string

var oneEarthYear = float64(31557600)

var planets = map[Planet]float64{
	"Earth":   oneEarthYear,
	"Mercury": oneEarthYear * 0.2408467,
	"Venus":   oneEarthYear * 0.61519726,
	"Mars":    oneEarthYear * 1.8808158,
	"Jupiter": oneEarthYear * 11.862615,
	"Saturn":  oneEarthYear * 29.447498,
	"Uranus":  oneEarthYear * 84.016846,
	"Neptune": oneEarthYear * 164.79132,
}

func Age(sec float64, planet Planet) float64 {
	if planet == "" {
		return -1
	}
	return sec / planets[planet]
}
