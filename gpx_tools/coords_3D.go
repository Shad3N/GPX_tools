package gpx_tools

type Coordinates3D struct {
	Coordinates Coordinates
	Altitude    float64
}

// Create new Coordinates3D object with normalized Latitude and longitude.
func NewCoordinates3D(latitude, longitude, altitude float64) Coordinates3D {
	return Coordinates3D{NewCoordinates(latitude, longitude), altitude}
}
