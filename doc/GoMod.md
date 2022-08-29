---
title: Go modules
time: 2020年04月21日
auto: 王军
---
# 使用

## 创建 gomod

```shell
go mod init github.com/WalterWj/go-study
```

* `go run`, `go build`, `go test` 时，会自动下载相关依赖包。

## 好处

* gomod 可以语义话版本号，可以 git 对应分支或者 tag 的依赖包

```shell
go get github.com/go-sql-driver/mysql@v1.5.0
```

## 常用命令

* 查看所以依赖版本

```shell
$ go list -u -m all
github.com/WalterWj/go-study
github.com/go-sql-driver/mysql v1.5.0
```

* 升级

```shell
go get -u github.com/go-sql-driver/mysql  
```

* 只升级补丁

```shell
go get -u=patch github.com/go-sql-driver/mysql  
```

* 升降级版本

```shell
运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
运行 go get -u=patch 将会升级到最新的修订版本
运行 go get package@version 将会升级到指定的版本号version
```