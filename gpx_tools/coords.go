package gpx_tools

import (
	"fmt"
	"math"
)

// Coordinates is a struct that holds Latitude and longitude
// in degrees
// Fields are public, however using
// setters and creators is strongly
// encouraged to avoid invalid values
type Coordinates struct {
	Latitude  float64
	longitude float64
}

// Calculate distance between two coordinates
// using Haversine's formula.
func (c *Coordinates) HaversineDistanceFrom(coordinates Coordinates) float64 {
	return Haversine(*c, coordinates)
}

// Calculate distance between two coordinates
// using Vincenty's formula.
func (c *Coordinates) VincentyDistanceFrom(coordinates Coordinates) float64 {
	return Vincenty(*c, coordinates)
}

// Return Latitude in radians.
func (c *Coordinates) GetLatitudeRadians() float64 {
	return c.Latitude * math.Pi / 180
}

// Return longitude in radians.
func (c *Coordinates) GetLongitudeRadians() float64 {
	return c.longitude * math.Pi / 180
}

// Create new Coordinates object
// with normalized Latitude and longitude.
func NewCoordinates(latitude, longitude float64) Coordinates {
	return Coordinates{Normalize(latitude), Normalize(longitude)}
}

// Set Latitude to normalized value and
// return error if Latitude is NaN or Inf.
func (c *Coordinates) SetLatitude(latitude float64) (err error) {
	if math.IsNaN(latitude) {
		return fmt.Errorf("Latitude is NaN")
	}
	if math.IsInf(latitude, 0) {
		return fmt.Errorf("Latitude is Inf")
	}
	c.Latitude = Normalize(latitude)
	return nil
}

// Set longitude to normalized value and
// return error if longitude is NaN or Inf.
func (c *Coordinates) SetLongitude(longitude float64) (err error) {
	if math.IsNaN(longitude) {
		return fmt.Errorf("Longitude is NaN")
	}
	if math.IsInf(longitude, 0) {
		return fmt.Errorf("Longitude is Inf")
	}
	c.longitude = Normalize(longitude)
	return nil
}

// Return degrees part of Latitude.
func (c *Coordinates) GetLatitudeDegrees() int64 {
	return int64(math.Floor(c.Latitude))
}

// Return degrees part of longitude.
func (c *Coordinates) GetLongitudeDegrees() int64 {
	return int64(math.Floor(c.longitude))
}

// Return minutes part of Latitude with degrees and seconds removed.
func (c *Coordinates) GetLatitudeMinutes() int64 {
	return int64(math.Floor(c.Latitude - float64(c.GetLatitudeDegrees())*60))
}

// Return minutes part of longitude with degrees and seconds removed.
func (c *Coordinates) GetLongitudeMinutes() int64 {
	return int64(math.Floor(c.longitude - float64(c.GetLongitudeDegrees())*60))
}

// Return seconds part of Latitude with degrees and minutes removed and with decimal places.
func (c *Coordinates) GetLatitudeSeconds() float64 {
	return ((c.Latitude-float64(c.GetLatitudeDegrees()))*60 - float64(c.GetLatitudeMinutes())) * 60
}

// Return seconds part of longitude with degrees and minutes removed and with decimal places.
func (c *Coordinates) GetLongitudeSeconds() float64 {
	return ((c.longitude-float64(c.GetLongitudeDegrees()))*60 - float64(c.GetLongitudeMinutes())) * 60
}

// Return Latitude split into degrees, minutes, seconds and formatted as string with symbols.
func (c *Coordinates) GetLatitudeAsString() string {
	return fmt.Sprintf("%d°%d'%f\"", c.GetLatitudeDegrees(), c.GetLatitudeMinutes(), c.GetLatitudeSeconds())
}

// Return longitude split into degrees, minutes, seconds and formatted as string with symbols.
func (c *Coordinates) GetLongitudeAsString() string {
	return fmt.Sprintf("%d°%d'%f\"", c.GetLongitudeDegrees(), c.GetLongitudeMinutes(), c.GetLongitudeSeconds())
}

// Return Latitude and longitude formatted as string.
func (c *Coordinates) ToString() string {
	return fmt.Sprintf("%s, %s", c.GetLatitudeAsString(), c.GetLongitudeAsString())
}

// Normalize Latitude or longitude to range [-180, 180].
func Normalize(degrees float64) float64 {
	if degrees == 0 {
		return 0
	}
	// wtf is minus zero
	degrees = math.Mod(degrees, 360)
	if degrees == -0 {
		return 0
	}

	if degrees > 180 {
		return degrees - 360
	}
	if degrees < -180 {
		return degrees + 360
	}
	return degrees
}
