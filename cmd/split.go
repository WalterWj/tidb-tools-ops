// Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split table",
	Short: "split table a_tmp like a",
	Long:  `跑批过程中，split 临时表，按照老表的结构进行解析，生产 split 语句.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("split called")
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// splitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// splitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	splitCmd.Flags().StringVarP(&dbusername, "user", "u", "root", "Username with privileges to select information_schema.tikv_region_status")
	splitCmd.Flags().StringVarP(&dbhost, "host", "H", "127.0.0.1", "Database host")
	splitCmd.Flags().StringVarP(&dbpassword, "password", "p", "123456", "Database passowrd")
	splitCmd.Flags().StringVarP(&dbtable, "table", "t", "", "table names, eg: db.table")
	splitCmd.Flags().StringVarP(&dbtable, "to_table", "T", "", "table names, eg: db.table_tmp")
	splitCmd.Flags().IntVarP(&dbport, "port", "P", 4000, "Database Port")

}
