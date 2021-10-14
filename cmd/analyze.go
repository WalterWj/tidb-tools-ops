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
	"sync"
	"time"

	"github.com/WalterWj/tidb-tools-ops/common"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze table",
	Long:  `Collect statistics. For example: analyze table test;`,
	Run: func(cmd *cobra.Command, args []string) {
		// get start time
		st := time.Now()
		// connect db
		dsn := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhost, ":", fmt.Sprint(dbport), ")/", "mysql?charset=utf8"}, "")
		db := common.MysqlConnect(dsn)

		// table list
		var tableList = make(map[string][]string)

		// parser args, make db & table: tableList = map[string][]string
		if len(dbname) == 0 {
			// args database is null, args tables is null:
			if len(dbtable) == 0 {
				// Get all
				dblist := common.GetAllDb(db, mode)
				for _, dbTmp := range dblist {
					// get all database and tables:
					tableList[dbTmp] = common.MapToArryString(common.GetTables(db, dbTmp))
				}
			} else {
				// args: database is null, tables is not null:
				// get tables
				tableList = common.ParserTbArgs(dbtable)
			}
		} else {
			// args database is not null:
			// get databases
			dblist := common.ParserDbArgs(dbname)
			for _, dbTmp := range dblist {
				tableList[dbTmp] = common.MapToArryString(common.GetTables(db, dbTmp))
			}
		}

		// thread
		// channel + waitgroup
		wg := sync.WaitGroup{}
		ch := make(chan struct{}, thread)
		for _dbname, _tblist := range tableList {
			for _, _tb := range _tblist {
				wg.Add(1) // 添加计数
				ch <- struct{}{}
				go func(_tb string) {
					rc := test(db, _dbname, _tb)
					defer fmt.Printf("[%s] analyze table: %s.%s Sucessfull \n", time.Unix(0, time.Now().UnixMilli()*1000000), dbname, _tb)
					if rc == 0 {
						defer wg.Done() // 将计数减1
						<-ch            // 读取chan
					} else {
						errC := strings.Join([]string{"execute analyze ", _dbname, ".", _tb, "failed"}, "")
						defer common.IfErrPrintE(errC)
					}
				}(_tb)
			}
		}
		// 等待加入的协程全部完成
		wg.Wait()
		// Close database connect
		defer db.Close()

		// get end time
		et := time.Now()
		fmt.Println("Analyze All tables Succeeded")
		// total cost time
		fmt.Printf("Total Cost time: %s\n", et.Sub(st))
	},
}

func analyzeTable(db *sql.DB, dbname string, tbname string) int64 {
	st, err := db.Exec(fmt.Sprintf("analyze table `%s`.`%s`", dbname, tbname))
	if err == nil {
		common.IfErrLog(err)
	}
	rs, _ := st.RowsAffected()

	return rs
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&dbusername, "user", "u", "root", "Username with privileges to run the analyze")
	analyzeCmd.Flags().StringVarP(&dbname, "database", "d", "", "Database name, eg: db1,db2,db3")
	analyzeCmd.Flags().StringVarP(&dbhost, "host", "H", "127.0.0.1", "Database host")
	analyzeCmd.Flags().StringVarP(&dbpassword, "password", "p", "123456", "Database passowrd")
	analyzeCmd.Flags().StringVarP(&dbtable, "tables", "t", "", "table names, eg: db1.table1,db1.table2,db2.table3")
	analyzeCmd.Flags().IntVarP(&dbport, "port", "P", 4000, "Database Port")
	analyzeCmd.Flags().IntVarP(&thread, "thread", "T", 4, "Number of goroutines to use")
	analyzeCmd.Flags().IntVarP(&mode, "mode", "m", 0, "Ignore system database, eg: 1")
}
