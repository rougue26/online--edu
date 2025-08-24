# 在线教育系统后端API

## 项目简介
这是一个基于Go语言开发的在线教育系统后端API，提供课程管理、用户管理、学习记录和社区交流等功能。系统采用分层架构设计，具有良好的可扩展性和维护性。

## 核心功能
1. **用户管理**：注册、登录、个人信息管理、密码修改
2. **课程分类管理**：创建、查询、更新、删除课程分类
3. **课程管理**：创建、查询、更新、删除课程
4. **学习记录管理**：课程报名、学习进度跟踪
5. **社区交流**：帖子创建、查询、更新、删除
6. **视频管理**：视频上传、查询、播放

## 技术栈
- **编程语言**：Go 1.18+
- **Web框架**：Gorilla Mux
- **数据库**：MySQL 8.0
- **认证机制**：JWT (JSON Web Token)
- **依赖管理**：Go Modules

## 项目结构
```
online-education-api/
├── BACKEND_EXTENSION_PLAN.md  # 后端扩展计划
├── GUIDE.md                   # 开发指南
├── README.md                  # 项目说明
├── config/                    # 配置文件
│   └── db.go                  # 数据库配置
├── controllers/               # 控制器
│   ├── course_category_controller.go  # 课程分类控制器
│   ├── course_controller.go           # 课程控制器
│   ├── post_controller.go             # 帖子控制器
│   ├── user_controller.go             # 用户控制器
│   ├── user_course_controller.go      # 用户课程控制器
│   └── video_controller.go            # 视频控制器
├── go.mod                     # Go模块定义
├── go.sum                     # 依赖列表
├── main.go                    # 入口文件
├── middleware/                # 中间件
│   └── auth.go                # JWT认证中间件
├── models/                    # 数据模型
│   ├── chapter.go             # 章节模型
│   ├── course.go              # 课程模型
│   ├── course_category.go     # 课程分类模型
│   ├── lesson.go              # 课时模型
│   ├── post.go                # 帖子模型
│   ├── user.go                # 用户模型
│   └── user_course.go         # 用户课程关系模型
├── routes/                    # 路由配置
│   └── routes.go              # 路由定义
├── scripts/                   # 脚本
│   └── init_db.sql            # 数据库初始化脚本
├── services/                  # 服务层
│   ├── course_category_service.go  # 课程分类服务
│   ├── course_service.go           # 课程服务
│   ├── post_service.go             # 帖子服务
│   ├── user_course_service.go      # 用户课程服务
│   ├── user_service.go             # 用户服务
│   └── video_service.go            # 视频服务
├── test_deps.go               # 依赖测试文件
└── utils/                     # 工具类
    └── jwt.go                 # JWT工具
```

## 快速开始
### 环境要求
- Go 1.18+ 已安装
- MySQL 8.0 已安装并运行

### 步骤
1. **克隆项目**
   ```bash
   git clone https://github.com/yourusername/online-education-api.git
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

4. **运行项目**
   ```bash
   go run main.go
   ```
   服务将启动在 `http://localhost:8081`

## API文档
API文档详细说明了所有可用接口、请求参数和响应格式。请参考 `GUIDE.md` 文件获取更多信息。

## 测试
- 运行依赖测试：`go run test_deps.go`
- 单元测试：`go test ./...`

## 部署
1. **编译项目**
   ```bash
   go build -o online-education-api
   ```

2. **运行编译后的二进制文件**
   ```bash
   ./online-education-api
   ```

## 贡献指南
1.  Fork 项目
2.  创建特性分支 (`git checkout -b feature/your-feature`)
3.  提交更改 (`git commit -am 'Add some feature'`)
4.  推送到分支 (`git push origin feature/your-feature`)
5.  创建 Pull Request

## 许可证
本项目采用 MIT 许可证。详情请见 LICENSE 文件。

## 联系方式
如有任何问题或建议，请联系：
- 邮箱：example@example.com
- GitHub：@yourusername