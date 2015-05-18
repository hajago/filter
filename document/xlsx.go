package document

import (
	"archive/zip"
	"os"

	"github.com/hajago/filter/filetype"
	"gopkg.in/xmlpath.v2"
)

type Xlsx struct {
	fileName string
}

func NewXlsx(fileName string) *Xlsx {
	return &Xlsx{fileName: fileName}
}

func (d *Xlsx) Close() {

}

func (d *Xlsx) FileType() filetype.Type {
	return filetype.XLSX
}

func (d *Xlsx) Filter(oFile *os.File) error {
	r, err := zip.OpenReader(d.fileName)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == "xl/sharedStrings.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			pPath := xmlpath.MustCompile("//t")
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
