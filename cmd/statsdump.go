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
	"strings"
	"time"

	// import common
	"github.com/WalterWj/tidb-tools-ops/common"

	"github.com/spf13/cobra"
)

var (
	dbhost, dbname, dbusername, dbpassword string
	dbport                                 int
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
		dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", "mysql?charset=utf8"}, "")
		db := mysqlConnect(dsn)
		dir := strings.Join([]string{"stats-", dbname, "-", time.Now().Format("2006-01-02-15:04:05")}, "")
		err := os.Mkdir(dir, os.ModePerm)
		ifErrLog(err)
		err = os.MkdirAll(fmt.Sprintf("%s/stats", dir), os.ModePerm)
		ifErrLog(err)
		// tidb version
		vs := common.GetVersion(db)
		common.Addfile(fmt.Sprintf("%s/schema.sql", dir), `/*`)
		common.Addfile(fmt.Sprintf("%s/schema.sql", dir), vs[0])
		common.Addfile(fmt.Sprintf("%s/schema.sql", dir), `*/`)

		tbn := common.GetTables(db, "'test'")
		for _, tableName := range tbn {
			tableMap := common.ParserTables(db, dbname, tableName)
			tbc, suc := tableMap[tableName]
			if suc {
				common.Addfile(fmt.Sprintf("%s/schema.sql", dir), fmt.Sprintf("\n-- Table %s schema", tableName))
				common.Addfile(fmt.Sprintf("%s/schema.sql", dir), tbc+";")
			}
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
	statsdumpCmd.Flags().IntVarP(&dbport, "statusport", "s", 10080, "TiDB Status Port")
}
