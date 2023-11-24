package GPS_tool

import (
	"fmt"
	"math"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) GetLatitudeRadians() float64 {
	return c.GetLatitude() * math.Pi / 180
}

func (c *Coordinates) GetLongitudeRadians() float64 {
	return c.longitude * math.Pi / 180
}

func NewCoordinates(latitude, longitude float64) Coordinates {
	return Coordinates{normalize(latitude), normalize(longitude)}
}

func (c *Coordinates) SetLatitude(latitude float64) {
	if math.IsNaN(latitude) {
		panic("Latitude is NaN")
	}
	if math.IsInf(latitude, 0) {
		panic("Latitude is Inf")
	}
	c.latitude = normalize(latitude)
}
func (c *Coordinates) SetLongitude(longitude float64) {
	if math.IsNaN(longitude) {
		panic("Latitude is NaN")
	}
	if math.IsInf(longitude, 0) {
		panic("Latitude is Inf")
	}
	c.longitude = normalize(longitude)
}

func (c *Coordinates) GetCoordinates() (float64, float64) {
	return c.latitude, c.longitude
}

func (c *Coordinates) GetLatitude() float64 {
	return c.latitude
}

func (c *Coordinates) GetLongitude() float64 {
	return c.longitude
}

func (c *Coordinates) GetLatitudeDegrees() int64 {
	return int64(math.Floor(c.latitude))
}

func (c *Coordinates) GetLongitudeDegrees() int64 {
	return int64(math.Floor(c.longitude))
}

func (c *Coordinates) GetLatitudeMinutes() int64 {
	return int64(math.Floor(c.latitude - float64(c.GetLatitudeDegrees())*60))
}

func (c *Coordinates) GetLongitudeMinutes() int64 {
	return int64(math.Floor(c.longitude - float64(c.GetLongitudeDegrees())*60))
}

func (c *Coordinates) GetLatitudeSeconds() float64 {
	return ((c.latitude-float64(c.GetLatitudeDegrees()))*60 - float64(c.GetLatitudeMinutes())) * 60
}

func (c *Coordinates) GetLongitudeSeconds() float64 {
	return ((c.longitude-float64(c.GetLongitudeDegrees()))*60 - float64(c.GetLongitudeMinutes())) * 60
}

func (c *Coordinates) GetLatitudeAsString() string {
	return fmt.Sprintf("%d°%d'%f\"", c.GetLatitudeDegrees(), c.GetLatitudeMinutes(), c.GetLatitudeSeconds())
}

func (c *Coordinates) GetLongitudeAsString() string {
	return fmt.Sprintf("%d°%d'%f\"", c.GetLongitudeDegrees(), c.GetLongitudeMinutes(), c.GetLongitudeSeconds())
}

func (c *Coordinates) ToString() string {
	return fmt.Sprintf("%s, %s", c.GetLatitudeAsString(), c.GetLongitudeAsString())
}

func normalize(degrees float64) float64 {
	return degrees - math.Floor(degrees/90)*90
}
