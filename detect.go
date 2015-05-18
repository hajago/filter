package main

import (
	"archive/zip"

	"github.com/hajago/filter/document"
	"github.com/hajago/filter/errors"
	"github.com/hajago/filter/types"
)

func detect(name string) (Document, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == types.ContentTypesXMLFileName {
			contentType := types.ContentType{}
			if err := types.UnmarshalZipFile(f, &contentType); err != nil {
				return nil, err
			}

			for _, override := range contentType.Override {
				if override.ContentType == types.ContentTypeWord {
					return document.NewDocx(name)
				} else if override.ContentType == types.ContentTypePresentation {
					return document.NewPptx(name), nil
				} else if override.ContentType == types.ContentTypeSpreadsheet {
					return document.NewXlsx(name), nil
				}
			}
		}
	}

	return nil, errors.ErrNotSupportedFormat
}
