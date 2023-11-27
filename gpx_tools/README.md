# GPX Tools

A collection of functions for working with GPX files version 1.1.

## Functionalities provided

- Parsing GPX files into Go struct
- Outputting GPX files from Go struct
- Calculating distance between two points by either Haversine or Vincenty formula
- Calculating total distance of a track
- Sorting points by time
- Calculating total time and average speed of a track
- Calculating velocity between two points
- Formatting time to string
- Formatting coordinates to string
- Normalizing angle to be in range of -180 to 180 degrees

## Usage

### Parsing .gpx file
```
ParseGpxFile(path string) (gpx Gpx, err error) 
```

### Writing .gpx file
```
WriteGpxFile(gpx Gpx, path string) (err error)
```

### Distance between points
```
(c *Coordinates) HaversineDistanceFrom(coordinates Coordinates) float64

(c *Coordinates) VincentyDistanceFrom(coordinates Coordinates) float64
```

### Total distance of track
```
TotalLength(track *[]CoordConvertible, algorithm DistanceAlgorithm) float64
```

### Sorting points by time
```
SortByTime(points *[]CoordConvertible) error
```

### Total time and average speed of track
```
TotalTime(track *[]CoordConvertible) (time.Duration, error)

AverageVelocity(points *[]CoordConvertible, algorithm DistanceAlgorithm) (float64, error)
```

### Velocity between two points
```
VelocityBetweenPoints(p1 CoordConvertible, p2 CoordConvertible, algorithm DistanceAlgorithm) (float64, error)
```

### Formatting time to string
```
FormatTime(t time.Time) string
```

### Formatting coordinates to string
```
(c *Coordinates) GetLatitudeAsString() string

(c *Coordinates) GetLongitudeAsString() string

(c *Coordinates) ToString() string
```

### Normalizing coordinates
```
Normalize(degrees float64) float64
```

### Other

```
NewCoordinates(latitude, longitude float64) Coordinates

(c *Coordinates) GetLatitudeRadians() float64

(c *Coordinates) GetLongitudeRadians() float64

(c *Coordinates) GetLatitudeDegrees() float64

(c *Coordinates) GetLongitudeDegrees() float64

(c *Coordinates) GetLatitudeMinutes() int64

(c *Coordinates) GetLongitudeMinutes() int64

(c *Coordinates) GetLatitudeSeconds() float64

(c *Coordinates) GetLongitudeSeconds() float64

(p *EmailType) GetEmailStr() string

ParseGpxBytes(bytes []byte) (gpx Gpx, err error)

NewCoordinates3D(latitude, longitude, altitude float64) Coordinates3D

(wpt *WptType) ToCoordinates() Coordinates

(pt *PtType) ToCoordinates() Coordinates

(wpt *WptType) ToCoordinates3D() Coordinates3D

(pt *PtType) ToCoordinates3D() Coordinates3D
```

