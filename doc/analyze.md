# 说明

```shell
./tidb-tools-ops analyze -h
Collect statistics. For example: analyze table test;

Usage:
  tidb-tools-ops analyze [flags]

Flags:
  -d, --database string   Database name, eg: db1,db2,db3
  -s, --healthy int       Table stats healthy, If it is below the threshold, then analyze (default 100)
  -h, --help              help for analyze
  -H, --host string       Database host (default "127.0.0.1")
  -m, --mode int          Ignore system database, eg: 1
  -p, --password string   Database passowrd (default "123456")
  -P, --port int          Database Port (default 4000)
  -t, --tables string     table names, eg: db1.table1,db1.table2,db2.table3
  -T, --thread int        Number of goroutines to use (default 4)
  -u, --user string       Username with privileges to run the analyze (default "root")
```

**参数说明：**
1. `-H, --dbhost`：数据库登录 IP
2. `-p, --dbpassword`：数据库登录密码
3. `-P, --dbport`：数据库登录端口
4. `-u, --dbusername`：数据库登录账户
5. `-d, --database`：需要收集统计信息的库，可以写多个
6. `-s, --healthy`：过滤条件：健康度，默认低于 100 的健康度的表才进行统计信息收集。
7. `-t, --tables`：需要收集的表，可以写多个，配置这个后 `-d` 配置就会失效
8. `-T, --thread`：收集统计信息的并行度，默认为 4

**使用案例**

```shell
./tidb-tools-ops analyze -d test -u root -P 4201 -p tidb@123 
Analyze All tables Succeeded
Total Cost time: 4.713658154s

# 查看日志
cat tools.log 
[INFO] 2022/10/11 09:46:00.135777 analyze.go:118: analyze table: test.sbtest3 Sucessfull 
[INFO] 2022/10/11 09:46:00.211260 analyze.go:118: analyze table: test.sbtest2 Sucessfull 
[INFO] 2022/10/11 09:46:00.211340 analyze.go:102: Analyze All tables Succeeded
[INFO] 2022/10/11 09:46:00.211372 analyze.go:105: Total Cost time: 4.713658154s
```

