#!/bin/bash

# 简单的 Post API 测试脚本（开发环境使用）
# 使用方法: ./scripts/post_test_simple.sh

BASE_URL="http://localhost:8888/api/v1"
API_PREFIX="$BASE_URL/p"

# 测试用户ID（可以根据需要修改）
TEST_USER_ID="1"

# 颜色定义
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

# 带测试用户ID的 curl 函数
test_curl() {
    local method=$1
    local url=$2
    local data=$3
    
    if [ -n "$data" ]; then
        curl -s -X "$method" "$url" \
            -H "Content-Type: application/json" \
            -H "X-Test-User-Id: $TEST_USER_ID" \
            -d "$data"
    else
        curl -s -X "$method" "$url" \
            -H "X-Test-User-Id: $TEST_USER_ID"
    fi
}

# 测试创建帖子
test_create_post() {
    print_info "=== 测试创建帖子 ==="
    
    print_info "创建技术讨论帖子..."
    response=$(test_curl "POST" "$API_PREFIX/create" '{
        "title": "Go语言并发编程实践",
        "content": "本文介绍Go语言中的goroutine和channel的使用方法...",
        "category_id": 1
    }')
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    print_info "创建生活随笔帖子..."
    response=$(test_curl "POST" "$API_PREFIX/create" '{
        "title": "生活随笔",
        "content": "今天天气很好，心情也不错..."
    }')
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

# 测试获取帖子列表
test_list_posts() {
    print_info "=== 测试获取帖子列表 ==="
    
    print_info "获取所有帖子..."
    response=$(test_curl "GET" "$API_PREFIX/list?category_id=1&last_index=0&page_size=10&order_by=created_time&order_type=desc")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

# 测试获取帖子详情
test_get_post() {
    print_info "=== 测试获取帖子详情 ==="
    
    print_info "获取帖子ID=1的详情..."
    response=$(test_curl "GET" "$API_PREFIX/detail?id=1")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

# 测试更新帖子
test_update_post() {
    print_info "=== 测试更新帖子 ==="
    
    print_info "更新帖子标题..."
    response=$(test_curl "PUT" "$API_PREFIX/update" '{
        "id": 1,
        "title": "更新后的Go语言并发编程实践"
    }')
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

# 主函数
main() {
    print_info "简单 Post API 测试脚本"
    print_info "测试用户ID: $TEST_USER_ID"
    print_info "基础URL: $BASE_URL"
    echo ""
    
    test_create_post
    echo ""
    
    test_list_posts
    echo ""
    
    test_get_post
    echo ""
    
    test_update_post
    echo ""
    
    print_success "测试完成！"
}

# 运行主函数
main "$@" 