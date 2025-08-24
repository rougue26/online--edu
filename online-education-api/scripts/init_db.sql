-- 创建数据库
DROP DATABASE IF EXISTS online_education_system;
CREATE DATABASE IF NOT EXISTS online_education_system;
USE online_education_system;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('student', 'teacher', 'admin') DEFAULT 'student',
    avatar VARCHAR(255),
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 课程分类表
CREATE TABLE IF NOT EXISTS course_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    parent_id INT,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES course_categories(id) ON DELETE CASCADE
);

-- 课程表
CREATE TABLE IF NOT EXISTS courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    cover_image VARCHAR(255),
    price DECIMAL(10, 2) DEFAULT 0.00,
    original_price DECIMAL(10, 2) DEFAULT 0.00,
    category_id INT,
    teacher_id INT,
    level INT DEFAULT 1,
    duration INT DEFAULT 0,
    student_count INT DEFAULT 0,
    rating DECIMAL(3, 2) DEFAULT 0.00,
    status INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES course_categories(id) ON DELETE SET NULL,
    FOREIGN KEY (teacher_id) REFERENCES users(id) ON DELETE SET NULL
);

-- 章节表
CREATE TABLE IF NOT EXISTS chapters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

-- 课时表
CREATE TABLE IF NOT EXISTS lessons (
    id INT AUTO_INCREMENT PRIMARY KEY,
    chapter_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    video_url VARCHAR(255),
    duration INT DEFAULT 0,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (chapter_id) REFERENCES chapters(id) ON DELETE CASCADE
);

-- 用户课程关系表
CREATE TABLE IF NOT EXISTS user_courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    course_id INT NOT NULL,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_course (user_id, course_id)
);

-- 社区帖子表
CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) DEFAULT '未分类',
    views INT DEFAULT 0,
    comment_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    status ENUM('draft', 'published', 'archived') DEFAULT 'published',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 视频表
CREATE TABLE IF NOT EXISTS videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    video_url VARCHAR(255) NOT NULL,
    cover_image_url VARCHAR(255),
    duration INT DEFAULT 0,
    user_id INT,
    category_id INT,
    view_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    favorite_count INT DEFAULT 0,
    is_public BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

-- 视频分类表
CREATE TABLE IF NOT EXISTS video_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入初始数据（如果不存在）
INSERT INTO users (username, email, password, role)
SELECT 'admin', 'admin@example.com', '$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW', 'admin'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin');

INSERT INTO users (username, email, password, role)
SELECT 'teacher1', 'teacher1@example.com', '$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW', 'teacher'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'teacher1');

INSERT INTO users (username, email, password, role)
SELECT 'student1', 'student1@example.com', '$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW', 'student'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'student1');

-- 密码都是: password

INSERT INTO course_categories (name) VALUES
('编程开发'),
('数据分析'),
('人工智能'),
('设计创意');

INSERT INTO courses (title, description, cover_image, price, original_price, category_id, teacher_id, level, duration, status) VALUES
('Go语言基础入门', '掌握Go语言的基本语法和编程技巧', 'https://example.com/go_cover.jpg', 99.00, 199.00, 1, 2, 1, 360, 1),
('Python数据分析实战', '使用Python进行数据清洗、分析和可视化', 'https://example.com/python_cover.jpg', 129.00, 259.00, 2, 2, 2, 480, 1);

INSERT INTO chapters (course_id, title, sort_order) VALUES
(1, 'Go语言简介', 1),
(1, '基本语法', 2),
(2, 'Python基础', 1),
(2, 'Pandas数据分析', 2);

INSERT INTO lessons (chapter_id, title, video_url, duration, sort_order) VALUES
(1, 'Go语言的起源与特点', 'https://example.com/video1.mp4', 300, 1),
(1, '环境搭建', 'https://example.com/video2.mp4', 420, 2),
(2, '变量与数据类型', 'https://example.com/video3.mp4', 360, 1),
(3, 'Python安装与配置', 'https://example.com/video4.mp4', 300, 1);

INSERT INTO video_categories (name) VALUES
('技术教程'),
('职业发展'),
('生活技能');

INSERT INTO videos (title, description, video_url, cover_image_url, duration, user_id, category_id, is_public) VALUES
('Git版本控制基础', '学习Git的基本操作和工作流程', 'https://example.com/git.mp4', 'https://example.com/git_cover.jpg', 600, 2, 1, true),
('RESTful API设计', '学习如何设计高质量的RESTful API', 'https://example.com/rest.mp4', 'https://example.com/rest_cover.jpg', 720, 2, 1, true);

-- 插入测试帖子
INSERT INTO posts (title, content, user_id, category, status, created_at, updated_at) VALUES
('第一个帖子', '这是一个测试帖子的内容。', 1, '技术讨论', 'published', NOW(), NOW()),
('第二个帖子', '这是第二个测试帖子的内容。', 2, '学习资源', 'published', NOW(), NOW()),
('第三个帖子', '这是第三个测试帖子的内容。', 3, '未分类', 'draft', NOW(), NOW());

-- 添加索引以提高查询性能
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_courses_title ON courses(title);
CREATE INDEX idx_courses_category_id ON courses(category_id);
CREATE INDEX idx_courses_teacher_id ON courses(teacher_id);
CREATE INDEX idx_user_courses_user_id ON user_courses(user_id);
CREATE INDEX idx_user_courses_course_id ON user_courses(course_id);
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_title ON posts(title);

-- 完成
SELECT '数据库初始化完成！' AS message;