package gengo

import (
	"strings"
	"unicode"

	"github.com/crazytyper/swagme"
)

func publicName(name string) string {
	return fixAbbrevatedSuffixes(strings.ToUpper(name[:1]) + name[1:])
}

func fixAbbrevatedSuffixes(name string) string {
	for _, suffix := range []string{"Id", "Url"} {
		if strings.HasSuffix(name, suffix) {
			name = strings.TrimSuffix(name, suffix)
			if name == "" {
				return strings.ToUpper(suffix)
			}

			lastchar := name[len(name)-1] // <- will not work with unicode
			if unicode.IsLower(rune(lastchar)) {
				return name + strings.ToUpper(suffix)
			}

			return name
		}
	}
	return name
}

func typeName(t swagme.DataType) string {
	switch t {
	case swagme.TString:
		return "string"
	case swagme.TBoolean:
		return "bool"
	case swagme.TInteger:
		return "int"
	}
	return ""
}

func typeNameForRef(ref swagme.Ref) string {
	if p := strings.LastIndex(string(ref), "/"); p >= 0 {
		return string(ref)[p+1:]
	}
	return ""
}

func stripWellKnownSuffixes(name string) string {
	for _, suffix := range []string{"Model", "Request", "Response"} {
		name = strings.TrimSuffix(name, suffix)
	}
	return name
}
