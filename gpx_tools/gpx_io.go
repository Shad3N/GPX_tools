package gpx_tools

import (
	"encoding/xml"
	"io"
	"os"
)

// This function parses an XML file in the
// GPX version 1.1 format and returns a Gpx struct
// or an error if the file could not be parsed.
//
// Returned Gpx is the root element in the XML file and a pointer to GpxType
func ParseGpxFile(path string) (gpx Gpx, err error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	bytes, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	return ParseGpxBytes(bytes)
}

// Basically alias for Unmarshal
func ParseGpxBytes(bytes []byte) (gpx Gpx, err error) {
	err = xml.Unmarshal(bytes, &gpx)
	return gpx, err
}

// WriteGpxFile writes a Gpx struct to a file
// or return an error if the file could not be written.
//
// Beware, the output file may differ from the input file
// because the XML marshalling will fill in the optional
// fields that were not present in the input file.
func WriteGpxFile(gpx Gpx, path string) (err error) {
	xmlFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	bytes, err := xml.MarshalIndent(gpx, "", "    ")
	if err != nil {
		return err
	}

	_, err = xmlFile.Write(bytes)
	return err
}
