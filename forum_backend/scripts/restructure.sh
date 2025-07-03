#!/bin/bash

echo "=== QQ论坛项目重构脚本 ==="

# 创建新的目录结构
echo "1. 创建新的目录结构..."
mkdir -p backend frontend/web frontend/admin shared/types shared/constants shared/utils docs/api docs/design docs/deployment docs/development scripts infrastructure/docker infrastructure/k8s

# 移动后端代码
echo "2. 移动后端代码..."
if [ -d "api" ]; then
    mv api backend/
    echo "  ✓ 移动 api/ 到 backend/"
fi

if [ -d "model" ]; then
    mv model backend/
    echo "  ✓ 移动 model/ 到 backend/"
fi

if [ -d "cmd" ]; then
    mv cmd backend/
    echo "  ✓ 移动 cmd/ 到 backend/"
fi

if [ -d "database" ]; then
    mv database backend/
    echo "  ✓ 移动 database/ 到 backend/"
fi

# 移动Go模块文件
if [ -f "go.mod" ]; then
    mv go.mod backend/
    echo "  ✓ 移动 go.mod 到 backend/"
fi

if [ -f "go.sum" ]; then
    mv go.sum backend/
    echo "  ✓ 移动 go.sum 到 backend/"
fi

# 移动文档
if [ -d "doc" ]; then
    mv doc/* docs/
    rmdir doc
    echo "  ✓ 移动 doc/ 到 docs/"
fi

# 移动脚本
if [ -f "test.sh" ]; then
    mv test.sh scripts/
    echo "  ✓ 移动 test.sh 到 scripts/"
fi

if [ -f "test_register.sh" ]; then
    mv test_register.sh scripts/
    echo "  ✓ 移动 test_register.sh 到 scripts/"
fi

if [ -f "goctl_model_cmd.sh" ]; then
    mv goctl_model_cmd.sh scripts/
    echo "  ✓ 移动 goctl_model_cmd.sh 到 scripts/"
fi

# 创建配置文件
echo "3. 创建配置文件..."

# 创建根目录的docker-compose.yml
cat > docker-compose.yml << 'EOF'
version: '3.8'

services:
  # 数据库服务
  mysql:
    image: mysql:8.0
    container_name: qq_forum_mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root542
      MYSQL_DATABASE: qq_forum
      MYSQL_USER: qq_forum_user
      MYSQL_PASSWORD: qq_forum_pass
      MYSQL_ROOT_HOST: '%'
    ports:
      - "3307:3306"
    volumes:
      - ./backend/database/sql:/docker-entrypoint-initdb.d
      - mysql_data:/var/lib/mysql
    networks:
      - qq_forum_network

  # Redis服务
  redis:
    image: redis:7.0
    container_name: qq_forum_redis
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    environment:
      - TZ=Asia/Shanghai
    command: redis-server --requirepass redisqiuqiu542
    networks:
      - qq_forum_network

  # 后端API服务
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: qq_forum_backend
    restart: always
    ports:
      - "8888:8888"
    depends_on:
      - mysql
      - redis
    environment:
      - ENV=production
    networks:
      - qq_forum_network

  # 前端Web服务
  frontend:
    build:
      context: ./frontend/web
      dockerfile: Dockerfile
    container_name: qq_forum_frontend
    restart: always
    ports:
      - "3000:3000"
    depends_on:
      - backend
    networks:
      - qq_forum_network

volumes:
  mysql_data:
  redis_data:

networks:
  qq_forum_network:
    driver: bridge
EOF

# 创建Makefile
cat > Makefile << 'EOF'
.PHONY: help dev dev-backend dev-frontend build build-backend build-frontend test clean

help: ## 显示帮助信息
	@echo "QQ论坛项目构建脚本"
	@echo "可用命令："
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## 启动所有开发服务
	@echo "启动开发环境..."
	docker-compose -f docker-compose.dev.yml up -d

dev-backend: ## 只启动后端开发服务
	@echo "启动后端开发服务..."
	cd backend && go run main.go

dev-frontend: ## 只启动前端开发服务
	@echo "启动前端开发服务..."
	cd frontend/web && npm run dev

build: build-backend build-frontend ## 构建所有项目

build-backend: ## 构建后端
	@echo "构建后端..."
	cd backend && go build -o bin/server .

build-frontend: ## 构建前端
	@echo "构建前端..."
	cd frontend/web && npm run build

test: ## 运行测试
	@echo "运行测试..."
	cd backend && go test ./...
	cd frontend/web && npm test

clean: ## 清理构建文件
	@echo "清理构建文件..."
	rm -rf backend/bin
	rm -rf frontend/web/dist
	docker-compose down -v
EOF

# 创建开发环境的docker-compose
cat > docker-compose.dev.yml << 'EOF'
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    container_name: qq_forum_mysql_dev
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root542
      MYSQL_DATABASE: qq_forum
      MYSQL_USER: qq_forum_user
      MYSQL_PASSWORD: qq_forum_pass
      MYSQL_ROOT_HOST: '%'
    ports:
      - "3307:3306"
    volumes:
      - ./backend/database/sql:/docker-entrypoint-initdb.d
      - mysql_dev_data:/var/lib/mysql
    networks:
      - qq_forum_dev_network

  redis:
    image: redis:7.0
    container_name: qq_forum_redis_dev
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - redis_dev_data:/data
    environment:
      - TZ=Asia/Shanghai
    command: redis-server --requirepass redisqiuqiu542
    networks:
      - qq_forum_dev_network

volumes:
  mysql_dev_data:
  redis_dev_data:

networks:
  qq_forum_dev_network:
    driver: bridge
EOF

# 更新.gitignore
cat >> .gitignore << 'EOF'

# 前端相关
frontend/web/node_modules/
frontend/web/dist/
frontend/web/.env.local
frontend/web/.env.development.local
frontend/web/.env.test.local
frontend/web/.env.production.local
frontend/web/npm-debug.log*
frontend/web/yarn-debug.log*
frontend/web/yarn-error.log*

# 后端相关
backend/bin/
backend/tmp/
*.exe
*.exe~
*.dll
*.so
*.dylib

# 数据库
*.db
*.sqlite

# 日志文件
*.log

# 环境变量
.env
.env.local
.env.development.local
.env.test.local
.env.production.local

# IDE
.vscode/
.idea/
*.swp
*.swo

# 系统文件
.DS_Store
Thumbs.db
EOF

echo "4. 创建前端项目..."
cd frontend/web

# 检查是否已存在package.json
if [ ! -f "package.json" ]; then
    echo "  创建React + TypeScript项目..."
    npm create vite@latest . -- --template react-ts --yes
    npm install
    echo "  ✓ 前端项目创建完成"
else
    echo "  ✓ 前端项目已存在"
fi

cd ../..

echo "5. 创建共享类型定义..."
cat > shared/types/api.ts << 'EOF'
// API响应类型定义
export interface BaseResponse {
  code: number;
  message: string;
}

export interface User {
  id: number;
  email: string;
  username: string;
  avatar?: string;
  signature?: string;
  role: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
}

export interface RegisterResponse extends BaseResponse {
  user_id: number;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse extends BaseResponse {
  token: string;
  user: User;
}
EOF

echo "6. 创建API服务..."
cat > frontend/web/src/services/api.ts << 'EOF'
import axios from 'axios';
import type { RegisterRequest, RegisterResponse, LoginRequest, LoginResponse } from '../../../shared/types/api';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8888';

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    console.error('API Error:', error);
    return Promise.reject(error);
  }
);

// 用户相关API
export const userApi = {
  register: (data: RegisterRequest): Promise<RegisterResponse> => {
    return api.post('/api/user/register', data);
  },
  
  login: (data: LoginRequest): Promise<LoginResponse> => {
    return api.post('/api/user/login', data);
  },
  
  getProfile: (): Promise<any> => {
    return api.get('/api/user/profile');
  },
};

export default api;
EOF

echo "7. 创建环境变量文件..."
cat > frontend/web/.env.development << 'EOF'
VITE_API_BASE_URL=http://localhost:8888
VITE_APP_TITLE=QQ论坛 (开发环境)
EOF

cat > frontend/web/.env.production << 'EOF'
VITE_API_BASE_URL=https://api.qqforum.com
VITE_APP_TITLE=QQ论坛
EOF

echo "=== 项目重构完成！ ==="
echo ""
echo "下一步操作："
echo "1. 安装前端依赖: cd frontend/web && npm install"
echo "2. 启动开发环境: make dev"
echo "3. 启动后端服务: make dev-backend"
echo "4. 启动前端服务: make dev-frontend"
echo ""
echo "项目结构已重新组织完成！" 