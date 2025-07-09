#1. ------------------------------ 获取分类 ------------------------------
# 获取所有分类（使用默认分页）
curl -X GET "http://localhost:8888/api/v1/ca/list"
# 指定分页参数
curl -X GET "http://localhost:8888/api/v1/ca/list?page=1&page_size=3"
# 只指定页码
curl -X GET "http://localhost:8888/api/v1/ca/list?page=1"
# 获取 ID 为 1 的分类详情
curl -X GET "http://localhost:8888/api/v1/ca/detail?id=1"

#2. ------------------------------ 创建分类 ------------------------------
# 创建技术分类
curl -X POST "http://localhost:8888/api/v1/ca/create" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "技术讨论",
    "description": "技术相关话题讨论",
    "sort_order": 1
  }'
# 创建生活分类
curl -X POST "http://localhost:8888/api/v1/ca/create" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "生活分享",
    "description": "日常生活分享",
    "sort_order": 2
  }'
# 创建游戏分类（不包含描述）
curl -X POST "http://localhost:8888/api/v1/ca/create" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "游戏讨论",
    "sort_order": 3
  }'

#3. ------------------------------ 更新分类 ------------------------------
# 更新分类信息
curl -X PUT "http://localhost:8888/api/v1/ca/update" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 2,
    "name": "技术交流",
    "description": "技术交流与分享",
    "sort_order": 1,
    "is_active": 1
  }'

# 只更新部分字段
curl -X PUT "http://localhost:8888/api/v1/ca/update" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 2,
    "name": "技术交流升级版"
  }'

#4. ------------------------------ 删除分类 ------------------------------
# 删除分类（软删除）
curl -X DELETE "http://localhost:8888/api/v1/ca/delete" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 3
  }'