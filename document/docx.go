package document

import (
	"archive/zip"
	"os"
	"strings"

	"github.com/hajago/filter/errors"
	"github.com/hajago/filter/filetype"
	"github.com/hajago/filter/types"
	"gopkg.in/xmlpath.v2"
)

type Docx struct {
	fileName   string
	fileReader *zip.ReadCloser
	fileMap    map[string]*zip.File
}

func NewDocx(fileName string) (docx *Docx, err error) {
	docx = &Docx{fileName: fileName, fileMap: map[string]*zip.File{}}
	docx.fileReader, err = zip.OpenReader(fileName)
	if err != nil {
		return
	}

	for _, f := range docx.fileReader.File {
		docx.fileMap[f.Name] = f
	}
	return
}

func (d *Docx) Close() {
	d.fileReader.Close()
}

func (d *Docx) FileType() filetype.Type {
	return filetype.DOCX
}

func (d *Docx) Filter(oFile *os.File) error {
	fName, err := d.documentMainXMLName()
	if err != nil {
		return err
	}
	if f, ok := d.fileMap[fName]; ok {
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
	return nil
}

func (d *Docx) documentMainXMLName() (string, error) {
	if f, ok := d.fileMap[types.ContentTypesXMLFileName]; ok {
		contentType := types.ContentType{}
		if err := types.UnmarshalZipFile(f, &contentType); err != nil {
			return "", err
		}

		for _, o := range contentType.Override {
			if o.ContentType == types.ContentTypeWord {
				return strings.TrimPrefix(o.PartName, "/"), nil
			}
		}
	}
	return "", errors.ErrNotFoundMainXMLOfWord
}
