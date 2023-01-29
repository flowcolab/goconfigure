package goconfigure

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type ConfigProperties map[string]any

type ConfigPropertiesVisitor func(indices []string, memberValue any)

func (props *ConfigProperties) HasProperties() bool {
	return len(*props) > 0
}

func (props *ConfigProperties) String() string {
	writer := strings.Builder{}
	encoder := json.NewEncoder(&writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(props); err == nil {
		return writer.String()
	} else {
		return "error encoding properties"
	}
}

func (props *ConfigProperties) Difference(other *ConfigProperties) *ConfigProperties {
	diffProps := ConfigProperties{}

	props.Visit(func(indices []string, memberValue any) {
		// if current index value is not found on other properties then add it to the difference
		if otherValue, otherOk := other.FindByIndex(indices); !otherOk || otherValue != memberValue {
			if err := diffProps.SetIndex(indices, memberValue); err != nil {
				panic(fmt.Sprintf("error setting difference at index '%s' to value '%v': %s",
					strings.Join(indices, "."), memberValue, err.Error()))
			}
		}
	})

	return &diffProps
}

func (props *ConfigProperties) SetIndex(indices []string, value any) error {
	var currentValue any = *props
	var currentIndex = make([]string, len(indices))

	for index, propIndex := range indices {
		currentIndex = append(currentIndex, propIndex)

		// if last index then set the value on map otherwise navigate to the next map
		if index == len(indices)-1 {
			if props, ok := currentValue.(ConfigProperties); ok {
				props[propIndex] = value
			} else {
				return fmt.Errorf("could not set properties at index '%s' because index is not a map",
					strings.Join(indices[:len(indices)-1], "."))
			}
		} else {
			if props, ok := currentValue.(ConfigProperties); ok {
				props[propIndex] = ConfigProperties{}
				currentValue = props[propIndex]
			} else {
				return fmt.Errorf("could not set properties at index '%s' because index is not a map",
					strings.Join(indices, "."))
			}
		}
	}

	return nil
}

func (props *ConfigProperties) Visit(visitor ConfigPropertiesVisitor) {
	visitNode([]string{}, props, visitor)
}

func (props *ConfigProperties) VisitIndex(indices []string, visitor ConfigPropertiesVisitor) {
	if value, ok := props.FindByIndex(indices); ok {
		visitNode(indices, value, visitor)
	}
}

func visitNode(currentIndices []string, value any, visitor ConfigPropertiesVisitor) {
	if value == nil {
		return
	}

	if valueValue := reflect.ValueOf(value); valueValue.Kind() == reflect.Ptr {
		visitNode(currentIndices, valueValue.Elem().Interface(), visitor)
		return
	}

	if valueAsMap, ok := value.(map[string]any); ok {
		for propName, propValue := range valueAsMap {
			visitNode(append(currentIndices, propName), propValue, visitor)
		}
		return
	}

	visitor(currentIndices, value)
}

func (props *ConfigProperties) SubProperties(indices []string) (subProps *ConfigProperties, ok bool) {
	if currentValue, ok := props.FindByIndex(indices); ok {
		if currentProps, ok := currentValue.(ConfigProperties); ok {
			return &currentProps, true
		}
	}

	return nil, false
}

func (props *ConfigProperties) FindByIndex(indices []string) (any, bool) {
	var currentValue any = props
	for _, propIndex := range indices {
		if currentValueValue := reflect.ValueOf(currentValue); currentValueValue.Kind() == reflect.Ptr {
			currentValue = currentValueValue.Elem().Interface()
		}

		if subProps, ok := currentValue.(ConfigProperties); ok {
			currentValue = subProps[propIndex]
		} else {
			return nil, false
		}
	}

	return currentValue, true
}
