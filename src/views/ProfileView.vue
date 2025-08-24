<template>
  <div class="profile-container">
    <div class="profile-sidebar">
      <div class="user-info">
        <img :src="userInfo.avatar" alt="用户头像" class="avatar">
        <div class="user-basic-info">
          <h3 class="username">{{ userInfo.username }}</h3>
          <p class="email">{{ userInfo.email }}</p>
        </div>
      </div>
      <el-menu
        :default-active="activeMenu"
        class="profile-menu"
        @select="handleMenuSelect"
      >
        <el-menu-item index="info">
          <el-icon slot="icon"><User /></el-icon>
          <span slot="title">个人信息</span>
        </el-menu-item>
        <el-menu-item index="learning">
          <el-icon slot="icon"><Book /></el-icon>
          <span slot="title">学习记录</span>
        </el-menu-item>
        <el-menu-item index="favorite">
          <el-icon slot="icon"><Star /></el-icon>
          <span slot="title">我的收藏</span>
        </el-menu-item>
        <el-menu-item index="settings">
          <el-icon slot="icon"><Setting /></el-icon>
          <span slot="title">账户设置</span>
        </el-menu-item>
        <el-menu-item index="logout"
          @click.native="handleLogout"
        >
          <el-icon slot="icon"><Logout /></el-icon>
          <span slot="title">退出登录</span>
        </el-menu-item>
      </el-menu>
    </div>

    <div class="profile-content">
      <!-- 个人信息 -->
      <div v-if="activeMenu === 'info'" class="info-content">
        <el-card class="info-card">
          <div class="info-item">
            <span class="info-label">用户名</span>
            <span class="info-value">{{ userInfo.username }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">邮箱</span>
            <span class="info-value">{{ userInfo.email }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">注册时间</span>
            <span class="info-value">{{ userInfo.registerTime }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">最后登录</span>
            <span class="info-value">{{ userInfo.lastLoginTime }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">学习课程</span>
            <span class="info-value">{{ userInfo.courseCount }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">学习时长</span>
            <span class="info-value">{{ userInfo.learningTime }}</span>
          </div>
        </el-card>
      </div>

      <!-- 学习记录 -->
      <div v-if="activeMenu === 'learning'" class="learning-content">
        <div class="learning-stats">
          <el-card class="stat-card">
            <div class="stat-title">总学习时长</div>
            <div class="stat-value">{{ userInfo.learningTime }}</div>
          </el-card>
          <el-card class="stat-card">
            <div class="stat-title">已完成课程</div>
            <div class="stat-value">{{ userInfo.completedCourses }}</div>
          </el-card>
          <el-card class="stat-card">
            <div class="stat-title">进行中课程</div>
            <div class="stat-value">{{ userInfo.ongoingCourses }}</div>
          </el-card>
        </div>

        <div class="learning-courses">
          <h3 class="section-title">最近学习的课程</h3>
          <div class="course-list">
            <el-card
              v-for="course in recentCourses"
              :key="course.id"
              class="course-card"
            >
              <div class="course-info">
                <img :src="course.image" alt="课程封面" class="course-cover">
                <div class="course-detail">
                  <h4 class="course-title">{{ course.title }}</h4>
                  <p class="course-teacher">讲师：{{ course.teacher }}</p>
                  <div class="course-progress">
                    <el-progress
                      :percentage="course.progress"
                      stroke-width="6"
                      :show-text="false"
                    />
                    <span class="progress-text">{{ course.progress }}%</span>
                  </div>
                </div>
              </div>
              <div class="course-actions">
                <el-button type="primary" size="small" @click="continueLearning(course.id)">继续学习</el-button>
              </div>
            </el-card>
          </div>
        </div>
      </div>

      <!-- 我的收藏 -->
      <div v-if="activeMenu === 'favorite'" class="favorite-content">
        <h3 class="section-title">收藏的课程</h3>
        <div class="course-list">
          <el-card
            v-for="course in favoriteCourses"
            :key="course.id"
            class="course-card"
          >
            <div class="course-info">
              <img :src="course.image" alt="课程封面" class="course-cover">
              <div class="course-detail">
                <h4 class="course-title">{{ course.title }}</h4>
                <p class="course-teacher">讲师：{{ course.teacher }}</p>
                <p class="course-price">{{ course.price }}</p>
              </div>
            </div>
            <div class="course-actions">
              <el-button type="primary" size="small" @click="goToCourseDetail(course.id)">查看详情</el-button>
              <el-button type="danger" size="small" icon="StarFilled" @click="cancelFavorite(course.id)">已收藏</el-button>
            </div>
          </el-card>
        </div>
      </div>

      <!-- 账户设置 -->
      <div v-if="activeMenu === 'settings'" class="settings-content">
        <el-form ref="settingsForm" :model="settingsForm" :rules="rules" class="settings-form">
          <el-form-item prop="username">
            <el-input v-model="settingsForm.username" placeholder="用户名" />
          </el-form-item>
          <el-form-item prop="email">
            <el-input v-model="settingsForm.email" placeholder="邮箱" />
          </el-form-item>
          <el-form-item prop="oldPassword">
            <el-input v-model="settingsForm.oldPassword" type="password" placeholder="原密码" />
          </el-form-item>
          <el-form-item prop="newPassword">
            <el-input v-model="settingsForm.newPassword" type="password" placeholder="新密码" />
          </el-form-item>
          <el-form-item prop="confirmPassword">
            <el-input v-model="settingsForm.confirmPassword" type="password" placeholder="确认新密码" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="saveSettings">保存设置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import { User, Book, Star, Setting, Logout, StarFilled } from '@element-plus/icons-vue'
import { userAPI } from '@/api/index'

export default {
  name: 'ProfileView',
  components: {
    User,
    Book,
    Star,
    Setting,
    Logout,
    StarFilled
  },
  data() {
    return {
      activeMenu: 'info',
      userInfo: {
        username: '',
        email: '',
        avatar: '/photos/100/100?random=1',
        registerTime: '',
        lastLoginTime: '',
        courseCount: 0,
        learningTime: '0小时',
        completedCourses: 0,
        ongoingCourses: 0
      },
      recentCourses: [
        {
          id: 1,
          title: 'Vue.js从入门到精通',
          teacher: '张老师',
          image: '/photos/200/150?random=8',
          progress: 60
        },
        {
          id: 2,
          title: 'React实战开发',
          teacher: '李老师',
          image: '/photos/200/150?random=9',
          progress: 30
        },
        {
          id: 3,
          title: 'Node.js后端开发',
          teacher: '王老师',
          image: '/photos/200/150?random=10',
          progress: 80
        }
      ],
      favoriteCourses: [
        {
          id: 4,
          title: 'JavaScript高级编程',
          teacher: '赵老师',
          image: '/photos/200/150?random=11',
          price: '¥199'
        },
        {
          id: 5,
          title: 'TypeScript完全指南',
          teacher: '刘老师',
          image: '/photos/200/150?random=12',
          price: '¥259'
        }
      ],
      settingsForm: {
        username: '',
        email: '',
        oldPassword: '',
        newPassword: '',
        confirmPassword: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '用户名长度在3到20个字符之间', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ],
        oldPassword: [
          { required: true, message: '请输入原密码', trigger: 'blur' },
          { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
          { pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$/, message: '密码必须包含字母和数字', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请确认新密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleMenuSelect(index) {
      this.activeMenu = index
    },
    handleLogout() {
      this.$confirm('确定要退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        localStorage.removeItem('token')
        this.$message.success('退出登录成功')
        this.$router.push({ name: 'login' })
      }).catch(() => {
        // 取消操作
      })
    },
    continueLearning(courseId) {
      // 跳转到课程详情页
      this.$router.push({ name: 'courseDetail', params: { id: courseId } })
    },
    goToCourseDetail(courseId) {
      // 跳转到课程详情页
      this.$router.push({ name: 'courseDetail', params: { id: courseId } })
    },
    cancelFavorite(courseId) {
      this.$confirm('确定要取消收藏吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.favoriteCourses = this.favoriteCourses.filter(course => course.id !== courseId)
        this.$message.success('取消收藏成功')
      }).catch(() => {
        // 取消操作
      })
    },
    validateConfirmPassword(rule, value, callback) {
      if (value !== this.settingsForm.newPassword) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    },
    saveSettings() {
      this.$refs.settingsForm.validate(valid => {
        if (valid) {
          // 模拟保存设置请求
          setTimeout(() => {
            // 实际应用中，这里应该调用保存设置的API
            this.userInfo.username = this.settingsForm.username
            this.userInfo.email = this.settingsForm.email
            this.$message.success('设置保存成功')
            this.activeMenu = 'info'
          }, 1000)
        }
      })
    }
  },
  created() {
    // 检查是否登录
    if (!localStorage.getItem('token')) {
      this.$router.push({ name: 'login', query: { redirect: '/profile' } })
      return
    }

    // 调用API获取用户信息
    userAPI.getProfile()
      .then(response => {
        // 更新用户信息
        this.userInfo = {
          username: response.username,
          email: response.email,
          avatar: response.avatar || '/photos/100/100?random=1',
          registerTime: response.created_at || '2023-01-15',
          lastLoginTime: response.last_login_time || '2023-08-15 14:30',
          courseCount: response.course_count || 0,
          learningTime: response.learning_time || '0小时',
          completedCourses: response.completed_courses || 0,
          ongoingCourses: response.ongoing_courses || 0
        }
        // 更新设置表单
        this.settingsForm = {
          username: response.username,
          email: response.email,
          oldPassword: '',
          newPassword: '',
          confirmPassword: ''
        }
      })
      .catch(error => {
        console.error('获取用户信息失败:', error)
        this.$message.error('获取用户信息失败')
      })
  }
}
</script>

<style scoped>
.profile-container {
  display: flex;
  min-height: 100vh;
  background-color: #f5f7fa;
}

.profile-sidebar {
  width: 250px;
  background-color: #fff;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.user-info {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e5e5e5;
}

.avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 15px;
}

.user-basic-info {
  flex: 1;
}

.username {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 5px;
}

.email {
  color: #666;
  font-size: 14px;
}

.profile-menu {
  width: 100%;
}

.profile-content {
  flex: 1;
  padding: 30px;
}

.info-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.info-card {
  width: 100%;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 15px 0;
  border-bottom: 1px solid #e5e5e5;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  color: #999;
}

.info-value {
  font-weight: bold;
}

.learning-stats {
  display: flex;
  gap: 20px;
  margin-bottom: 30px;
  flex-wrap: wrap;
}

.stat-card {
  flex: 1;
  min-width: 200px;
  text-align: center;
}

.stat-title {
  color: #999;
  margin-bottom: 10px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #409eff;
}

.section-title {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 20px;
}

.course-list {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
}

.course-card {
  width: 100%;
  max-width: 350px;
  margin-bottom: 20px;
}

.course-info {
  display: flex;
  gap: 15px;
  margin-bottom: 15px;
}

.course-cover {
  width: 120px;
  height: 90px;
  object-fit: cover;
}

.course-detail {
  flex: 1;
}

.course-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 10px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.course-teacher {
  color: #666;
  font-size: 14px;
  margin-bottom: 10px;
}

.course-price {
  color: #f56c6c;
  font-weight: bold;
  font-size: 16px;
}

.course-progress {
  display: flex;
  align-items: center;
  gap: 10px;
}

.progress-text {
  font-size: 14px;
  color: #666;
}

.course-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.settings-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.settings-form {
  max-width: 500px;
}
</style>