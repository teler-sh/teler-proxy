package common

import "errors"

var (
	ErrCfgFileFormatInv = errors.New("invalid teler configuration file format")
	ErrCfgFileFormatUnd = errors.New("undefined teler configuration file format")
	ErrDestAddressEmpty = errors.New("empty destination address")
)
