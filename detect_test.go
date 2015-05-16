package main

import (
	"github.com/hajago/filter/filetype"
	"testing"
)

func TestDetect(t *testing.T) {
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
