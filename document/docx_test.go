package document

import (
	"os"
	"testing"
)

func TestDocxFilter(t *testing.T) {
	doc, err := NewDocx("../Files/test.docx")
	if err != nil {
		t.Error(err)
	}
	if err := doc.Filter(os.Stdout); err != nil {
		t.Error(err)
	}
}
