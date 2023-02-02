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

//func TestNewJsonConfigDecoder_ReadProperties_WhenReaderContainsJson_ThenReturnsCorrectProperties(t *testing.T) {
//	reader := strings.NewReader(CONFIG_COMPLEX_STRING)
//	decoder := NewJsonConfigDecoder(reader)
//	expectedProps := CONFIG_COMPLEX_PROPERTIES_FACTORY()
//	if props, err := decoder.ReadProperties(); err != nil {
//		t.Fatalf("error reading properties: %s", err.Error())
//	} else if ok := reflect.DeepEqual(props, expectedProps); !ok {
//		t.Fatalf(`unexpected properties`)
//	}
//}
