package gpx_tools

import (
	"fmt"
	"sort"
	"time"
)

func (p *EmailType) GetEmailStr() string {
	return p.IdAttr + "@" + p.DomainAttr
}

// Gpx uses ISO 8601 format for time, but
// with a small difference, this function
// accomodates for this.
func ParseGpxTimeStr(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z0700", timeStr)
}

func pointsToTimeSlice(points *[]CoordConvertible) (*[]time.Time, error) {
	timestamps := make([]time.Time, len(*points))
	for i := 0; i < len(*points); i++ {
		timestamp, err := (*points)[i].getTimestamp()
		if err != nil {
			return nil, err
		}
		timestamps[i] = timestamp
	}
	return &timestamps, nil
}

// SortByTime sorts a slice of CoordConvertible (WptType, PtType)
// by their timestamp.
// If the timestamp is not defined or could not be parsed, error is returned.
func SortByTime(points *[]CoordConvertible) error {
	timestamps, err := pointsToTimeSlice(points)
	if err != nil {
		return err
	}
	sort.Slice(points, func(i, j int) bool {
		return (*timestamps)[i].Before((*timestamps)[j])
	})
	return nil
}

// Calculate length of segment.
// This function relies on elements
// being in order.
// If you cannot guarantee order
// use SortByTime first.
func TotalLength(track *[]CoordConvertible, algorithm DistanceAlgorithm) float64 {
	distance := 0.0
	convertibles := (*track)
	if len(convertibles) <= 1 {
		return distance
	}

	prev := convertibles[0].ToCoordinates()
	for i := 1; i < len(convertibles); i++ {
		coordinates := convertibles[i].ToCoordinates()
		distance += algorithm(prev, coordinates)
		prev = coordinates
	}
	return distance
}

// TotalTime returns the total time between the first and last
// element in a slice of CoordConvertible (WptType, PtType).
//
// If the slice is empty or has only one element, it returns 0.
// This function will work on unordered slices.
// Returns error if timestamp is not defined or could not be parsed in any element.
func TotalTime(points *[]CoordConvertible) (time.Duration, error) {
	if len(*points) <= 1 {
		return 0, nil
	}
	timestampsPtr, err := pointsToTimeSlice(points)
	if err != nil {
		return 0, err
	}

	timestamps := *timestampsPtr

	minimum := 0
	maximum := 0
	for i := 0; i < len(timestamps); i++ {
		if timestamps[i].Before(timestamps[minimum]) {
			minimum = i
		}
		if timestamps[i].After(timestamps[maximum]) {
			maximum = i
		}
	}
	return timestamps[maximum].Sub(timestamps[minimum]), nil
}

// VelocityBetweenPoints returns the average speed between two points with
// coordinates and timestamps.
//
// Algorithm is a function that takes two coordinates and returns the distance,
// you can use Haversine or Vincenty algorithms from this package.
//
// Result is in m/s.
// Returns error if the two points have the same timestamp as the
// denominator would be 0.
// Returns error if timestamp is not defined or could not be parsed.
func VelocityBetweenPoints(p1 CoordConvertible, p2 CoordConvertible, algorithm DistanceAlgorithm) (float64, error) {
	distance := algorithm(p1.ToCoordinates(), p2.ToCoordinates())
	p1Time, err := p1.getTimestamp()
	if err != nil {
		return 0, err
	}
	p2Time, err := p2.getTimestamp()
	if err != nil {
		return 0, err
	}
	timeDif := p2Time.Sub(p1Time)

	if timeDif.Seconds() == 0.0 {
		return 0, fmt.Errorf("The time difference between points is zero")
	}

	return distance / timeDif.Seconds(), nil
}

// AverageVelocity returns the speed between two points with
// coordinates and timestamps.
//
// Algorithm is a function that takes two coordinates and returns the distance,
// you can use Haversine or Vincenty algorithms from this package.
//
// Result is in m/s.
//
// Returns error if the two points have the same timestamp as the
// denominator would be 0.
// Returns error if any timestamp is not defined or could not be parsed.
func AverageVelocity(points *[]CoordConvertible, algorithm DistanceAlgorithm) (float64, error) {
	totalTime, err := TotalTime(points)
	if err != nil {
		return 0, err
	}

	totalDistance := TotalLength(points, algorithm)

	if totalTime == 0.0 {
		return 0, fmt.Errorf("The time difference between start and end points is zero")
	}

	return totalDistance / totalTime.Seconds(), nil
}
