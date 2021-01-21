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
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	host, username, password, port string
)

const (
	userQ = "select user,host,authentication_string from user;"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", "mysql?charset=utf8"}, "")
		db, err := sql.Open("mysql", path)
		if err != nil {
			fmt.Println("connect is fail")
		}
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

			fmt.Println(createuser)
			fmt.Println(userinfo)

			// ierr := ioutil.WriteFile("user.sql", []byte(createuser), 0644)
			// if ierr != nil {
			// 	fmt.Println("error is: ", ierr)
			// }
			gRows, err := db.Query(grantQ)
			if err != nil {
				fmt.Printf("execute %v fail", grantQ)
			}
			for gRows.Next() {
				err := gRows.Scan(&grant)
				if err != nil {
					fmt.Println("error is ", err)
				}
				fmt.Println(grant, ";")
			}
		}

		fmt.Println("Successfully introduce all users and permissions.")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	exportCmd.Flags().StringVarP(&username, "dbusername", "u", "root", "Database user")
	exportCmd.Flags().StringVarP(&host, "dbhost", "H", "127.0.0.1", "Database host")
	exportCmd.Flags().StringVarP(&password, "dbpassword", "p", "123456", "Database passowrd")
	exportCmd.Flags().StringVarP(&port, "dbport", "P", "4000", "Database Port")
	// exportCmd.Flags().IntVarP(&port, "statusport", "s", 10080, "TiDB Status Port")

}

func addfile(context string) {
	file, err := os.OpenFile("users.sql", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file fail:", err)
	}
}
