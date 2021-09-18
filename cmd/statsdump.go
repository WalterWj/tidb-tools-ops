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
	"strconv"
	"strings"
	"time"

	// import common
	"github.com/WalterWj/tidb-tools-ops/common"

	"github.com/spf13/cobra"
)

var (
	dbhost, dbname, dbusername, dbpassword, dbtable string
	dbport, dbStatusPort, mode                      int
)

const (
	tablesQ = "show tables"
)

// statsdumpCmd represents the statsdump command
var statsdumpCmd = &cobra.Command{
	Use:   "statsdump",
	Short: "Export statistics and table structures",
	Long:  `Export statistics and table structures`,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", "mysql?charset=utf8"}, "")
		db := mysqlConnect(dsn)
		// mkdir dir
		dir := strings.Join([]string{"stats-", time.Now().Format("2006-01-02-15:04:05")}, "")
		err := os.Mkdir(dir, os.ModePerm)
		ifErrLog(err)
		statsDir := filepath.Join(dir, "stats")
		err = os.MkdirAll(statsDir, os.ModePerm)
		ifErrLog(err)
		schemaFile := filepath.Join(dir, "schema.sql")
		// tidb version
		vs := common.GetVersion(db)
		common.Addfile(schemaFile, `/*`)
		common.Addfile(schemaFile, vs[0])
		common.Addfile(schemaFile, `*/`)
		// parser args
		if len(dbname) == 0 {
			if len(dbtable) == 0 {
				// Get all
				dblist := common.GetAllDb(db, mode)
				for _, dbTmp := range dblist {
					// Write db info
					wDbInfo(db, schemaFile, dbTmp)
					// table name
					tbn := common.GetTables(db, strconv.Quote(dbname))
					// Write tables information
					for _, tableName := range tbn {
						wTableInfo(db, schemaFile, dbname, tableName)
					}
				}
			} else {
				// get tables
				tablelist := common.ParserTbArgs(dbtable)
				for dbTmp, tbTmp := range tablelist {
					// write db info
					wDbInfo(db, schemaFile, dbTmp)
					// table name
					for _, tb := range tbTmp {
						// write table info
						wTableInfo(db, schemaFile, dbTmp, tb)
					}
				}
			}
		} else {
			// get databases
			dbTmp := common.ParserDbArgs(dbname)
			for _, dbName := range dbTmp {
				// write db info
				wDbInfo(db, schemaFile, dbName)
				// tablme name
				tbName := common.GetTables(db, dbName)
				for _, tb := range tbName {
					// write table info
					wTableInfo(db, schemaFile, dbName, tb)
				}
			}
		}

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
	tableMap := common.ParserTables(db, dbname, tbName)
	// tables
	common.Addfile(fileName, fmt.Sprintf("\n-- Table %s schema", tbName))
	common.Addfile(fileName, tableMap+";")
	// stats
	statsContent := common.ParserTs(dbhost, dbStatusPort, dbname, tbName)
	statsFile := filepath.Join(fileName, "stats", fmt.Sprintf("%s.%s.json", dbname, tbName))
	common.Addfile(statsFile, statsContent)
	common.Addfile(tbName, fmt.Sprintf("\nLOAD STATS '%s';", statsFile))
}

func init() {
	rootCmd.AddCommand(statsdumpCmd)

	statsdumpCmd.Flags().StringVarP(&dbusername, "dbusername", "u", "root", "Database user")
	statsdumpCmd.Flags().StringVarP(&dbname, "dbname", "d", "", "Database name, eg: db1,db2,db3")
	statsdumpCmd.Flags().StringVarP(&dbhost, "dbhost", "H", "127.0.0.1", "Database host")
	statsdumpCmd.Flags().StringVarP(&dbpassword, "dbpassword", "p", "123456", "Database passowrd")
	statsdumpCmd.Flags().StringVarP(&dbtable, "dbtable", "t", "", "table names, eg: db1.table1,db1.table2,db2.table3")
	statsdumpCmd.Flags().IntVarP(&dbport, "dbport", "P", 4000, "Database Port")
	statsdumpCmd.Flags().IntVarP(&dbStatusPort, "statusport", "s", 10080, "TiDB Status Port")
	statsdumpCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Ignore system database")
}
