package main

import (
	"github.com/hajago/filter/filetype"
	"os"
)

type Document interface {
	FileType() filetype.Type
	Filter(oFile *os.File) error
}
