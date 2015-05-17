package document

import (
	"os"
	"testing"
)

func TestXlsxFilter(t *testing.T) {
	doc := NewXlsx("../Files/test.xlsx")
	if err := doc.Filter(os.Stdout); err != nil {
		t.Error(err)
	}
}
