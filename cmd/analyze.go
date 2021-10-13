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

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze table",
	Long:  `Collect statistics. For example: analyze table test;`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyze called")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	analyzeCmd.Flags().StringVarP(&dbusername, "user", "u", "root", "Database user")
	analyzeCmd.Flags().StringVarP(&dbname, "database", "d", "", "Database name, eg: db1,db2,db3")
	analyzeCmd.Flags().StringVarP(&dbhost, "host", "H", "127.0.0.1", "Database host")
	analyzeCmd.Flags().StringVarP(&dbpassword, "password", "p", "123456", "Database passowrd")
	analyzeCmd.Flags().StringVarP(&dbtable, "tables", "t", "", "table names, eg: db1.table1,db1.table2,db2.table3")
	analyzeCmd.Flags().IntVarP(&dbport, "port", "P", 4000, "Database Port")
	analyzeCmd.Flags().IntVarP(&thread, "thread", "T", 4, "Number of threads executing analyze")
	analyzeCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Ignore system database")
}
