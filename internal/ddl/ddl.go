package ddl

import (
	"bytes"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/model"
	"reflect"
	"text/template"
)

type Ddl struct {
	AthenaDB string
	S3Prefix string
	Table    *model.Table
}

// https://stackoverflow.com/a/22375000/397457
var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func Sql(athenaDB string, s3Prefix string, t *model.Table) (string, error) {
	template, err := template.New("DDL").Funcs(fns).Parse(DdlTemplate)
	if err != nil {
		return "", err
	}
	model := Ddl{
		AthenaDB: athenaDB,
		S3Prefix: s3Prefix,
		Table:    t,
	}
	var buff bytes.Buffer
	err = template.ExecuteTemplate(&buff, "DDL", model)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}
