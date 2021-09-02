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
	"io/ioutil"
	"strings"

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
		res := common.GetTables(db, "'test'")
		for _, tableName := range res {
			showQ := fmt.Sprintf("show create table %s", tableName)
			db.Exec(fmt.Sprintf("use %s;", dbname))
			rows, err := db.Query(showQ)
			ifErrLog(err)
			for rows.Next() {
				var t, Ct string
				err := rows.Scan(&t, &Ct)
				ifErrLog(err)
				// fmt.Printf("%s;\n", Ct)
				ctc := []byte(Ct)
				err = ioutil.WriteFile(fmt.Sprintf("%s-%s.sql", dbname, tableName), ctc, 0644)
				ifErrLog(err)
			}
			rows.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(statsdumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsdumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsdumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	statsdumpCmd.Flags().StringVarP(&dbusername, "dbusername", "u", "root", "Database user")
	statsdumpCmd.Flags().StringVarP(&dbname, "dbname", "d", "test", "Database name")
	statsdumpCmd.Flags().StringVarP(&dbhost, "dbhost", "H", "127.0.0.1", "Database host")
	statsdumpCmd.Flags().StringVarP(&dbpassword, "dbpassword", "p", "123456", "Database passowrd")
	statsdumpCmd.Flags().IntVarP(&dbport, "dbport", "P", 4000, "Database Port")
	statsdumpCmd.Flags().IntVarP(&dbport, "statusport", "s", 10080, "TiDB Status Port")
}