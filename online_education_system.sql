-- MySQL dump 10.13  Distrib 8.4.6, for macos15 (x86_64)
--
-- Host: localhost    Database: online_education_system
-- ------------------------------------------------------
-- Server version	8.4.6

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `chapters`
--

DROP TABLE IF EXISTS `chapters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chapters` (
  `id` int NOT NULL AUTO_INCREMENT,
  `course_id` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `sort_order` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  CONSTRAINT `chapters_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chapters`
--

LOCK TABLES `chapters` WRITE;
/*!40000 ALTER TABLE `chapters` DISABLE KEYS */;
INSERT INTO `chapters` VALUES (1,1,'Go语言简介',1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,1,'基本语法',2,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(3,2,'Python基础',1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(4,2,'Pandas数据分析',2,'2025-08-22 12:44:34','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `chapters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course_categories`
--

DROP TABLE IF EXISTS `course_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `course_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `parent_id` int DEFAULT NULL,
  `sort_order` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `course_categories_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `course_categories` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_categories`
--

LOCK TABLES `course_categories` WRITE;
/*!40000 ALTER TABLE `course_categories` DISABLE KEYS */;
INSERT INTO `course_categories` VALUES (1,'编程开发',NULL,0,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,'数据分析',NULL,0,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(3,'人工智能',NULL,0,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(4,'设计创意',NULL,0,'2025-08-22 12:44:34','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `course_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `cover_image` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT '0.00',
  `original_price` decimal(10,2) DEFAULT '0.00',
  `category_id` int DEFAULT NULL,
  `teacher_id` int DEFAULT NULL,
  `level` int DEFAULT '1',
  `duration` int DEFAULT '0',
  `student_count` int DEFAULT '0',
  `rating` decimal(3,2) DEFAULT '0.00',
  `status` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_courses_title` (`title`),
  KEY `idx_courses_category_id` (`category_id`),
  KEY `idx_courses_teacher_id` (`teacher_id`),
  CONSTRAINT `courses_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `course_categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `courses_ibfk_2` FOREIGN KEY (`teacher_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,'Go语言基础入门','掌握Go语言的基本语法和编程技巧','https://example.com/go_cover.jpg',99.00,199.00,1,2,1,360,0,0.00,1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,'Python数据分析实战','使用Python进行数据清洗、分析和可视化','https://example.com/python_cover.jpg',129.00,259.00,2,2,2,480,0,0.00,1,'2025-08-22 12:44:34','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lessons`
--

DROP TABLE IF EXISTS `lessons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lessons` (
  `id` int NOT NULL AUTO_INCREMENT,
  `chapter_id` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `video_url` varchar(255) DEFAULT NULL,
  `duration` int DEFAULT '0',
  `sort_order` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `chapter_id` (`chapter_id`),
  CONSTRAINT `lessons_ibfk_1` FOREIGN KEY (`chapter_id`) REFERENCES `chapters` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lessons`
--

LOCK TABLES `lessons` WRITE;
/*!40000 ALTER TABLE `lessons` DISABLE KEYS */;
INSERT INTO `lessons` VALUES (1,1,'Go语言的起源与特点','https://example.com/video1.mp4',300,1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,1,'环境搭建','https://example.com/video2.mp4',420,2,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(3,2,'变量与数据类型','https://example.com/video3.mp4',360,1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(4,3,'Python安装与配置','https://example.com/video4.mp4',300,1,'2025-08-22 12:44:34','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `lessons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `category` varchar(100) DEFAULT '未分类',
  `views` int DEFAULT '0',
  `comment_count` int DEFAULT '0',
  `like_count` int DEFAULT '0',
  `status` enum('draft','published','archived') DEFAULT 'published',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_posts_user_id` (`user_id`),
  KEY `idx_posts_title` (`title`),
  CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (1,1,'第一个帖子','这是一个测试帖子的内容。','技术讨论',0,0,0,'published','2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,2,'第二个帖子','这是第二个测试帖子的内容。','学习资源',0,0,0,'published','2025-08-22 12:44:34','2025-08-22 12:44:34'),(3,3,'第三个帖子','这是第三个测试帖子的内容。','未分类',0,0,0,'draft','2025-08-22 12:44:34','2025-08-22 12:44:34'),(4,4,'这是一个测试帖子','瞬间觉得撒娇就睡觉啊计算机计算机技术','未分类',0,0,0,'published','2025-08-22 04:55:59','2025-08-22 04:55:59'),(5,4,'test这是一个','test这是一个test67699','未分类',0,0,0,'published','2025-08-22 05:11:09','2025-08-22 05:11:09'),(6,4,'测试分类项目','1234567890','未分类',0,0,0,'published','2025-08-22 05:18:58','2025-08-22 05:18:58'),(7,4,'number123','uuuuuuuuuuuuuuuuuu','求职经验',0,0,0,'published','2025-08-22 05:25:24','2025-08-22 05:25:24'),(8,4,'1234567test','12345678ueiuhdskja','资源分享',0,0,0,'published','2025-08-22 05:35:23','2025-08-22 05:35:23'),(9,4,'test 最后一次测试','最后一次测试 1234567','资源分享',0,0,0,'published','2025-08-23 17:02:57','2025-08-23 17:02:57');
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_courses`
--

DROP TABLE IF EXISTS `user_courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_courses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `course_id` int NOT NULL,
  `enrolled_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_course` (`user_id`,`course_id`),
  KEY `idx_user_courses_user_id` (`user_id`),
  KEY `idx_user_courses_course_id` (`course_id`),
  CONSTRAINT `user_courses_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_courses_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_courses`
--

LOCK TABLES `user_courses` WRITE;
/*!40000 ALTER TABLE `user_courses` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('student','teacher','admin') DEFAULT 'student',
  `avatar` varchar(255) DEFAULT NULL,
  `bio` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_username` (`username`),
  KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','admin@example.com','$2a$10$u0IV1arIZLDAK7iTpgVQN.fZCORRXFGwXczAODfLkWlfj91mUsWf.','admin',NULL,NULL,'2025-08-22 12:44:34','2025-08-24 00:42:23'),(2,'teacher1','teacher1@example.com','$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW','teacher',NULL,NULL,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(3,'student1','student1@example.com','$2a$10$EixZaYVK1fsbw1Zfbx3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW','student',NULL,NULL,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(4,'test678','123@qq.com','$2a$10$jBdATrM5.1D9g6cdgq43PezTCJ6dIIyDksXSayohGtjLIjGHn3/mu','student',NULL,NULL,'2025-08-22 04:45:56','2025-08-22 04:45:56'),(5,'admin_new','admin_new@example.com','$2a$10$JheYIxjIvcxtJx1rivzQneWrv6BE8K123Eyrp6rixNE4PSwgjuZKu','admin',NULL,NULL,'2025-08-22 22:24:42','2025-08-22 22:24:42');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video_categories`
--

DROP TABLE IF EXISTS `video_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `video_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video_categories`
--

LOCK TABLES `video_categories` WRITE;
/*!40000 ALTER TABLE `video_categories` DISABLE KEYS */;
INSERT INTO `video_categories` VALUES (1,'技术教程','2025-08-22 12:44:34'),(2,'职业发展','2025-08-22 12:44:34'),(3,'生活技能','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `video_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video_comments`
--

DROP TABLE IF EXISTS `video_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `video_comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `video_id` int NOT NULL,
  `user_id` int NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `video_id` (`video_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `video_comments_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`),
  CONSTRAINT `video_comments_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video_comments`
--

LOCK TABLES `video_comments` WRITE;
/*!40000 ALTER TABLE `video_comments` DISABLE KEYS */;
/*!40000 ALTER TABLE `video_comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `videos`
--

DROP TABLE IF EXISTS `videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `videos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `video_url` varchar(255) NOT NULL,
  `cover_image_url` varchar(255) DEFAULT NULL,
  `duration` int DEFAULT '0',
  `user_id` int DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  `view_count` int DEFAULT '0',
  `like_count` int DEFAULT '0',
  `favorite_count` int DEFAULT '0',
  `is_public` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (1,'Git版本控制基础','学习Git的基本操作和工作流程','https://example.com/git.mp4','https://example.com/git_cover.jpg',600,2,1,0,0,0,1,'2025-08-22 12:44:34','2025-08-22 12:44:34'),(2,'RESTful API设计','学习如何设计高质量的RESTful API','https://example.com/rest.mp4','https://example.com/rest_cover.jpg',720,2,1,0,0,0,1,'2025-08-22 12:44:34','2025-08-22 12:44:34');
/*!40000 ALTER TABLE `videos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-24  9:22:52
