package gengo

import (
	"fmt"
	"strings"

	"github.com/crazytyper/swagme"
)

type enumType struct {
	Name        string
	Description string
	Type        swagme.DataType
	Values      []string
}

// buildEnums scans the definitions for enums and tries to build common enum types for them.
// enums have no names in swagger, so the name is deferred from the property name the
// enum is used for and (if not unique) from the enclosing type.
func buildEnums(definitions swagme.Definitions) []enumType {
	enumsByName := map[string]enumType{}

	buildEnumName := func(outerName, propName string) string {
		propName = publicName(propName)

		if _, ok := enumsByName[propName]; !ok {
			return propName
		}

		outerName = publicName(outerName)
		outerName = stripWellKnownSuffixes(outerName)
		outerName = strings.TrimSuffix(outerName, propName) // do not stutter
		return outerName + propName
	}

	buildEnumsForProperty := func(outerName, propertyName string, property *swagme.Property) {
		if property.Enum == nil {
			return
		}
		if property.Ref != "" {
			panic(fmt.Errorf("Enums using type refs not supported. Property %q, ref %q", propertyName, property.Ref))
		}
		if t := typeName(property.Type); t == "" {
			panic(fmt.Errorf("Enums using type %q are supported. Property %q", property.Type, propertyName))
		}

		name := buildEnumName(outerName, propertyName)

		if enum, ok := enumsByName[name]; !ok {
			enumsByName[name] = enumType{
				Name:        name,
				Description: property.Description,
				Type:        property.Type,
				Values:      []string(property.Enum),
			}
			// tweak the type of the property
			property.Type = ""
			property.Ref = swagme.Ref("/" + name)
		} else {
			if enum.Type != "" && enum.Type != property.Type {
				panic(fmt.Errorf("Enum type mismatch. Enum %q was found with types %s and %s", name, enum.Type, property.Type))
			}
			if strings.Join(enum.Values, "|") != strings.Join([]string(property.Enum), "|") {
				panic(fmt.Errorf("Enum values mismatch. Enum %q was found with values %v but also with values %v", name, enum.Values, property.Enum))
			}
		}
	}

	for name, definition := range definitions {
		switch definition.Type {
		case swagme.TObject:
			for propname, property := range definition.Properties {
				buildEnumsForProperty(name, propname, property)
			}
		}
	}

	enums := []enumType{}
	for _, enum := range enumsByName {
		enums = append(enums, enum)
	}
	return enums
}
