# QQ论坛项目

基于go-zero框架开发的论坛系统，采用自底向上的开发方式。

## 项目结构

```
qq_forum/
├── cmd/                    # 命令行工具和测试程序
│   └── test_func/         # 模型层功能测试
├── model/                 # 数据模型层
│   ├── user/             # 用户模型
│   ├── post/             # 帖子模型
│   └── comment/          # 评论模型
├── database/             # 数据库相关
│   ├── docker-compose.yml # Docker编排文件
│   └── sql/              # SQL脚本
└── test.sh               # 测试脚本
```

## 快速开始

### 1. 启动数据库

```bash
cd database
docker-compose up -d
```

### 2. 运行测试

```bash
# 运行所有测试
./test.sh

# 或者单独运行模型层测试
cd cmd/test_func && go run main.go    # 模型层功能测试
go test ./model/user -v               # 单元测试
```

## 测试方法

### 1. 单元测试（推荐）

使用Go标准测试框架进行单元测试：

```bash
# 运行用户模型测试
go test ./model/user -v

# 运行所有测试
go test ./... -v
```

### 2. 功能测试

直接运行测试程序验证模型层功能：

```bash
# 测试模型层基本功能
cd cmd/test_func
go run main.go
```

### 3. 自动化测试脚本

使用提供的测试脚本一键运行所有测试：

```bash
./test.sh
```

## 开发建议

### 自底向上开发流程

1. **数据模型层** - 先完成数据库表设计和模型生成
2. **业务逻辑层** - 实现服务层的业务逻辑（待开发）
3. **API接口层** - 开发HTTP API接口（待开发）
4. **前端界面** - 最后开发用户界面（待开发）

### 测试策略

- **模型层**：使用单元测试验证CRUD操作
- **服务层**：待开发时添加服务层测试
- **API层**：待开发时添加API层测试

### 快速验证

在开发过程中，可以快速验证模型层功能：

```bash
# 验证数据库连接
mysql -h localhost -P 3307 -u root -proot542 -e "USE qq_forum; SELECT 1;"

# 验证模型功能
cd cmd/test_func && go run main.go
```

## 环境要求

- Go 1.24+
- Docker & Docker Compose
- MySQL 8.0
- Redis 7.0

## 配置说明

数据库连接配置在测试文件中：
- 主机：localhost
- 端口：3307
- 用户名：root
- 密码：root542
- 数据库：qq_forum

## 注意事项

1. 确保Docker容器正在运行
2. 测试前确保数据库已初始化
3. 测试数据会在数据库中创建，注意数据清理
4. 生产环境中请修改数据库连接配置

## 基于go-zero的论坛项目
（从基本的数据库crud不断扩展其他中间件了解整个框架如何使用）