/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analysisCmd represents the analysis command
var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "Analysis TiDB",
	Long:  `Analyze the bottleneck and top n SQL of the current tidb cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analysis called")
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analysisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analysisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	analysisCmd.Flags().StringVarP(&dbusername, "user", "u", "root", "Username with privileges to run the analyze")
	analysisCmd.Flags().StringVarP(&dbname, "database", "d", "", "Database name, eg: db1,db2,db3")
	analysisCmd.Flags().StringVarP(&dbhost, "host", "H", "127.0.0.1", "Database host")
	analysisCmd.Flags().StringVarP(&dbpassword, "password", "p", "123456", "Database passowrd")
	analysisCmd.Flags().StringVarP(&dbtable, "tables", "t", "", "table names, eg: db1.table1,db1.table2,db2.table3")
	analysisCmd.Flags().IntVarP(&dbport, "port", "P", 4000, "Database Port")
	analysisCmd.Flags().IntVarP(&thread, "thread", "T", 4, "Number of goroutines to use")
	analysisCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Ignore system database, eg: 1")
	analysisCmd.Flags().IntVarP(&stats_healthy, "healthy", "s", 100, "Table stats healthy, If it is below the threshold, then analyze")
}
