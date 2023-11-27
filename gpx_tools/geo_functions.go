package gpx_tools

import (
	geo "github.com/olivermichel/vincenty"
	"github.com/umahmood/haversine"
	"time"
)

type DistanceAlgorithm func(Coordinates, Coordinates) float64

type CoordConvertible interface {
	ToCoordinates() Coordinates
	getTimestamp() (time.Time, error)
}

// Perform Vincenty's calculation on two coordinates
// and return the distance between them in meters.
// See: http://en.wikipedia.org/wiki/Vincenty%27s_formulae
//
// This algorithm is very accurate, but also very slow.
func Vincenty(c1 Coordinates, c2 Coordinates) float64 {
	return geo.Vincenty(geo.Point{c1.Latitude, c1.longitude},
		geo.Point{c2.Latitude, c2.longitude})
}

// Perform Haversine's calculation on two coordinates
// and return the distance between them in meters.
// See: http://en.wikipedia.org/wiki/Haversine_formula
//
// This algorithm is less accurate than Vincenty's, but much faster,
// because it assumes the earth is a perfect sphere.
func Haversine(c1 Coordinates, c2 Coordinates) float64 {
	_, km := haversine.Distance(haversine.Coord{c1.Latitude, c1.longitude},
		haversine.Coord{c2.Latitude, c2.longitude})
	return km * 1000.0
}
