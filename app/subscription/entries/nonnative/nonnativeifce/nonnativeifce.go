package nonnativeifce

import (
	"io/fs"

	"github.com/vdonkey/accelerator/v5/app/subscription/entries"
)

type NonNativeConverterConstructorT func(fs fs.FS) (entries.Converter, error)

var NewNonNativeConverterConstructor NonNativeConverterConstructorT
