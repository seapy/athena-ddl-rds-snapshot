package parser

import (
	"encoding/json"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/model"
	"io/ioutil"
	"log"
)

func Parse(path string) []model.Table {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	var exportStatus model.ExportStatus
	json.Unmarshal(data, &exportStatus)
	return exportStatus.Tables
}