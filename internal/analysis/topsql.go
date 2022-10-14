package analysis

import (
	"database/sql"
	"fmt"
	"tidb-tools-ops/pkg/logutil"
)

const topsql = `SELECT
s.Digest,
count(*) as Exec_count,
round(SUM(s.Query_time)/ count(*), 2) as Avg_Query_time,
round(SUM(s.Query_time), 0) as Sum_Query_time,
round(SUM(s.Cop_time), 0) as Sum_Cop_time,
round(SUM(s.Request_count), 0) as Sum_Request_count,
round(SUM(s.Process_keys), 0) as Sum_Process_keys,
round(SUM(s.Total_keys), 0) as Sum_Total_keys,
max(Query) as Query
FROM
INFORMATION_SCHEMA.CLUSTER_SLOW_QUERY s
WHERE
	Time >= ?
AND time <= ?
GROUP BY
Digest
ORDER BY
SUM(Request_count) DESC
LIMIT 10;`

type Topsql struct {
	Digest            string
	Exec_count        int
	Avg_Query_time    string
	Sum_Query_time    string
	Sum_Cop_time      string
	Sum_Request_count int
	Sum_Process_keys  int64
	Sum_Total_keys    int64
	Query             string
}

func AnalysisTopSql(db *sql.DB, Stime string, Etime string) {
	rows, err := db.Query(topsql, Stime, Etime)
	if err != nil {
		logutil.ErrorLog(err.Error())
	}
	var c Topsql
	for rows.Next() {
		err := rows.Scan(&c.Digest, &c.Exec_count, &c.Avg_Query_time, &c.Sum_Query_time, &c.Sum_Cop_time, &c.Sum_Request_count, &c.Sum_Process_keys, &c.Sum_Total_keys, &c.Query)

		if err != nil {
			logutil.ErrorLog(err.Error())
		}
		c.Out()
	}

}

func (c Topsql) Out() {
	fmt.Println(c)
}
