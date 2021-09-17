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
	dbhost, dbname, dbusername, dbpassword string
	dbport, dbStatusPort                   int
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
		// table name
		tbn := common.GetTables(db, strconv.Quote(dbname))
		dbn := common.ParserDb(db, dbname)
		// db information
		common.Addfile(schemaFile, fmt.Sprintf("-- DB %s info", dbname))
		common.Addfile(schemaFile, dbn)
		common.Addfile(schemaFile, fmt.Sprintf("use %s;", dbname))
		// tables information
		for _, tableName := range tbn {
			tableMap := common.ParserTables(db, dbname, tableName)
			// tables
			common.Addfile(schemaFile, fmt.Sprintf("\n-- Table %s schema", tableName))
			common.Addfile(schemaFile, tableMap+";")
			// stats
			statsContent := common.ParserTs(dbhost, dbStatusPort, dbname, tableName)
			statsFile := filepath.Join(dir, "stats", fmt.Sprintf("%s.%s.json", dbname, tableName))
			common.Addfile(statsFile, statsContent)
			common.Addfile(schemaFile, fmt.Sprintf("\nLOAD STATS '%s';", statsFile))
		}
	},
}

func init() {
	rootCmd.AddCommand(statsdumpCmd)

	statsdumpCmd.Flags().StringVarP(&dbusername, "dbusername", "u", "root", "Database user")
	statsdumpCmd.Flags().StringVarP(&dbname, "dbname", "d", "test", "Database name")
	statsdumpCmd.Flags().StringVarP(&dbhost, "dbhost", "H", "127.0.0.1", "Database host")
	statsdumpCmd.Flags().StringVarP(&dbpassword, "dbpassword", "p", "123456", "Database passowrd")
	statsdumpCmd.Flags().IntVarP(&dbport, "dbport", "P", 4000, "Database Port")
	statsdumpCmd.Flags().IntVarP(&dbStatusPort, "statusport", "s", 10080, "TiDB Status Port")
}
