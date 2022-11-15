package splittable

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"tidb-tools-ops/pkg/logutil"
)

type RegionStatus struct {
	index_id   sql.NullInt64
	index_name sql.NullString
	start_ts   string
	end_ts     string
}

func Keys(db *sql.DB, dbname string, tbname string, ct chan RegionStatus) {
	const keys_sql = `select
	INDEX_ID,
	INDEX_NAME,
	tidb_decode_key(start_key) start,
	tidb_decode_key(END_KEY) end
from
	information_schema.tikv_region_status
where
	DB_NAME = ?
	and TABLE_NAME = ?
order by
	INDEX_ID;`
	rows, err := db.Query(keys_sql, dbname, tbname)
	if err != nil {
		logutil.ErrorLog(err.Error())
	}

	var rs RegionStatus
	for rows.Next() {
		err := rows.Scan(&rs.index_id, &rs.index_name, &rs.start_ts, &rs.end_ts)
		if err != nil {
			logutil.ErrorLog(err.Error())
		}
		ct <- rs
	}

	defer close(ct)

}

func Out(c chan RegionStatus) {
	// var rt map[string]string
	for {
		ct, ok := <-c
		if !ok {
			return
		} else {
			// ts to map
			st := JsonToMap(ct.start_ts)
			et := JsonToMap(ct.end_ts)
			fmt.Printf("index id: %v\n index name: %v\n start ts: %v\n end ts: %v \n", ct.index_id, ct.index_name, st, et)
			fmt.Print("====", ParserKeys(st), "\n")
			// ParserKeys(et)
		}
	}
}

func ParserKeys(keys map[string]interface{}) string {
	var rt string
	if _, ok := keys["index_vals"]; ok {
		val := keys["index_vals"]
		rt = "("
		for _, ct := range val.(map[string]interface{}) {
			rt = rt + fmt.Sprint(ct) + ","
		}
		rt = strings.Trim(rt, ",")
		rt += ")"
		return rt
	}

	return rt
}

func JsonToMap(str string) map[string]interface{} {
	// ts to map
	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		logutil.ErrorLog(err.Error())
	}
	// decode result;Formatting prevents scientific counting
	data, _ := json.Marshal(tempMap)
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	_ = d.Decode(&tempMap)

	return tempMap
}
