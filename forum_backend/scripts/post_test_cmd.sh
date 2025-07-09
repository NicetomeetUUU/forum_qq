#!/bin/bash

# Post API 测试脚本
# 使用方法: ./scripts/post_test_cmd.sh

BASE_URL="http://localhost:8888/api/v1"
API_PREFIX="$BASE_URL/p"

# 测试用户ID（可以根据需要修改）
TEST_USER_ID="1"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 测试函数
test_create_post() {
    print_info "=== 测试创建帖子 ==="
    
    # 测试1: 创建带分类的帖子
    print_info "创建技术讨论帖子..."
    request_data="{
        \"user_id\": $TEST_USER_ID,
        \"title\": \"Go语言并发编程实践\",
        \"content\": \"本文介绍Go语言中的goroutine和channel的使用方法...\",
        \"category_id\": 1
    }"
    print_info "请求数据: $request_data"
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d "$request_data")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 创建无分类的帖子
    print_info "创建无分类帖子..."
    request_data="{
        \"user_id\": $TEST_USER_ID,
        \"title\": \"生活随笔\",
        \"content\": \"今天天气很好，心情也不错...\"
    }"
    print_info "请求数据: $request_data"
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d "$request_data")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试3: 创建帖子（缺少必填字段）
    print_info "测试缺少标题字段..."
    request_data="{
        \"user_id\": $TEST_USER_ID,
        \"content\": \"只有内容，没有标题\"
    }"
    print_info "请求数据: $request_data"
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d "$request_data")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试4: 创建帖子（不存在的分类）
    print_info "测试不存在的分类..."
    request_data="{
        \"user_id\": $TEST_USER_ID,
        \"title\": \"测试帖子\",
        \"content\": \"测试内容\",
        \"category_id\": 999
    }"
    print_info "请求数据: $request_data"
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d "$request_data")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

