package analysis

const topsql = `SELECT
s.Digest,
count(*) as exec_count,
round(SUM(s.Query_time)/ count(*), 2) as avg_Query_time,
round(SUM(s.Query_time), 0) as sum_Query_time,
round(SUM(s.Cop_time), 0) as sum_Cop_time,
round(SUM(s.Request_count), 0) as sum_Request_count,
round(SUM(s.Process_keys), 0) as sum_Process_keys,
round(SUM(s.Total_keys), 0) as sum_Total_keys,
max(Query) as Query
FROM
INFORMATION_SCHEMA.CLUSTER_SLOW_QUERY s
WHERE
	Time >= '?'
AND time <= '?'
GROUP BY
Digest
ORDER BY
SUM(Request_count) DESC
LIMIT 10;`
