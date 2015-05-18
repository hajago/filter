package document

import (
	"archive/zip"
	"os"
	"strings"

	"github.com/hajago/filter/filetype"
	"gopkg.in/xmlpath.v2"
)

type Pptx struct {
	fileName     string
	slideMasters []slideMaster
	slideLayouts []slideLayout
}

type slideMaster struct {
	fileName string
}

type slideLayout struct {
	fileName string
}

type LayoutSlide struct {
	fileName string
}

func NewPptx(fileName string) *Pptx {
	return &Pptx{fileName: fileName}
}

func (d *Pptx) Close() {

}

func (d *Pptx) FileType() filetype.Type {
	return filetype.PPTX
}

func (d *Pptx) buildSlideMasters(file *zip.File) {

}

func (d *Pptx) buildSlideLayouts(file *zip.File) {

}

func (d *Pptx) filterSlide(oFile *os.File, f *zip.File) error {
	file, err := f.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	pPath := xmlpath.MustCompile("//t")
	pRoot, err := xmlpath.Parse(file)
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

	return nil
}

func (d *Pptx) Filter(oFile *os.File) error {
	r, err := zip.OpenReader(d.fileName)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if !strings.HasSuffix(f.Name, ".xml") {
			continue
		}

		if strings.HasPrefix(f.Name, "ppt/slideMasters") {
			d.buildSlideMasters(f)
		} else if strings.HasPrefix(f.Name, "ppt/slideLayouts") {
			d.buildSlideLayouts(f)
		} else if strings.HasPrefix(f.Name, "ppt/slides") {
			d.filterSlide(oFile, f)
		}
	}
	return nil
}
