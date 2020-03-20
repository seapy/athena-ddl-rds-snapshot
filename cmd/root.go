package cmd

import (
	"fmt"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/athena"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/ddl"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/model"
	"github.com/seapy/athena-ddl-rds-snapshot/internal/parser"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {
		tableInfos := parser.Parse(infoPath)
		for _, t := range tableInfos  {
			if len(tables) > 0 {
				for _, filterName := range tables {
					if t.AthenaName() == filterName {
						sqlProcess(t)
					}
				}
			} else {
				sqlProcess(t)
			}
		}
	},
}

var athenaDB string
var s3Prefix string
var infoPath string
var tables []string
var execute bool
var workgroup string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&athenaDB, "athenaDB", "a", "", "Your athena database name.")
	rootCmd.MarkFlagRequired("athenaDB")
	rootCmd.Flags().StringVarP(&s3Prefix, "s3Prefix", "s", "", "Exported snapshot s3 location. include database name. not include table name.")
	rootCmd.MarkFlagRequired("s3Prefix")
	rootCmd.Flags().StringVarP(&infoPath, "infoPath", "i", "", "Table exported information json file location. you can download from s3 exported result.")
	rootCmd.MarkFlagRequired("infoPath")
	rootCmd.Flags().StringSliceVarP(&tables, "tables", "t", nil, "Table name if you want DDL only some tables. Seprated by comma(,)")
	rootCmd.Flags().BoolVarP(&execute, "execute", "e", false, "Execute sql to Athena")
	rootCmd.Flags().StringVarP(&workgroup, "workgroup", "w", "", "Athena workgroup. Required if execute is true")
}

func sqlProcess(t model.Table) {
	sql, err := ddl.Sql(athenaDB, s3Prefix, &t)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("🚀 %s \n", t.Name)
	if execute {
		_, err := athena.Update(sql, workgroup)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println(sql)
	}
}