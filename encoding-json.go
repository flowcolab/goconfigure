package goconfigure

import (
	"encoding/json"
	"fmt"
	"io"
)

var configEncodingJson ConfigEncoding

func init() {
	configEncodingJson = NewConfigEncoding(
		func(reader io.Reader) ConfigDecoder { return NewJsonConfigDecoder(reader) },
		func(writer io.Writer) ConfigEncoder { return NewJsonConfigEncoder(writer) },
	)
}

func GetConfigEncodingJson() ConfigEncoding {
	return configEncodingJson
}

type JsonConfigDecoder struct {
	reader io.Reader
}

type JsonConfigEncoder struct {
	writer io.Writer
}

func NewJsonConfigDecoder(reader io.Reader) JsonConfigDecoder {
	return JsonConfigDecoder{reader}
}

func (dec JsonConfigDecoder) ReadProperties() (props *ConfigProperties, err error) {
	decoder := json.NewDecoder(dec.reader)

	props = &ConfigProperties{}
	if err = decoder.Decode(props); err != nil {
		err = fmt.Errorf("error decoding properties: %w", err)
	}

	return
}

func NewJsonConfigEncoder(writer io.Writer) JsonConfigEncoder {
	return JsonConfigEncoder{writer}
}

func (enc JsonConfigEncoder) WriteProperties(props *ConfigProperties) (err error) {
	encoder := json.NewEncoder(enc.writer)

	if err = encoder.Encode(props); err != nil {
		err = fmt.Errorf("error encoding properties: %w", err)
	}

	return
}
