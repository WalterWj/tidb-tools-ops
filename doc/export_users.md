# 说明

```shell
./tidb-tools-ops export -h
export your users for TiDB

Usage:
  tidb-tools-ops export [flags]

Flags:
  -H, --dbhost string       Database host (default "127.0.0.1")
  -p, --dbpassword string   Database passowrd (default "123456")
  -P, --dbport string       Database Port (default "4000")
  -u, --dbusername string   Database user (default "root")
  -h, --help                help for export
```

**参数说明：**
1. `-H, --dbhost`：数据库登录 IP
2. `-p, --dbpassword`：数据库登录密码
3. `-P, --dbport`：数据库登录端口
4. `-u, --dbusername`：数据库登录账户

**使用案例**

```shell
./tidb-tools-ops export -H 127.0.0.1 -p 123456 -P 4000 -u root
Successfully introduce all users and permissions.

cat users.sql
create user 'root'@'%';
update mysql.user set `authentication_string`='*6BB4837EB74329105EE4568DDA7DC67ED2CA2AD9' where user='root' and host='%';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;

create user 'readonly'@'%';
update mysql.user set `authentication_string`='*23AE809DDACAF96AF0FD78ED04B6A265E05AA257' where user='readonly' and host='%';
GRANT USAGE ON *.* TO 'readonly'@'%';
GRANT Select ON test.* TO 'readonly'@'%';
```

根据导出的内容，可以手动删除不需要迁移的用户代码。