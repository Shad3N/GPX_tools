package tests

import (
	"encoding/xml"
	"gpx_tools"
	"os"
	"testing"
)

func TestNormalize(t *testing.T) {
	if gpx_tools.Normalize(-360.0) != 0 {
		t.Errorf(`Normalize(-360.0) = %f; want 0`, gpx_tools.Normalize(-360.0))
	}

	if gpx_tools.Normalize(360.0) != 0 {
		t.Errorf(`Normalize(360.0) = %f; want 0`, gpx_tools.Normalize(360.0))
	}

	if gpx_tools.Normalize(0.0) != 0 {
		t.Errorf(`Normalize(0.0) = %f; want 0`, gpx_tools.Normalize(0.0))
	}

	if gpx_tools.Normalize(180.0) != 180 {
		t.Errorf(`Normalize(180.0) = %f; want 180`, gpx_tools.Normalize(180.0))
	}

	if gpx_tools.Normalize(-180.0) != -180 {
		t.Errorf(`Normalize(-180.0) = %f; want -180`, gpx_tools.Normalize(-180.0))
	}

	if gpx_tools.Normalize(181.0) != -179 {
		t.Errorf(`Normalize(181.0) = %f; want -179`, gpx_tools.Normalize(181.0))
	}

	if gpx_tools.Normalize(-181.0) != 179 {
		t.Errorf(`Normalize(-181.0) = %f; want 179`, gpx_tools.Normalize(-181.0))
	}
}

func TestParseGpxTime(t *testing.T) {
	time := "2002-02-27T17:18:33Z"
	time2, err := gpx_tools.ParseGpxTimeStr(time)
	if err != nil {
		t.Errorf(`ParseTime("%s") = %v; want nil`, time, err)
	}
	if time2.String() != "2002-02-27 17:18:33 +0000 UTC" {
		t.Errorf(`ParseTime("%s") = %v; want 2002-02-27 17:18:33 +0000 UTC`, time, time2.String())
	}
}

func TestParseGpx(t *testing.T) {
	gpx, err := gpx_tools.ParseGpxFile("/home/shad3n/GolandProjects/GPS_toolkit/gpx_tools/tests/sample.gpx")
	if err != nil {
		t.Errorf(`ParseGpxFile("gpx_test.xml") = %v; want nil`, err)
	}

	if gpx.CreatorAttr != "ThE WiSeGeoGrapHer" {
		t.Errorf(`ParseGpxFile("gpx_test.xml") = %v; want "ThE WiSeGeoGrapHer"`, gpx.CreatorAttr)
	}
	if gpx.Metadata.Author.Name != "Running enthusiast" {
		t.Errorf(`ParseGpxFile("gpx_test.xml") = %v; want "Running enthusiast"`, gpx.Metadata.Author.Name)
	}
}

func TestGetEmailStr(t *testing.T) {
	testMail := gpx_tools.EmailType{
		XMLName:    xml.Name{},
		IdAttr:     "test",
		DomainAttr: "test.com",
	}
	if testMail.GetEmailStr() != "test@test.com" {
		t.Errorf(`GetEmailStr() = %v; want "test@test.com"`, testMail.GetEmailStr())
	}
}

func TestWrite(t *testing.T) {
	gpx, err := gpx_tools.ParseGpxFile("/home/shad3n/GolandProjects/GPS_toolkit/gpx_tools/tests/sample.gpx")
	if err != nil {
		t.Errorf(`ParseGpxFile("gpx_test.xml") = %v; want nil`, err)
	}

	err = gpx_tools.WriteGpxFile(gpx, "/home/shad3n/GolandProjects/GPS_toolkit/gpx_tools/tests/sample2.gpx")
	if err != nil {
		t.Errorf(`WriteGpxFile("gpx_test.xml") = %v; want nil`, err)
	}

	t.Cleanup(func() {
		err := os.Remove("/home/shad3n/GolandProjects/GPS_toolkit/gpx_tools/tests/sample2.gpx")
		if err != nil {
			return
		}
	})

}
