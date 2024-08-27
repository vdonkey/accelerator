//go:build !windows
// +build !windows

package platform

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

func LineSeparator() string {
	return "\n"
}

// GetAssetLocation search for `file` in certain locations
func GetAssetLocation(file string) string {
	const name = "accelerator.location.asset"
	assetPath := NewEnvFlag(name).GetValue(getExecutableDir)
	defPath := filepath.Join(assetPath, file)
	relPath := filepath.Join("accelerator", file)
	fullPath, err := xdg.SearchDataFile(relPath)
	if err != nil {
		return defPath
	}
	return fullPath
}
