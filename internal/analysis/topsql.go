package analysis

import (
	"database/sql"
	"fmt"
	"html/template"
	"os"
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
LIMIT 2;`

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

func AnalysisTopSql(db *sql.DB, Stime string, Etime string, ct chan Topsql) {
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

		ct <- c
		// c.Out()
		// c.OutHtml()
	}

	close(ct)
}

func Out(c chan Topsql) {
	for {
		ct, ok := <-c
		if !ok {
			return
		} else {
			fmt.Printf("指纹码: %v\n执行次数: %v\n", ct.Digest, ct.Exec_count)
		}
	}
}

func OutHtml(c chan Topsql) {
	const TableHtml = `
<tr>
<td>{{ .Digest }}</td>
<td>{{ .Exec_count }} </td>
<td>{{ .Avg_Query_time }}</td>
<td>{{ .Sum_Query_time }}</td>
<td>{{ .Sum_Cop_time }}</td>
<td>{{ .Sum_Request_count }}</td>
<td>{{ .Sum_Process_keys }}</td>
<td>{{ .Sum_Total_keys }}</td>
<td>{{ .Query }}</td>
</tr>
`

	for {
		ct, ok := <-c
		if !ok {
			return
		} else {
			t := template.Must(template.New("tables").Parse(TableHtml))
			t.Execute(os.Stdout, ct)
		}
	}

}
