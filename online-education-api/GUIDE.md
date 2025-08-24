# 在线教育系统后端开发指南

## 项目介绍
这是一个在线教育系统的后端API服务，提供用户管理、课程管理和社区交流等功能。系统采用Go语言开发，使用Gorilla Mux作为路由框架，MySQL作为数据库。

## 技术栈
- 编程语言：Go 1.18+
- Web框架：Gorilla Mux
- 数据库：MySQL 8.0
- 认证：JWT (JSON Web Token)
- 依赖管理：Go Modules

## 目录结构
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
│   ├── user_course.go         # 用户课程关系模型
│   └── ...                    # 其他模型
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

## 环境配置
### 1. 安装Go
请确保已安装Go 1.18或更高版本。可以从[Go官网](https://golang.org/)下载并安装。

### 2. 安装MySQL
安装MySQL 8.0数据库，并创建一个名为`online_education_system`的数据库。可以使用提供的`scripts/init_db.sql`脚本初始化数据库结构。

### 3. 配置数据库连接
修改`config/db.go`文件中的数据库连接信息，确保与你的本地MySQL配置匹配。

### 4. 安装依赖
在项目根目录下执行以下命令安装依赖：
```bash
cd /Users/linmumu/Desktop/online-education-api
go mod tidy
```

## 运行方法
### 1. 启动后端服务
在项目根目录下执行以下命令：
```bash
go run main.go
```
服务将启动在`http://localhost:8081`端口。

### 2. 测试依赖
可以使用`test_deps.go`文件测试数据库连接、JWT工具和服务初始化：
```bash
go run test_deps.go
```

## API接口文档
### 用户接口
- `POST /api/users/register` - 用户注册
- `POST /api/users/login` - 用户登录
- `GET /api/users/profile` - 获取用户资料 (需要认证)
- `PUT /api/users/profile` - 更新用户资料 (需要认证)
- `PUT /api/users/change-password` - 修改密码 (需要认证)

### 课程分类接口
- `GET /api/course-categories` - 获取所有课程分类
- `GET /api/course-categories/{id}` - 获取单个课程分类
- `POST /api/course-categories` - 创建课程分类 (需要认证)
- `PUT /api/course-categories/{id}` - 更新课程分类 (需要认证)
- `DELETE /api/course-categories/{id}` - 删除课程分类 (需要认证)

### 课程接口
- `GET /api/courses` - 获取课程列表
- `GET /api/courses/{id}` - 获取课程详情
- `POST /api/courses` - 创建课程 (需要认证)
- `PUT /api/courses/{id}` - 更新课程 (需要认证，仅教师)
- `DELETE /api/courses/{id}` - 删除课程 (需要认证，仅教师)

### 用户课程接口
- `POST /api/user-courses/{courseID}` - 报名课程 (需要认证)
- `GET /api/user-courses` - 获取用户报名的课程列表 (需要认证)
- `GET /api/user-courses/{courseID}` - 检查用户是否报名课程 (需要认证)
- `DELETE /api/user-courses/{courseID}` - 取消报名课程 (需要认证)

### 社区帖子接口
- `GET /api/posts` - 获取帖子列表
- `GET /api/posts/{id}` - 获取帖子详情
- `POST /api/posts` - 创建帖子 (需要认证)
- `PUT /api/posts/{id}` - 更新帖子 (需要认证，仅作者)
- `DELETE /api/posts/{id}` - 删除帖子 (需要认证，仅作者)
- `GET /api/users/posts` - 获取用户发布的帖子 (需要认证)

## 测试方法
### 1. 单元测试
可以为各个服务编写单元测试。测试文件应放在对应包下，命名为`xxx_test.go`。

### 2. API测试
可以使用Postman、curl或其他API测试工具测试API接口。

### 3. 集成测试
可以使用`test_deps.go`作为基础，扩展编写更全面的集成测试。

## 部署说明
### 1. 编译项目
在项目根目录下执行以下命令编译项目：
```bash
go build -o online-education-api
```

### 2. 运行编译后的二进制文件
```bash
./online-education-api
```

### 3. 生产环境配置
在生产环境中，建议：
- 使用环境变量或配置文件管理敏感信息
- 设置适当的日志级别
- 配置HTTPS
- 使用反向代理(如Nginx)转发请求
- 考虑使用Docker容器化部署

## 扩展建议
1. 添加缓存机制(如Redis)提高性能
2. 实现文件上传功能，用于课程封面、视频等
3. 添加邮件通知系统
4. 实现支付系统
5. 添加管理员后台
6. 优化搜索功能

希望这个指南对你有所帮助！如果有任何问题或建议，请随时提出。