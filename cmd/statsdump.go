// Copyright © 2021 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	// import common
	"github.com/WalterWj/tidb-tools-ops/common"

	"github.com/spf13/cobra"
)

// statsdumpCmd represents the statsdump command
var statsdumpCmd = &cobra.Command{
	Use:   "statsdump",
	Short: "Export statistics and table structures",
	Long:  `Export statistics and table structures`,
	Run: func(cmd *cobra.Command, args []string) {
		// get start time
		st := time.Now()

		// mkdir dir
		dir := strings.Join([]string{"stats-", time.Now().Format("2006-01-02-15:04:05")}, "")
		err := os.Mkdir(dir, os.ModePerm)
		common.IfErrLog(err)

		// mkdir stats dir
		statsDir := filepath.Join(dir, "stats")
		err = os.MkdirAll(statsDir, os.ModePerm)
		common.IfErrLog(err)

		// mkdir schema file
		schemaFile := filepath.Join(dir, "schema.sql")

		// connect db
		dsn := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhost, ":", fmt.Sprint(dbport), ")/", "mysql?charset=utf8"}, "")
		db := common.MysqlConnect(dsn)

		// tidb version
		vs := common.GetVersion(db)
		common.Addfile(schemaFile, `/*`)
		common.Addfile(schemaFile, vs[0])
		common.Addfile(schemaFile, `*/`)

		// table list
		var tableList = make(map[string][]string)

		// parser args, make db & table: tableList = map[string][]string
		if len(dbname) == 0 {
			// args database is null, args tables is null:
			if len(dbtable) == 0 {
				// Get all
				dblist := common.GetAllDb(db, mode)
				for _, dbTmp := range dblist {
					// get all database and tables:
					tableList[dbTmp] = common.MapToArryString(common.GetTables(db, dbTmp))
				}
			} else {
				// args: database is null, tables is not null:
				// get tables
				tableList = common.ParserTbArgs(dbtable)
			}
		} else {
			// args database is not null:
			// get databases
			dblist := common.ParserDbArgs(dbname)
			for _, dbTmp := range dblist {
				tableList[dbTmp] = common.MapToArryString(common.GetTables(db, dbTmp))
			}
		}

		// Write file
		for _db, _tbname := range tableList {
			// db schema,
			wDbInfo(db, schemaFile, _db)
			for _, _tb := range _tbname {
				// table schema
				wTableInfo(db, schemaFile, _db, _tb)
				// table states
				wStatsInfo(statsDir, dbhost, dbStatusPort, _db, _tb)
				fmt.Printf("Get %s.%s stats Succeeded~\n", _db, _tb)
			}
		}

		// Close database connection
		defer db.Close()
		// get end time
		et := time.Now()
		fmt.Println("Get All stats Succeeded!")
		// total cost time
		fmt.Printf("Cost time is: %s\n", et.Sub(st))
	},
}

func wDbInfo(db *sql.DB, fileName string, dbName string) {
	// db information
	common.Addfile(fileName, fmt.Sprintf("-- DB %s info", dbName))
	dbn := common.ParserDb(db, dbName)
	common.Addfile(fileName, dbn)
	common.Addfile(fileName, fmt.Sprintf("use %s;", dbName))
}

// Write table information to file
func wTableInfo(db *sql.DB, fileName string, dbName string, tbName string) {
	tableMap := common.ParserTables(db, dbName, tbName)
	// tables
	common.Addfile(fileName, fmt.Sprintf("\n-- Table %s schema", tbName))
	common.Addfile(fileName, tableMap+";")
	common.Addfile(fileName, fmt.Sprintf("\nLOAD STATS 'stats/%s.%s.json;'", dbName, tbName))
}

func wStatsInfo(fileName string, dbHost string, dbStatusPort int, dbName string, tbName string) {
	// stats
	statsContent := common.ParserTs(dbHost, dbStatusPort, dbName, tbName)
	statsFile := filepath.Join(fileName, fmt.Sprintf("%s.%s.json", dbName, tbName))
	common.Addfile(statsFile, statsContent)
}

func init() {
	rootCmd.AddCommand(statsdumpCmd)

	statsdumpCmd.Flags().StringVarP(&dbusername, "user", "u", "root", "Database user")
	statsdumpCmd.Flags().StringVarP(&dbname, "database", "d", "", "Database name, eg: db1,db2,db3")
	statsdumpCmd.Flags().StringVarP(&dbhost, "host", "H", "127.0.0.1", "Database host")
	statsdumpCmd.Flags().StringVarP(&dbpassword, "password", "p", "123456", "Database passowrd")
	statsdumpCmd.Flags().StringVarP(&dbtable, "tables", "t", "", "table names, eg: db1.table1,db1.table2,db2.table3")
	statsdumpCmd.Flags().IntVarP(&dbport, "port", "P", 4000, "Database Port")
	statsdumpCmd.Flags().IntVarP(&dbStatusPort, "statusport", "s", 10080, "TiDB Status Port")
	statsdumpCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Ignore system database, eg: 1 (default 0)")
}
