package goconfigure

import "io"

type ConfigEncoding struct {
	decoderFactory ConfigDecoderFactory
	encoderFactory ConfigEncoderFactory
}

type ConfigDecoderFactory func(reader io.Reader) ConfigDecoder

type ConfigEncoderFactory func(writer io.Writer) ConfigEncoder

type ConfigDecoder interface {
	ReadProperties() (*ConfigProperties, error)
}

type ConfigEncoder interface {
	WriteProperties(properties *ConfigProperties) error
}

func NewConfigEncoding(decoderFactory ConfigDecoderFactory, encoderFactory ConfigEncoderFactory) ConfigEncoding {
	return ConfigEncoding{
		decoderFactory,
		encoderFactory,
	}
}

func (enc ConfigEncoding) createDecoder(reader io.Reader) ConfigDecoder {
	return enc.decoderFactory(reader)
}

func (enc ConfigEncoding) createEncoder(writer io.Writer) ConfigEncoder {
	return enc.encoderFactory(writer)
}
