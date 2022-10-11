# tidb-tools-ops

**说明**：该项目为日常运维 TIDB 创建。

## 安装

**方法一**

```shell
# 1. 下载项目
git clone https://github.com/WalterWj/tidb-tools-ops.git

# 2. 打包
make

# 随后可以看到在当前目录下，生成 bin 目录下有 tidb-tools-ops 包，直接使用即可。
```

**方法二**

在 [release](https://github.com/WalterWj/tidb-tools-ops/releases) 页面下载 

# 功能说明

## 功能一

导出当前 TIDB 所有账户的密码，权限，生成 users.sql 文件。可以 [tidb-tools-ops export](./doc/export_users.md) 导出。

## 功能二

按照一定条件对 TiDB 集群的表进行并行的统计信息收集。可以使用 [tidb-tools-ops analyze](./doc/analyze.md) 命令。
