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
	"strings"

	"tidb-tools-ops/common"

	// import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

const (
	userQ = "select user,host,authentication_string from user;"
	// userQ = "select user,host,password from user;"  v2.1
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export your users for TiDB",
	Long:  `export your users and passowrd for TiDB`,
	Run: func(cmd *cobra.Command, args []string) {
		dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", fmt.Sprint(port), ")/", "mysql?charset=utf8"}, "")
		db := common.MysqlConnect(dsn)
		rows, err := db.Query(userQ)
		if err != nil {
			fmt.Printf("execute %v fail", userQ)
		}
		var user, host, pas, grant string
		for rows.Next() {
			err := rows.Scan(&user, &host, &pas)
			if err != nil {
				fmt.Println("error is ", err)
			}
			createuser := strings.Join([]string{"create user ", user, "@", host, ";"}, "'")
			userinfo := strings.Join([]string{"update mysql.user set `authentication_string`=", pas, " where user=", user, " and host=", host, ";"}, "'")
			grantQ := strings.Join([]string{"SHOW GRANTS FOR ", user, "@", host, ";"}, "'")

			common.Addfile("users.sql", createuser)
			common.Addfile("users.sql", userinfo)
			gRows, err := db.Query(grantQ)
			if err != nil {
				fmt.Printf("execute %v fail", grantQ)
			}
			for gRows.Next() {
				err := gRows.Scan(&grant)
				if err != nil {
					fmt.Println("error is ", err)
				}
				grant = strings.Join([]string{grant, ";"}, "")

				common.Addfile("users.sql", grant)
			}
			common.Addfile("users.sql", "")
		}

		fmt.Println("Successfully introduce all users and permissions.")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&username, "user", "u", "root", "Database user")
	exportCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "Database host")
	exportCmd.Flags().StringVarP(&password, "password", "p", "123456", "Database passowrd")
	exportCmd.Flags().IntVarP(&port, "port", "P", 4000, "Database Port")

}
