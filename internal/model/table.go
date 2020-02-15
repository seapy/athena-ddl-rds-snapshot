package model

import "strings"

type ExportStatus struct {
	Tables []Table `json:"perTableStatus"`
}

type Table struct {
	Name string `json:"target"`
	SchemaMetaData SchemaMetaData `json:"schemaMetadata"`
}

func (t *Table) Columns() []Column {
	return t.SchemaMetaData.Columns
}

func (t *Table) AthenaName() string {
	names := strings.Split(t.Name, ".")
	return names[len(names)-1]
}

type SchemaMetaData struct {
	Columns []Column `json:"OriginalTypeMappings"`
}

type Column struct {
	Name string `json:"columnName"`
	OriginalType string `json:"originalType"`
	ExpectedExportedType string `json:"expectedExportedType"`
}

func (c *Column) AthenaType() string {
	switch c.ExpectedExportedType {
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "double":
		return "double"
	case "boolean":
		return "boolean"
	default:
		return "string"
	}
}