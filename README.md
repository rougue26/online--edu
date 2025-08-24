# 在线教育系统

## 项目简介
这是一个用go+vue3做的在线教育系统，包含前端和后端两部分。系统提供课程管理、用户管理、学习记录、社区交流和视频播放等功能，采用现代化的前后端分离架构设计，具有良好的可扩展性和维护性，可能会有些bug（比如支付功能并未测试），欢迎使用者指正

## 技术栈
### 前端
- **框架**: Vue 3
- **路由**: Vue Router 4
- **UI组件库**: Element Plus
- **HTTP客户端**: Axios
- **构建工具**: Vue CLI
- **代码规范**: ESLint

### 后端
- **编程语言**: Go 1.18+
- **Web框架**: Gorilla Mux
- **数据库**: MySQL 8.0
- **认证机制**: JWT (JSON Web Token)
- **依赖管理**: Go Modules

## 项目结构
```
在线教育系统/
├── .gitignore
├── README.md
├── babel.config.js
├── database/
│   └── video_schema.sql
├── jsconfig.json
├── online-education-api/       # 后端API目录
│   ├── BACKEND_EXTENSION_PLAN.md
│   ├── GUIDE.md
│   ├── config/                 # 配置文件
│   ├── controllers/            # 控制器
│   ├── generate_admin_password.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go                 # 后端入口文件
│   ├── middleware/             # 中间件
│   ├── models/                 # 数据模型
│   ├── routes/                 # 路由配置
│   ├── scripts/                # 脚本
│   ├── services/               # 服务层
│   ├── test_deps
│   ├── test_deps.go
│   ├── tools/
│   └── utils/                  # 工具类
├── online-education-system/    # 前端目录
│   ├── .gitignore
│   ├── babel.config.js
│   ├── database/
│   ├── jsconfig.json
│   ├── package-lock.json
│   ├── package.json
│   ├── public/
│   ├── src/
│   │   ├── App.vue
│   │   ├── api/                # API请求
│   │   ├── assets/             # 静态资源
│   │   ├── components/         # 公共组件
│   │   ├── main.js             # 前端入口文件
│   │   ├── router/             # 路由配置
│   │   └── views/              # 页面视图
│   └── vue.config.js
├── online_education_system.sql
├── package-lock.json
├── package.json
├── public/
├── src/
│   ├── App.vue
│   ├── api/
│   ├── assets/
│   ├── components/
│   ├── main.js
│   ├── router/
│   └── views/
└── vue.config.js
```

## 核心功能
1. **用户管理**
   - 注册、登录、个人信息管理
   - 密码修改、权限控制
   - 管理员后台用户管理

2. **课程管理**
   - 课程创建、查询、更新、删除
   - 课程分类管理
   - 章节和课时管理

3. **学习记录管理**
   - 课程报名、学习进度跟踪
   - 视频观看记录
   - 学习统计分析

4. **社区交流**
   - 帖子创建、查询、更新、删除
   - 评论和回复功能
   - 点赞和收藏功能

5. **视频管理**
   - 视频上传、查询、播放
   - 视频分类管理
   - 视频评论和点赞

6. **支付功能**
   - 课程购买支付接口
   - 支付记录查询

## 快速开始
### 环境要求
- Go 1.18+ 已安装
- MySQL 8.0 已安装并运行
- Node.js 14+ 已安装
- npm 6+ 已安装

### 后端启动步骤
1. **进入后端目录**
   ```bash
   cd online-education-api
   ```

2. **配置数据库**
   - 创建数据库 `online_education_system`
   - 运行初始化脚本 `scripts/init_db.sql`
   - 修改 `config/db.go` 中的数据库连接信息

3. **安装依赖**
   ```bash
   go mod tidy
   ```

4. **运行后端服务**
   ```bash
   go run main.go
   ```
   服务将启动在 `http://localhost:8081`

### 前端启动步骤
1. **进入前端目录**
   ```bash
   cd online-education-system
   ```

2. **安装依赖**
   ```bash
   npm install
   ```

3. **运行前端服务**
   ```bash
   npm run serve
   ```
   服务将启动在 `http://localhost:8080`

## API文档
API文档详细说明了所有可用接口、请求参数和响应格式。请参考 `online-education-api/GUIDE.md` 文件获取更多信息。

## 测试
- 后端依赖测试：`cd online-education-api && go run test_deps.go`
- 后端单元测试：`cd online-education-api && go test ./...`
- 前端测试：`cd online-education-system && npm run test`

## 部署
### 后端部署
1. 编译后端代码
   ```bash
   cd online-education-api
   go build -o online-education-api main.go
   ```

2. 部署编译后的二进制文件到服务器

### 前端部署
1. 构建前端代码
   ```bash
   cd online-education-system
   npm run build
   ```

2. 将 `dist` 目录下的文件部署到Web服务器

## 扩展计划
请参考 `online-education-api/BACKEND_EXTENSION_PLAN.md` 文件了解后端扩展计划。

## 贡献
欢迎对本项目进行贡献，您可以通过提交Issue或Pull Request的方式参与项目开发。
