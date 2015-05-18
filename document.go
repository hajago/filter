package main

import (
	"os"

	"github.com/hajago/filter/filetype"
)

type Document interface {
	FileType() filetype.Type
	Filter(oFile *os.File) error
	Close()
}
