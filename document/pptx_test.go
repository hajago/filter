package document

import (
	"os"
	"testing"
)

func TestPptxFilter(t *testing.T) {
	doc := NewPptx("../Files/test.pptx")
	if err := doc.Filter(os.Stdout); err != nil {
		t.Error(err)
	}
}
