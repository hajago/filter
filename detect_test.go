package main

import (
	"github.com/hajago/filter/filetype"
	"testing"
)

func TestDetectDocx(t *testing.T) {
	doc, err := detect("Files/test.docx")
	if err != nil {
		t.Error(err)
	}
	if doc == nil {
		t.Error("document is null")
	}
	if doc.FileType() != filetype.DOCX {
		t.Error("file is not DOCX format")
	}
}

func TestDetectPptx(t *testing.T) {
	doc, err := detect("Files/test.pptx")
	if err != nil {
		t.Error(err)
	}
	if doc == nil {
		t.Error("document is null")
	}
	if doc.FileType() != filetype.PPTX {
		t.Error("file is not PPTX format")
	}
}

func TestDetectXlsx(t *testing.T) {
	doc, err := detect("Files/test.xlsx")
	if err != nil {
		t.Error(err)
	}
	if doc == nil {
		t.Error("document is null")
	}
	if doc.FileType() != filetype.XLSX {
		t.Error("file is not XLSX format")
	}
}