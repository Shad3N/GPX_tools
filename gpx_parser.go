package GPS_tool

import (
	"encoding/xml"
	"os"
	"time"
)

type GpxData struct {
	Creator    string      `xml:"creator,attr"`
	Version    string      `xml:"version,attr"`
	Metadata   GpxMetadata `xml:"metadata,omitempty"`
	Extensions Extensions  `xml:"extensions,omitempty"`

	Waypoints []GpxWaypoint `xml:"wpt,omitempty"`
	Routes    []GpxRoute    `xml:"rte,omitempty"`
	Tracks    []GpxTrack    `xml:"trk,omitempty"`
}

type Extensions struct {
	Elements []xml.Attr `xml:",any"`
}

type GpxWaypoint struct {
	Latitude      float64 `xml:"lat,attr"`
	Longitude     float64 `xml:"lon,attr"`
	Elevation     float64
	Time          time.Time
	Magvar        float64
	Geoidheight   float64
	Name          string
	Cmt           string
	Desc          string
	Src           string
	Links         []GpxLink
	Sym           string
	Type          string
	Fix           GSPFix
	Sat           uint64
	Hdop          float64
	Vdop          float64
	Pdop          float64
	Ageofdgpsdata float64
	Dgpsid        uint64
	Extensions    map[string]string
}

type GpxRoute struct {
	Name        string
	Cmt         string
	Desc        string
	Src         string
	Links       []GpxLink
	Number      uint64
	Type        string
	Extensions  map[string]string
	RoutePoints []GpxWaypoint
}

type GpxTrack struct {
	Name       string
	Cmt        string
	Desc       string
	Src        string
	Links      []GpxLink
	Number     uint64
	Type       string
	Extensions map[string]string
	Segments   []GpxTrackSegment
}

type GpxTrackSegment struct {
	TrackPoints []GpxWaypoint
	Extensions  map[string]string
}

type GSPFix string

const (
	FixNone GSPFix = "none"
	Fix2d   GSPFix = "2d"
	Fix3d   GSPFix = "3d"
	FixDGPS GSPFix = "dgps"
	FixPPS  GSPFix = "pps"
)

func (wpt *GpxWaypoint) GetCoordinates() Coordinates {
	return NewCoordinates(wpt.Latitude, wpt.Longitude)
}

type GpxMetadata struct {
	Name        string
	Description string
	Author      GpxPerson
	Copyright   GpxCopyright
	Links       []GpxLink
	Timestamp   time.Time
	Keywords    string
	Bounds      GpxBounds
}

type GpxPerson struct {
	Name  string
	Email string
	Link  string
}

type GpxLink struct {
	Href string
	Text string
	Type string
}

type GpxCopyright struct {
	Author  string
	Year    int
	License string
}

type GpxBounds struct {
	MinLat float64
	MinLon float64
	MaxLat float64
	MaxLon float64
}

func ParseGpx(path string) (*GpxData, error) {
	data := GpxData{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func parseGpxMetadata(decoder *xml.Decoder, metadata *GpxMetadata) error {
	return nil
}
