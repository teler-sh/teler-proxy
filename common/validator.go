package common

func (opt *Options) Validate() error {
	validFormats := map[string]bool{
		"yaml": true,
		"json": true,
	}

	if opt.Destination == "" {
		return ErrDestAddressEmpty
	}

	if opt.Config.Path != "" && !validFormats[opt.Config.Format] {
		return ErrCfgFileFormatInv
	}

	return nil
}
