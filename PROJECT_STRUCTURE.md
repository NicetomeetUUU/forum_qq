# QQ论坛项目结构规划

## 推荐结构：Monorepo（单仓库）

```
qq_forum/
├── README.md                    # 项目总览
├── .gitignore                   # Git忽略文件
├── docker-compose.yml           # 整体服务编排
├── Makefile                     # 构建脚本
│
├── backend/                     # 后端代码
│   ├── api/                     # API服务
│   │   ├── user/                # 用户服务
│   │   │   └── userService/
│   │   ├── post/                # 帖子服务
│   │   │   └── postService/
│   │   └── comment/             # 评论服务
│   │       └── commentService/
│   ├── model/                   # 数据模型
│   │   ├── user/
│   │   ├── post/
│   │   └── comment/
│   ├── cmd/                     # 命令行工具
│   ├── database/                # 数据库相关
│   │   ├── docker-compose.yml   # 数据库服务
│   │   ├── sql/
│   │   └── migrations/
│   ├── configs/                 # 配置文件
│   │   ├── dev/
│   │   ├── test/
│   │   └── prod/
│   ├── scripts/                 # 构建脚本
│   ├── docs/                    # API文档
│   ├── tests/                   # 测试文件
│   ├── go.mod
│   └── go.sum
│
├── frontend/                    # 前端代码
│   ├── web/                     # Web应用
│   │   ├── public/
│   │   ├── src/
│   │   │   ├── components/      # React组件
│   │   │   ├── pages/          # 页面组件
│   │   │   ├── services/       # API服务
│   │   │   ├── utils/          # 工具函数
│   │   │   ├── styles/         # 样式文件
│   │   │   └── types/          # TypeScript类型
│   │   ├── package.json
│   │   └── vite.config.ts
│   │
│   ├── mobile/                  # 移动端应用（可选）
│   │   ├── android/
│   │   └── ios/
│   │
│   └── admin/                   # 管理后台（可选）
│       ├── src/
│       ├── package.json
│       └── vite.config.ts
│
├── shared/                      # 共享代码
│   ├── types/                   # 共享类型定义
│   ├── constants/               # 共享常量
│   └── utils/                   # 共享工具函数
│
├── docs/                        # 项目文档
│   ├── api/                     # API文档
│   ├── design/                  # 设计文档
│   ├── deployment/              # 部署文档
│   └── development/             # 开发文档
│
├── scripts/                     # 项目级脚本
│   ├── build.sh                 # 构建脚本
│   ├── deploy.sh                # 部署脚本
│   └── test.sh                  # 测试脚本
│
└── infrastructure/              # 基础设施
    ├── docker/                  # Docker配置
    ├── k8s/                     # Kubernetes配置
    └── terraform/               # 基础设施即代码
```

## 方案二：Multi-repo（多仓库）

```
qq-forum-backend/               # 后端仓库
├── api/
├── model/
├── cmd/
└── ...

qq-forum-frontend/              # 前端仓库
├── src/
├── public/
└── ...

qq-forum-shared/                # 共享仓库
├── types/
└── ...

qq-forum-docs/                  # 文档仓库
└── ...
```

## 当前项目迁移建议

### 第一步：创建新的目录结构
```bash
# 在项目根目录执行
mkdir -p backend frontend shared docs scripts infrastructure

# 移动现有后端代码
mv api backend/
mv model backend/
mv cmd backend/
mv database backend/
mv go.mod backend/
mv go.sum backend/
```

### 第二步：创建前端项目
```bash
# 创建React项目
cd frontend
npm create vite@latest web -- --template react-ts
cd web
npm install
```

### 第三步：配置开发环境
```bash
# 根目录创建docker-compose.yml
# 配置前后端开发环境
```

## 开发工作流

### 本地开发
```bash
# 启动所有服务
make dev

# 只启动后端
make dev-backend

# 只启动前端
make dev-frontend
```

### 构建部署
```bash
# 构建所有项目
make build

# 部署到生产环境
make deploy
```

## 优势对比

### Monorepo优势：
- ✅ 代码共享方便
- ✅ 版本管理统一
- ✅ 部署协调简单
- ✅ 开发体验好

### Multi-repo优势：
- ✅ 团队独立开发
- ✅ 技术栈独立
- ✅ 部署独立
- ✅ 权限管理精细

## 推荐选择

对于中小型项目，推荐使用 **Monorepo** 结构，因为：
1. 团队规模适中
2. 前后端关联紧密
3. 部署和维护简单
4. 开发效率高 