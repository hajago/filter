package document

import (
	"archive/zip"
	"os"

	"github.com/hajago/filter/filetype"
	"gopkg.in/xmlpath.v2"
)

type Docx struct {
	fileName string
}

func NewDocx(fileName string) *Docx {
	return &Docx{fileName: fileName}
}

func (d *Docx) FileType() filetype.Type {
	return filetype.DOCX
}

func (d *Docx) Filter(oFile *os.File) error {
	r, err := zip.OpenReader(d.fileName)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			pPath := xmlpath.MustCompile("//p")
			pRoot, err := xmlpath.Parse(rc)
			if err != nil {
				return err
			}
			pIter := pPath.Iter(pRoot)
			for pIter.Next() {
				byts := pIter.Node().Bytes()
				if len(byts) > 0 {
					if _, err := oFile.Write(byts); err != nil {
						return err
					}
					oFile.Write([]byte("\n"))
				}
			}
		}
	}
	return nil
}
