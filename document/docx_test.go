package document

import (
	"os"
	"testing"
)

func TestDocxFilter(t *testing.T) {
	doc := NewDocx("../Files/test.docx")
	if err := doc.Filter(os.Stdout); err != nil {
		t.Error(err)
	}
}
