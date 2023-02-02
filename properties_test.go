package goconfigure

import (
	"testing"
)

func TestConfigProperties_SetIndex_WhenIndexFound_ThenSetValueOnExistingIndex(t *testing.T) {
	props := CONFIG_COMPLEX_PROPERTIES_FACTORY()
	if err := props.SetIndex([]string{"server", "port"}, 123); err != nil {
		t.Fatalf("error setting index")
	}

	val, ok := props.FindByIndex([]string{"server", "port"})
	if !ok {
		t.Fatalf("index not found")
	}
	if val != 123 {
		t.Fatalf("index set to incorect value")
	}
}

func TestConfigProperties_SetIndex_WhenIndexNotFound_ThenSetValueOnNewIndex(t *testing.T) {
	props := CONFIG_COMPLEX_PROPERTIES_FACTORY()
	if err := props.SetIndex([]string{"new1", "new2"}, 123); err != nil {
		t.Fatalf("error setting index")
	}

	val, ok := props.FindByIndex([]string{"new1", "new2"})
	if !ok {
		t.Fatalf("index not found")
	}
	if val != 123 {
		t.Fatalf("index set to incorect value")
	}
}

func TestConfigProperties_FindByIndex_WhenNotFound_ThenReturnNil(t *testing.T) {
	props := CONFIG_COMPLEX_PROPERTIES_FACTORY()
	if val, ok := props.FindByIndex([]string{"noop", "noop"}); ok || val != nil {
		t.Fatalf("incorrect return")
	}
}

func TestConfigProperties_FindByIndex_WhenIndexFound_ThenReturnValue(t *testing.T) {
	props := CONFIG_COMPLEX_PROPERTIES_FACTORY()
	if val, ok := props.FindByIndex([]string{"database", "employees", "postgres", "url"}); val == nil || !ok {
		t.Fatalf("incorrect return")
	} else if strVal, strValOk := val.(string); !ok || !strValOk || strVal != "postgres://localhost:12345/employees_db" {
		t.Fatalf("incorrect return")
	}
}

func TestConfigProperties_IsSameAs_WhenSamePointer_ThenReturnTrue(t *testing.T) {
	props := CONFIG_COMPLEX_PROPERTIES_FACTORY()

	if ok := props.IsSameAs(props); !ok {
		t.Fatalf("Same properties reported as not being the same.")
	}
}

func TestConfigProperties_IsSameAs_WhenSameContents_ThenReturnTrue(t *testing.T) {
	props1 := CONFIG_COMPLEX_PROPERTIES_FACTORY()
	props2 := CONFIG_COMPLEX_PROPERTIES_FACTORY()

	if ok := props1.IsSameAs(props2); !ok {
		t.Fatalf("Same properties reported as not being the same.\n%s\n%s",
			props1.Difference(props2).String(), props2.Difference(props1).String())
	}
}
