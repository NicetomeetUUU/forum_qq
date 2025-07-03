#!/bin/bash

echo "=== 更新Go文件中的import路径 ==="

# 进入后端目录
cd "$(dirname "$0")/.."

# 更新所有Go文件中的import路径
echo "1. 更新import路径从 qq_forum 到 forum_backend..."

# 使用find和sed批量替换
find . -name "*.go" -type f -exec sed -i '' 's|qq_forum/|forum_backend/|g' {} \;

echo "2. 更新数据库连接字符串..."
# 更新数据库连接字符串中的数据库名
find . -name "*.go" -type f -exec sed -i '' 's|/qq_forum?|/forum_backend?|g' {} \;

echo "3. 运行go mod tidy..."
go mod tidy

echo "4. 验证import路径..."
# 检查是否还有旧的import路径
if grep -r "qq_forum/" . --include="*.go"; then
    echo "⚠️  发现残留的旧import路径，请手动检查："
    grep -r "qq_forum/" . --include="*.go"
else
    echo "✅ 所有import路径已更新完成"
fi

echo "=== 更新完成 ===" 