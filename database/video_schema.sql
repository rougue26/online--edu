-- 创建数据库
CREATE DATABASE IF NOT EXISTS online_education_system DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE online_education_system;

-- 用户表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    nickname VARCHAR(50),
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    status TINYINT DEFAULT 1 COMMENT '1: 正常, 0: 禁用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 课程分类表
CREATE TABLE course_categories (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    parent_id BIGINT DEFAULT NULL,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES course_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 课程表
CREATE TABLE courses (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    cover_image VARCHAR(255),
    price DECIMAL(10,2) DEFAULT 0.00,
    original_price DECIMAL(10,2) DEFAULT 0.00,
    category_id BIGINT NOT NULL,
    teacher_id BIGINT NOT NULL,
    level TINYINT DEFAULT 1 COMMENT '1: 初级, 2: 中级, 3: 高级',
    duration INT DEFAULT 0 COMMENT '课程时长(分钟)',
    student_count INT DEFAULT 0,
    rating DECIMAL(3,2) DEFAULT 0.00,
    status TINYINT DEFAULT 1 COMMENT '1: 上架, 0: 下架',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES course_categories(id),
    FOREIGN KEY (teacher_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 章节表
CREATE TABLE chapters (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    course_id BIGINT NOT NULL,
    title VARCHAR(100) NOT NULL,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 课时表
CREATE TABLE lessons (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    chapter_id BIGINT NOT NULL,
    title VARCHAR(100) NOT NULL,
    video_url VARCHAR(255),
    duration INT DEFAULT 0 COMMENT '课时时长(分钟)',
    sort_order INT DEFAULT 0,
    free TINYINT DEFAULT 0 COMMENT '1: 免费, 0: 收费',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (chapter_id) REFERENCES chapters(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户课程关系表 (记录用户购买/学习的课程)
CREATE TABLE user_courses (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    course_id BIGINT NOT NULL,
    order_id VARCHAR(50),
    payment_amount DECIMAL(10,2) DEFAULT 0.00,
    payment_time TIMESTAMP,
    status TINYINT DEFAULT 1 COMMENT '1: 已购买, 0: 已退款',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_course (user_id, course_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 社区帖子表
CREATE TABLE posts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    category VARCHAR(50),
    view_count INT DEFAULT 0,
    comment_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 评论表
CREATE TABLE comments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    content TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    parent_id BIGINT DEFAULT NULL,
    like_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 点赞表
CREATE TABLE likes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    target_type TINYINT NOT NULL COMMENT '1: 帖子, 2: 评论',
    target_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_target (user_id, target_type, target_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 学习进度表
CREATE TABLE learning_progress (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    lesson_id BIGINT NOT NULL,
    progress INT DEFAULT 0 COMMENT '学习进度(秒)',
    completed TINYINT DEFAULT 0 COMMENT '1: 已完成, 0: 未完成',
    last_learned_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_lesson (user_id, lesson_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引优化
CREATE INDEX idx_courses_category ON courses(category_id);
CREATE INDEX idx_courses_teacher ON courses(teacher_id);
CREATE INDEX idx_posts_user ON posts(user_id);
CREATE INDEX idx_comments_post ON comments(post_id);
CREATE INDEX idx_comments_user ON comments(user_id);
-- 图片表：存储图片元数据
CREATE TABLE images (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(512) NOT NULL,
    file_size INT NOT NULL COMMENT '文件大小(字节)',
    mime_type VARCHAR(100) NOT NULL COMMENT 'MIME类型',
    width INT COMMENT '图片宽度(像素)',
    height INT COMMENT '图片高度(像素)',
    alt_text VARCHAR(255) COMMENT '替代文本',
    entity_type VARCHAR(50) COMMENT '关联实体类型(如user, course, post等)',
    entity_id BIGINT COMMENT '关联实体ID',
    is_primary TINYINT DEFAULT 0 COMMENT '是否为主图(1:是, 0:否)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 索引优化
CREATE INDEX idx_images_entity ON images(entity_type, entity_id);
CREATE INDEX idx_images_file_path ON images(file_path);
-- 视频分类表
CREATE TABLE IF NOT EXISTS video_categories (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL COMMENT '分类名称',
  description TEXT COMMENT '分类描述',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频分类表';

-- 视频表
CREATE TABLE IF NOT EXISTS videos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL COMMENT '视频标题',
  description TEXT COMMENT '视频描述',
  video_url VARCHAR(512) NOT NULL COMMENT '视频文件路径',
  cover_image_url VARCHAR(512) COMMENT '封面图片路径',
  duration INT NOT NULL DEFAULT 0 COMMENT '视频时长(秒)',
  author_id BIGINT NOT NULL COMMENT '作者ID',
  category_id INT COMMENT '分类ID',
  view_count INT NOT NULL DEFAULT 0 COMMENT '观看次数',
  like_count INT NOT NULL DEFAULT 0 COMMENT '点赞次数',
  favorite_count INT NOT NULL DEFAULT 0 COMMENT '收藏次数',
  is_public BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否公开',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_author_id (author_id),
  INDEX idx_category_id (category_id),
  INDEX idx_created_at (created_at),
  INDEX idx_view_count (view_count),
  FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES video_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

-- 视频点赞表
CREATE TABLE IF NOT EXISTS video_likes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  video_id INT NOT NULL COMMENT '视频ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_video_user (video_id, user_id),
  INDEX idx_user_id (user_id),
  FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频点赞表';

-- 视频收藏表
CREATE TABLE IF NOT EXISTS video_favorites (
  id INT AUTO_INCREMENT PRIMARY KEY,
  video_id INT NOT NULL COMMENT '视频ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_video_user (video_id, user_id),
  INDEX idx_user_id (user_id),
  FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频收藏表';

-- 视频评论表
CREATE TABLE IF NOT EXISTS video_comments (
  id INT AUTO_INCREMENT PRIMARY KEY,
  video_id INT NOT NULL COMMENT '视频ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  content TEXT NOT NULL COMMENT '评论内容',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_video_id (video_id),
  INDEX idx_user_id (user_id),
  INDEX idx_created_at (created_at),
  FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频评论表';

-- 视频评论点赞表
CREATE TABLE IF NOT EXISTS video_comment_likes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  comment_id INT NOT NULL COMMENT '评论ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_comment_user (comment_id, user_id),
  INDEX idx_user_id (user_id),
  FOREIGN KEY (comment_id) REFERENCES video_comments(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频评论点赞表';

-- 初始视频分类数据
INSERT IGNORE INTO video_categories (name, description) VALUES
('编程开发', '编程语言、框架、开发工具等相关视频'),
('设计创意', 'UI设计、平面设计、创意设计等相关视频'),
('职业技能', '职场能力、沟通技巧、项目管理等相关视频'),
('兴趣爱好', '音乐、运动、美食等兴趣爱好相关视频');

-- 修改权限检查视图，移除对u.role的引用
CREATE OR REPLACE VIEW user_video_permissions AS
SELECT
  v.id AS video_id,
  v.author_id,
  u.id AS user_id,
  (v.author_id = u.id) AS is_author,
  (v.is_public OR v.author_id = u.id) AS can_view
FROM videos v
CROSS JOIN users u;