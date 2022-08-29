---
title: Go Cobra
time: 2020年06月16日
auto: 王军
---
# 使用

命令行项目框架生成器。

基本模型

```shell
APPNAME COMMAND ARG --FLAG
```

## 安装

- 下载

```shell
go get -u github.com/spf13/cobra/cobra
```

- 包

```shell
cd $GOPATH/bin
# 可以看到 cobra 包
```

- 初始化

```shell
$GOPATH/bin/cobra init --pkg-name github.com/WalterWj/go-study

# 这样在我的 go-study 项目中，创建了 cmd 目录(初始化会有一个 root.go)和 main.go,LICENSE 两个文件。
```

- 创建 test.go (CLIs程序)

```shell
$GOPATH/bin/cobra add test

# 在 cmd 目录下就有了 test.go
# build 之后，可以 test -h 调用
```

- 创建附加命令，如 version.go (版本)

```shell
# 在 cmd 目录下创建 version.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-study",
	Long:  `All software has versions. This is go-study's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-study Version: v1.0")
	},
}

# 编译完成后使用即可。
```