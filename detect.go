package main

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"

	"github.com/hajago/filter/document"
	"github.com/hajago/filter/errors"
)

const (
	contentTypesXMLFileName = "[Content_Types].xml"
	contentTypeWord         = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"
	contentTypePresentation = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"
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

func detect(name string) (Document, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == contentTypesXMLFileName {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			byts, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			contentType := ContentType{}
			if err := xml.Unmarshal([]byte(byts), &contentType); err != nil {
				return nil, err
			}

			for _, override := range contentType.Override {
				if override.ContentType == contentTypeWord {
					return document.NewDocx(name), nil
				}else if override.ContentType == contentTypePresentation {
					return document.NewPptx(name), nil
				}
			}
		}
	}

	return nil, errors.ErrNotSupportedFormat
}
