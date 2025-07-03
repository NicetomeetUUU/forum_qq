# QQ论坛项目 - 新结构说明

## 🚀 快速开始

### 1. 执行重构脚本
```bash
chmod +x scripts/restructure.sh
./scripts/restructure.sh
```

### 2. 安装依赖
```bash
# 安装前端依赖
cd frontend/web
npm install

# 安装后端依赖
cd backend
go mod tidy
```

### 3. 启动开发环境
```bash
# 启动数据库和Redis
make dev

# 启动后端服务
make dev-backend

# 启动前端服务
make dev-frontend
```

## 📁 项目结构

```
qq_forum/
├── backend/                    # 后端代码
│   ├── api/                   # API服务
│   ├── model/                 # 数据模型
│   ├── cmd/                   # 命令行工具
│   ├── database/              # 数据库相关
│   └── go.mod
│
├── frontend/                   # 前端代码
│   └── web/                   # Web应用
│       ├── src/
│       ├── public/
│       └── package.json
│
├── shared/                     # 共享代码
│   └── types/                 # 类型定义
│
├── docs/                       # 文档
├── scripts/                    # 脚本
└── infrastructure/            # 基础设施
```

## 🛠️ 开发命令

### 后端开发
```bash
# 启动后端服务
make dev-backend

# 构建后端
make build-backend

# 运行测试
cd backend && go test ./...
```

### 前端开发
```bash
# 启动前端开发服务器
make dev-frontend

# 构建前端
make build-frontend

# 运行测试
cd frontend/web && npm test
```

### 数据库操作
```bash
# 启动数据库服务
docker-compose -f docker-compose.dev.yml up -d

# 进入数据库
docker exec -it qq_forum_mysql_dev mysql -u root -proot542 qq_forum

# 查看Redis
docker exec -it qq_forum_redis_dev redis-cli -a redisqiuqiu542
```

## 🔧 配置说明

### 环境变量
- `VITE_API_BASE_URL`: 前端API基础URL
- `MYSQL_ROOT_PASSWORD`: MySQL root密码
- `REDIS_PASSWORD`: Redis密码

### 端口配置
- 前端: http://localhost:3000
- 后端API: http://localhost:8888
- MySQL: localhost:3307
- Redis: localhost:6379

## 📝 API测试

### 用户注册
```bash
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123"
  }'
```

### 使用测试脚本
```bash
./scripts/test_register.sh
```

## 🚀 部署

### 开发环境
```bash
make dev
```

### 生产环境
```bash
# 构建所有项目
make build

# 启动生产服务
docker-compose up -d
```

## 📚 文档

- [API文档](./docs/api/)
- [设计文档](./docs/design/)
- [部署文档](./docs/deployment/)
- [开发文档](./docs/development/)

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## �� 许可证

MIT License 