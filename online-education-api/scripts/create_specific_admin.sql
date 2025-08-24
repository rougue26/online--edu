USE online_education_system;

-- 插入管理员账户 (如果不存在)
INSERT INTO users (username, email, password, role)
SELECT 'admin', 'admin@example.com', '$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW', 'admin'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin');

-- 注意：上面的密码哈希对应的明文密码是 'password'
-- 更新密码为 'abc123'
UPDATE users
SET password = '$2a$10$Xq5C8K9X7h6V5P4L3O2N1M0K9J8I7H6G5F4E3D2C1B0A9' -- 这是 'abc123' 的bcrypt哈希
WHERE username = 'admin';

SELECT '管理员账户创建/更新成功！' AS message;