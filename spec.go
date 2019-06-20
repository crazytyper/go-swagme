package swagme

// DataType ...
type DataType string

// Data types
const (
	TString  DataType = "string"
	TInteger          = "integer"
	TBoolean          = "boolean"
	TObject           = "object"
	TArray            = "array"
)

// AdditionalProperties ...
type AdditionalProperties struct {
	Type DataType `json:"type"`
	Ref  Ref      `json:"$ref,omitempty"`
}

// Ref ...
type Ref string

// ArrayItems ...
type ArrayItems struct {
	Type DataType `json:"type"`
	Ref  Ref      `json:"$ref,omitempty"`
}

// Enum ...
type Enum []string

// Property ...
type Property struct {
	Type                 DataType              `json:"type"`
	Ref                  Ref                   `json:"$ref,omitempty"`
	Description          string                `json:"description,omitempty"`
	Enum                 Enum                  `json:"enum,omitempty"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
	Items                *ArrayItems           `json:"items,omitempty"`
}

// Definition ...
type Definition struct {
	Type       DataType             `json:"type"`
	Properties map[string]*Property `json:"properties,omitempty"`
}

// Definitions ...
type Definitions map[string]*Definition

// Info ...
type Info struct {
	Version string `json:"version"`
	Title   string `json:"title"`
}

// Spec ...
type Spec struct {
	Info        *Info       `json:"info,omitempty"`
	Definitions Definitions `json:"definitions,omitempty"`
}
