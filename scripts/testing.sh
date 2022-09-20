#!/bin/bash

dir=$1
cd "${dir}" || exit 1

# 生产 out 文件，run testing
go test -v -covermode=set -coverprofile=hint_test.out ./

# 查看测试覆盖率
go tool cover -func=hint_test.out

# 生产测试覆盖率 HTML
go tool cover -html=hint_test.out -o hint_test.html
