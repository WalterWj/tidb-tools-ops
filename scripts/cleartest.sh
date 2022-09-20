#!/bin/bash

dir=$1
cd "${dir}" || exit 1

# 删除生成的 html 和 out 文件
rm hint_test.html hint_test.out test.log
