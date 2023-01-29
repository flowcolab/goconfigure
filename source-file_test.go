package goconfigure

import (
	"testing"
)

func TestFileConfigSource_GetProperties(t *testing.T) {
	NewFileConfigSource(CONFIG_COMPLEX_FILE, GetConfigEncodingJson())
}
