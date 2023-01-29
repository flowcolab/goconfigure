package goconfigure

import (
	"strings"
	"testing"
)

func TestJsonConfigDecoder_ReadProperties_WhenReaderEmpty_ThenReturnsEmptyProperties(t *testing.T) {
	reader := strings.NewReader(CONFIG_EMPTY_STRING)
	decoder := NewJsonConfigDecoder(reader)
	if props, err := decoder.ReadProperties(); err != nil {
		t.Fatalf("error reading properties: %s", err.Error())
	} else if len(*props) > 0 {
		t.Fatalf("unexpected properties; properties should be empty;")
	}
}

func TestNewJsonConfigDecoder_ReadProperties_WhenReaderContainsJson_ThenReturnsCorrectProperties(t *testing.T) {
	reader := strings.NewReader(CONFIG_COMPLEX_STRING)
	decoder := NewJsonConfigDecoder(reader)
	if props, err := decoder.ReadProperties(); err != nil {
		t.Fatalf("error reading properties: %s", err.Error())
	} else if diff := props.Difference(&CONFIG_COMPLEX_PROPERTIES); diff.HasProperties() {
		t.Fatalf("unexpected properties:\n %s", diff.String())
	}
}
