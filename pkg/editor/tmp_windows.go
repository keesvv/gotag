//go:build windows

package editor

import (
	"os"
	path "path/filepath"
)

var TMP_PATH = path.Join(os.Getenv("LOCALAPPDATA"), "Temp", "gotag")
