package types

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"
)

const (
	ContentTypesXMLFileName = "[Content_Types].xml"
	ContentTypeWord         = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"
	ContentTypePresentation = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"
	ContentTypeSpreadsheet  = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
)

type ContentType struct {
	XMLName  xml.Name `xml:"Types"`
	Default  []Default
	Override []Override
}

type Default struct {
	Extension   string `xml:",attr"`
	ContentType string `xml:",attr"`
}

type Override struct {
	ContentType string `xml:",attr"`
	PartName    string `xml:",attr"`
}

func UnmarshalZipFile(f *zip.File, v interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	byts, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal([]byte(byts), v); err != nil {
		return err
	}
	return nil
}
