package analysis_test

import (
	"fmt"
	"strings"
	"testing"
	"tidb-tools-ops/internal/analysis"
	"tidb-tools-ops/pkg/dbutil"
	"tidb-tools-ops/pkg/logutil"
)

func TestTopsql(t *testing.T) {
	logutil.InitLog("test.log")
	// nomal
	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	db, err := dbutil.MysqlConnect(dsn)
	if err != nil {
		fmt.Println(err.Error())
	}

	ct := make(chan analysis.Topsql)

	go analysis.AnalysisTopSql(db, "2021/10/10 10:00:00", "2022/10/10 12:00:00", ct)
	analysis.OutHtml(ct)

	// ct.Out()

}
