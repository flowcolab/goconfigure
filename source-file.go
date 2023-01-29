package goconfigure

import (
	"fmt"
	"os"
)

type FileConfigSource struct {
	path     string
	encoding ConfigEncoding
}

func NewFileConfigSource(path string, encoding ConfigEncoding) FileConfigSource {
	return FileConfigSource{path, encoding}
}

func (cs FileConfigSource) GetProperties() (props *ConfigProperties, err error) {
	var file *os.File
	file, err = os.Open(cs.path)
	if err != nil {
		err = fmt.Errorf("error opening properties file '%s': %w", cs.path, err)
		return
	}

	decoder := cs.encoding.decoderFactory(file)
	if props, err = decoder.ReadProperties(); err != nil {
		err = fmt.Errorf("error decoding properties file '%s': %w", cs.path, err)
		return
	}

	return
}
