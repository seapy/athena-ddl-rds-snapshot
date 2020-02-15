package ddl

import (
	"github.com/seapy/athena-ddl-rds-snapshot/internal/model"
	"log"
	"os"
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

func Output(athenaDB string, s3Prefix string, t *model.Table) {
	template, err := template.New("DDL").Funcs(fns).Parse(DdlTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	model := Ddl{
		AthenaDB: athenaDB,
		S3Prefix: s3Prefix,
		Table:    t,
	}
	template.ExecuteTemplate(os.Stdout, "DDL", model)
}
