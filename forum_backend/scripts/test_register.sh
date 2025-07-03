#!/bin/bash

echo "=== 用户注册接口测试 ==="

# 测试1：正常注册
echo "1. 测试正常注册"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user1@example.com",
    "username": "user1",
    "password": "password123"
  }'
echo -e "\n"

# 测试2：重复邮箱
echo "2. 测试重复邮箱"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user1@example.com",
    "username": "user2",
    "password": "password123"
  }'
echo -e "\n"

# 测试3：重复用户名
echo "3. 测试重复用户名"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user2@example.com",
    "username": "user1",
    "password": "password123"
  }'
echo -e "\n"

# 测试4：空邮箱
echo "4. 测试空邮箱"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "",
    "username": "user3",
    "password": "password123"
  }'
echo -e "\n"

# 测试5：空用户名
echo "5. 测试空用户名"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user3@example.com",
    "username": "",
    "password": "password123"
  }'
echo -e "\n"

# 测试6：空密码
echo "6. 测试空密码"
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user4@example.com",
    "username": "user4",
    "password": ""
  }'
echo -e "\n"

echo "=== 测试完成 ===" 