test_get_post() {
    print_info "=== 测试获取帖子详情 ==="
    
    # 测试1: 获取存在的帖子
    print_info "获取帖子ID=1的详情..."
    request_url="$API_PREFIX/detail?id=1"
    print_info "请求URL: $request_url"
    response=$(curl -s -X GET "$request_url")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 获取不存在的帖子
    print_info "获取不存在的帖子ID=999..."
    request_url="$API_PREFIX/detail?id=999"
    print_info "请求URL: $request_url"
    response=$(curl -s -X GET "$request_url")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试3: 无效的ID参数
    print_info "测试无效ID参数..."
    request_url="$API_PREFIX/detail?id=abc"
    print_info "请求URL: $request_url"
    response=$(curl -s -X GET "$request_url")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

test_list_posts() {
    print_info "=== 测试获取帖子列表 ==="
    
    # 测试1: 获取默认分类的帖子
    print_info "获取默认分类的帖子列表..."
    request_url="$API_PREFIX/list?category_id=1&last_index=0&page_size=10&order_by=created_time&order_type=desc"
    print_info "请求URL: $request_url"
    response=$(curl -s -X GET "$request_url")
    print_info "响应结果:"
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 获取所有帖子（不指定分类）
    print_info "获取所有帖子列表..."
    response=$(curl -s -X GET "$API_PREFIX/list?category_id=0&last_index=0&page_size=5&order_by=created_time&order_type=desc")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试3: 按更新时间排序
    print_info "按更新时间排序..."
    response=$(curl -s -X GET "$API_PREFIX/list?category_id=1&last_index=0&page_size=10&order_by=updated_time&order_type=desc")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试4: 分页测试
    print_info "分页测试（第二页）..."
    response=$(curl -s -X GET "$API_PREFIX/list?category_id=1&last_index=10&page_size=5&order_by=created_time&order_type=desc")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试5: 无效参数
    print_info "测试无效参数..."
    response=$(curl -s -X GET "$API_PREFIX/list?category_id=1&last_index=0&page_size=0&order_by=invalid&order_type=invalid")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

test_update_post() {
    print_info "=== 测试更新帖子 ==="
    
    # 测试1: 更新帖子标题
    print_info "更新帖子标题..."
    response=$(curl -s -X PUT "$API_PREFIX/update" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": 1,
            \"user_id\": $TEST_USER_ID,
            \"title\": \"更新后的Go语言并发编程实践\"
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 更新帖子内容
    print_info "更新帖子内容..."
    response=$(curl -s -X PUT "$API_PREFIX/update" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": 1,
            \"user_id\": $TEST_USER_ID,
            \"content\": \"更新后的内容：Go语言并发编程是Go语言最重要的特性之一...\"
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试3: 更新分类
    print_info "更新帖子分类..."
    response=$(curl -s -X PUT "$API_PREFIX/update" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": 1,
            \"user_id\": $TEST_USER_ID,
            \"category_id\": 1
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试4: 更新状态
    print_info "更新帖子状态..."
    response=$(curl -s -X PUT "$API_PREFIX/update" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": 1,
            \"user_id\": $TEST_USER_ID,
            \"status\": 1
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试5: 更新不存在的帖子
    print_info "更新不存在的帖子..."
    response=$(curl -s -X PUT "$API_PREFIX/update" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": 999,
            \"user_id\": $TEST_USER_ID,
            \"title\": \"不存在的帖子\"
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

test_delete_post() {
    print_info "=== 测试删除帖子 ==="
    
    # 注意：删除操作需要谨慎，这里只是演示API调用
    print_warning "删除操作会永久删除数据，请谨慎使用！"
    
    # 测试1: 删除存在的帖子
    print_info "删除帖子ID=2..."
    response=$(curl -s -X DELETE "$API_PREFIX/delete/2" \
        -H "Content-Type: application/json" \
        -d "{
            \"user_id\": $TEST_USER_ID
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 删除不存在的帖子
    print_info "删除不存在的帖子ID=999..."
    response=$(curl -s -X DELETE "$API_PREFIX/delete/999" \
        -H "Content-Type: application/json" \
        -d "{
            \"user_id\": $TEST_USER_ID
        }")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

test_error_cases() {
    print_info "=== 测试错误情况 ==="
    
    # 测试1: 无效的JSON
    print_info "测试无效JSON..."
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d '{invalid json}')
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试2: 缺少Content-Type
    print_info "测试缺少Content-Type..."
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -d "{\"user_id\": $TEST_USER_ID, \"title\": \"test\", \"content\": \"test\"}")
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
    
    # 测试3: 空请求体
    print_info "测试空请求体..."
    response=$(curl -s -X POST "$API_PREFIX/create" \
        -H "Content-Type: application/json" \
        -d '{}')
    echo "$response" | jq '.' 2>/dev/null || echo "$response"
}

# 完整测试流程
run_full_test() {
    print_info "开始完整的Post API测试..."
    print_info "测试用户ID: $TEST_USER_ID"
    echo ""
    
    test_create_post
    echo ""
    
    test_get_post
    echo ""
    
    test_list_posts
    echo ""
    
    test_update_post
    echo ""
    
    test_error_cases
    echo ""
    
    print_success "Post API 测试完成！"
}

# 交互式测试
interactive_test() {
    while true; do
        echo ""
        print_info "请选择要测试的功能："
        echo "1. 创建帖子"
        echo "2. 获取帖子详情"
        echo "3. 获取帖子列表"
        echo "4. 更新帖子"
        echo "5. 删除帖子"
        echo "6. 错误情况测试"
        echo "7. 运行完整测试"
        echo "0. 退出"
        echo ""
        read -p "请输入选项 (0-7): " choice
        
        case $choice in
            1) test_create_post ;;
            2) test_get_post ;;
            3) test_list_posts ;;
            4) test_update_post ;;
            5) test_delete_post ;;
            6) test_error_cases ;;
            7) run_full_test ;;
            0) print_info "退出测试"; exit 0 ;;
            *) print_error "无效选项，请重新选择" ;;
        esac
    done
}

# 主函数
main() {
    print_info "Post API 测试脚本"
    print_info "基础URL: $BASE_URL"
    print_info "测试用户ID: $TEST_USER_ID"
    echo ""
    
    # 检查jq是否安装
    if ! command -v jq &> /dev/null; then
        print_warning "jq 未安装，JSON输出将不会被格式化"
        print_info "安装jq: brew install jq (macOS) 或 apt-get install jq (Ubuntu)"
    fi
    
    # 检查参数
    if [ "$1" = "interactive" ]; then
        interactive_test
    else
        run_full_test
    fi
}

# 运行主函数
main "$@"