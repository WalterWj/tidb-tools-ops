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
	"strings"

	// import mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
)

const (
	configQ = "select @@tidb_config;"
)

// killsCmd represents the kills command
var killsCmd = &cobra.Command{
	Use:   "kills",
	Short: "Kill TIDB session",
	Long:  `Kill TiDB session`,
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", "mysql?charset=utf8"}, "")
		db, err := sql.Open("mysql", path)
		if err != nil {
			fmt.Println("connect is fail")
		}
		rows, err := db.Query(configQ)
		if err != nil {
			fmt.Printf("execute %v fail", userQ)
		}
		var configC string
		for rows.Next() {
			err := rows.Scan(&configC)
			if err != nil {
				fmt.Println("error is ", err)
			}
			fmt.Println(configC)
			// configC := json.Unmarshal(configC)
		}
		fmt.Println("kills called")
	},
}

func init() {
	rootCmd.AddCommand(killsCmd)

	killsCmd.Flags().StringVarP(&username, "user", "u", "root", "Database user")
	killsCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "Database host")
	killsCmd.Flags().StringVarP(&password, "password", "p", "123456", "Database passowrd")
	killsCmd.Flags().StringVarP(&port, "port", "P", "4000", "Database Port")
}
