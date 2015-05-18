package errors

import (
	"errors"
)

var (
	ErrNotSupportedFormat    = errors.New("Not supported format")
	ErrNotFoundMainXMLOfWord = errors.New("Not found main xml of word")
)
