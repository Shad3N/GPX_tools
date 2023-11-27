package gpx_tools

import (
	"encoding/xml"
	"time"
)

// Gpx is the root element in the XML file.
type Gpx *GpxType

// GpxType is You can add extend GPX by adding your own elements from another schema here.
type GpxType struct {
	XMLName     xml.Name        `xml:"gpx"`
	VersionAttr string          `xml:"version,attr"`
	CreatorAttr string          `xml:"creator,attr"`
	Metadata    *MetadataType   `xml:"metadata"`
	Wpt         []*WptType      `xml:"wpt"`
	Rte         []*RteType      `xml:"rte"`
	Trk         []*TrkType      `xml:"trk"`
	Extensions  *ExtensionsType `xml:"extensions"`
}

// MetadataType is You can add extend GPX by adding your own elements from another schema here.
type MetadataType struct {
	XMLName    xml.Name        `xml:"metadata"`
	Name       string          `xml:"name"`
	Desc       string          `xml:"desc"`
	Author     *PersonType     `xml:"author"`
	Copyright  *CopyrightType  `xml:"copyright"`
	Link       []*LinkType     `xml:"link"`
	Time       string          `xml:"time"`
	Keywords   string          `xml:"keywords"`
	Bounds     *BoundsType     `xml:"bounds"`
	Extensions *ExtensionsType `xml:"extensions"`
}

// WptType is You can add extend GPX by adding your own elements from another schema here.
type WptType struct {
	XMLName       xml.Name        `xml:"wpt"`
	LatAttr       float64         `xml:"lat,attr"`
	LonAttr       float64         `xml:"lon,attr"`
	Ele           float64         `xml:"ele"`
	Time          string          `xml:"time"`
	Magvar        float64         `xml:"magvar"`
	Geoidheight   float64         `xml:"geoidheight"`
	Name          string          `xml:"name"`
	Cmt           string          `xml:"cmt"`
	Desc          string          `xml:"desc"`
	Src           string          `xml:"src"`
	Link          []*LinkType     `xml:"link"`
	Sym           string          `xml:"sym"`
	Type          string          `xml:"type"`
	Fix           string          `xml:"fix"`
	Sat           int             `xml:"sat"`
	Hdop          float64         `xml:"hdop"`
	Vdop          float64         `xml:"vdop"`
	Pdop          float64         `xml:"pdop"`
	Ageofdgpsdata float64         `xml:"ageofdgpsdata"`
	Dgpsid        int             `xml:"dgpsid"`
	Extensions    *ExtensionsType `xml:"extensions"`
}

// RteType is A list of route points.
type RteType struct {
	XMLName    xml.Name        `xml:"rte"`
	Name       string          `xml:"name"`
	Cmt        string          `xml:"cmt"`
	Desc       string          `xml:"desc"`
	Src        string          `xml:"src"`
	Link       []*LinkType     `xml:"link"`
	Number     int             `xml:"number"`
	Type       string          `xml:"type"`
	Extensions *ExtensionsType `xml:"extensions"`
	Rtept      []*WptType      `xml:"rtept"`
}

// TrkType is A Track Segment holds a list of Track Points which are logically connected in order.
// To represent a single GPS track where GPS reception was lost, or the GPS receiver was turned off, start a new Track Segment for each continuous span of track data.
type TrkType struct {
	XMLName    xml.Name        `xml:"trk"`
	Name       string          `xml:"name"`
	Cmt        string          `xml:"cmt"`
	Desc       string          `xml:"desc"`
	Src        string          `xml:"src"`
	Link       []*LinkType     `xml:"link"`
	Number     int             `xml:"number"`
	Type       string          `xml:"type"`
	Extensions *ExtensionsType `xml:"extensions"`
	Trkseg     []*TrksegType   `xml:"trkseg"`
}

// ExtensionsType is You can add extend GPX by adding your own elements from another schema here.
type ExtensionsType struct {
	XMLName xml.Name `xml:"extensions"`
}

// TrksegType is You can add extend GPX by adding your own elements from another schema here.
type TrksegType struct {
	XMLName    xml.Name        `xml:"trkseg"`
	Trkpt      []*WptType      `xml:"wpt"`
	Extensions *ExtensionsType `xml:"extensions"`
}

// CopyrightType is Link to external file containing license text.
type CopyrightType struct {
	XMLName    xml.Name `xml:"copyright"`
	AuthorAttr string   `xml:"author,attr"`
	Year       string   `xml:"year"`
	License    string   `xml:"license"`
}

// LinkType is Mime type of content (image/jpeg)
type LinkType struct {
	XMLName  xml.Name `xml:"link"`
	HrefAttr string   `xml:"href,attr"`
	Text     string   `xml:"text"`
	Type     string   `xml:"type"`
}

// EmailType is An email address.  Broken into two parts (id and domain) to help prevent email harvesting.
type EmailType struct {
	XMLName    xml.Name `xml:"email"`
	IdAttr     string   `xml:"id,attr"`
	DomainAttr string   `xml:"domain,attr"`
}

// PersonType is Link to Web site or other external information about person.
type PersonType struct {
	XMLName xml.Name   `xml:"author"`
	Name    string     `xml:"name"`
	Email   *EmailType `xml:"email"`
	Link    *LinkType  `xml:"link"`
}

// PtType is a point with timestamp
type PtType struct {
	XMLName xml.Name `xml:"ptType"`
	LatAttr float64  `xml:"lat,attr"`
	LonAttr float64  `xml:"lon,attr"`
	Ele     float64  `xml:"ele"`
	Time    string   `xml:"time"`
}

// PtsegType is Ordered list of geographic points.
type PtsegType struct {
	XMLName xml.Name  `xml:"ptsegType"`
	Pt      []*PtType `xml:"pt"`
}

// BoundsType is Two lat/lon pairs defining the extent of an element.
type BoundsType struct {
	XMLName    xml.Name `xml:"bounds"`
	MinlatAttr float64  `xml:"minlat,attr"`
	MinlonAttr float64  `xml:"minlon,attr"`
	MaxlatAttr float64  `xml:"maxlat,attr"`
	MaxlonAttr float64  `xml:"maxlon,attr"`
}

// LatitudeType is The Latitude of the point.  Decimal degrees, WGS84 datum.
type LatitudeType float64

// LongitudeType is The longitude of the point.  Decimal degrees, WGS84 datum.
type LongitudeType float64

// DegreesType is Used for bearing, heading, course.  Units are decimal degrees, true (not magnetic).
type DegreesType float64

// FixType is Type of GPS fix.  none means GPS had no fix.  To signify "the fix info is unknown, leave out fixType entirely. pps = military signal used
type FixType string

// DgpsStationType is Represents a differential GPS station.
type DgpsStationType int

// Convert BoundsType to a tuple of Coordinates.
func (bounds *BoundsType) ToCoordinates() (minCoords Coordinates, maxCoords Coordinates) {
	minCoords = NewCoordinates(bounds.MinlatAttr, bounds.MinlonAttr)
	maxCoords = NewCoordinates(bounds.MaxlatAttr, bounds.MaxlonAttr)
	return minCoords, maxCoords
}

// Convert WptType to Coordinates
// used by gpx_toolkit to represent
// universal coordinates and perform
// calculations on them.
func (wpt *WptType) ToCoordinates() Coordinates {
	return NewCoordinates(wpt.LatAttr, wpt.LonAttr)
}

// Convert WptType to Coordinates3D
// used byt gpx_toolkit to represent
// universal coordinates with elevation
// and perform calculations on them.
func (wpt *WptType) ToCoordinates3D() Coordinates3D {
	return NewCoordinates3D(wpt.LatAttr, wpt.LonAttr, wpt.Ele)
}

func (wpt *WptType) getTimestamp() (time.Time, error) {
	return ParseGpxTimeStr(wpt.Time)
}

// Convert PtType to Coordinates
// used by gpx_toolkit to represent
// universal coordinates and perform
// calculations on them.
func (pt *PtType) ToCoordinates() Coordinates {
	return NewCoordinates(pt.LatAttr, pt.LonAttr)
}

// Convert PtType to Coordinates3D
// used by gpx_toolkit to represent
// universal coordinates with elevation
// and perform calculations on them.
func (pt *PtType) ToCoordinates3D() Coordinates3D {
	return NewCoordinates3D(pt.LatAttr, pt.LonAttr, pt.Ele)
}

func (pt *PtType) getTimestamp() (time.Time, error) {
	return ParseGpxTimeStr(pt.Time)
}
