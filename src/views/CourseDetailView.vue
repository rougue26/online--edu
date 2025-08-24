<template>
  <div class="course-detail-container">
    <div v-if="loading" class="loading-container">
      <el-loading indicator-color="#409EFF" background="rgba(255, 255, 255, 0.8)"></el-loading>
    </div>
    <div v-else-if="error" class="error-container">
      <div class="error-message">{{ error }}</div>
    </div>
    <div v-else class="course-content">
      <div class="course-header">
        <img v-if="course.image" :src="course.image" alt="课程封面" class="course-cover"><div v-else class="course-cover-placeholder">加载中...</div>
        <div class="course-basic-info">
          <h1>{{ course.title }}</h1>
        </div>
        <div class="course-teacher">讲师：{{ course.teacher }}</div>
        <div class="course-rating">
          <el-rate v-model="course.rating" disabled :max="5" /> 
          <span class="rating-score">{{ course.rating }}</span>
          <span class="rating-count">({{ course.ratingCount }}人评价)</span>
        </div>
        <div class="course-price">
          <span class="current-price">{{ course.price }}</span>
          <span class="original-price" v-if="course.originalPrice">{{ course.originalPrice }}</span>
        </div>
        <div class="course-actions">
          <el-button type="primary" size="large" @click="purchaseCourse">立即购买</el-button>
          <el-button size="large" @click="addToCart">加入购物车</el-button>
        </div>
      </div>
    </div>

    <div class="course-tabs">
      <el-tabs v-model="activeTab" type="card">
        <el-tab-pane label="课程介绍" name="introduction" /> 
        <el-tab-pane label="课程章节" name="chapters" /> 
        <el-tab-pane label="讲师信息" name="teacher" /> 
        <el-tab-pane label="学员评价" name="reviews" /> 
      </el-tabs>
    </div>

    <div class="tab-content">
      <!-- 课程介绍 -->
      <div v-if="activeTab === 'introduction'" class="introduction-content">
        <div class="course-description" v-html="course.description"></div>
        <div class="course-info-cards">
          <el-card class="info-card">
            <div class="info-item">
              <span class="info-label">课程时长</span>
              <span class="info-value">{{ course.duration }}</span>
            </div>
          </el-card>
          <el-card class="info-card">
            <div class="info-item">
              <span class="info-label">课程难度</span>
              <span class="info-value">{{ course.level }}</span>
            </div>
          </el-card>
          <el-card class="info-card">
            <div class="info-item">
              <span class="info-label">学习人数</span>
              <span class="info-value">{{ course.studentCount }}</span>
            </div>
          </el-card>
          <el-card class="info-card">
            <div class="info-item">
              <span class="info-label">更新时间</span>
              <span class="info-value">{{ course.updateTime }}</span>
            </div>
          </el-card>
        </div>
      </div>

      <!-- 课程章节 -->
      <div v-if="activeTab === 'chapters'" class="chapters-content">
        <div v-for="chapter in course.chapters" :key="chapter.id" class="chapter">
          <div class="chapter-title" @click="toggleChapter(chapter.id)">
            <el-icon v-if="expandedChapters.includes(chapter.id)"><ChevronDown /></el-icon>
            <el-icon v-else><ChevronRight /></el-icon>
            <span>{{ chapter.title }}</span>
            <span class="chapter-duration">{{ chapter.duration }}</span>
          </div>
          <div v-if="expandedChapters.includes(chapter.id)" class="chapter-videos">
            <div v-for="video in chapter.videos" :key="video.id" class="video-item">
              <div class="video-title">
                <el-icon v-if="video.isFree"><Unlock /></el-icon>
                <el-icon v-else><Lock /></el-icon>
                <span>{{ video.title }}</span>
              </div>
              <div class="video-duration">{{ video.duration }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 讲师信息 -->
      <div v-if="activeTab === 'teacher'" class="teacher-content">
        <img v-if="course && course.teacherInfo && course.teacherInfo.avatar" :src="course.teacherInfo.avatar" alt="讲师头像" class="teacher-avatar"><div v-else class="avatar-placeholder">加载中...</div>
        <div class="teacher-info">
          <h3>{{ course.teacher }}</h3>
          <p class="teacher-title">{{ course.teacherInfo.title }}</p>
          <p class="teacher-intro" v-html="course.teacherInfo.introduction"></p>
        </div>
      </div>

      <!-- 学员评价 -->
      <div v-if="activeTab === 'reviews'" class="reviews-content">
        <div class="review-filter">
          <el-select v-model="reviewFilter" placeholder="筛选评价">
            <el-option label="全部评价" value="all" /> 
            <el-option label="好评" value="positive" /> 
            <el-option label="中评" value="neutral" /> 
            <el-option label="差评" value="negative" /> 
          </el-select>
        </div>
        <div class="review-item" v-for="review in paginatedReviews" :key="review.id">
          <img :src="review.userAvatar" alt="用户头像" class="user-avatar">
          <div class="review-content">
            <div class="review-header">
              <span class="user-name">{{ review.userName }}</span>
              <el-rate v-model="review.rating" disabled :max="5" size="small" /> 
              <span class="review-time">{{ review.time }}</span>
            </div>
            <div class="review-text">{{ review.content }}</div>
            <div class="review-images" v-if="review.images.length > 0">
              <img v-for="img in review.images" :key="img" :src="img" alt="评价图片" class="review-image">
            </div>
          </div>
        </div>
        <div class="pagination-container">
          <el-pagination
            v-model:current-page="reviewCurrentPage"
            :page-size="reviewPageSize"
            layout="total, prev, pager, next"
            :total="filteredReviews.length"
            @current-change="handleReviewCurrentChange"
          /> 
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ChevronDown, ChevronRight, Lock, Unlock } from '@element-plus/icons-vue'
import { ElLoading, ElMessage } from 'element-plus'
// 引入courseAPI模块
import { courseAPI } from '@/api/index'

// 路由和路由参数
const route = useRoute()
const router = useRouter()
const id = route.params.id

// 状态管理
const activeTab = ref('introduction')
const expandedChapters = ref([])
const reviewFilter = ref('all')
const reviewCurrentPage = ref(1)
const reviewPageSize = ref(5)
const course = ref(null)
const loading = ref(true)
const error = ref('')

// 默认数据备份
const defaultCourseData = {
  id: id || '1',
  title: 'Vue.js从入门到精通',
  teacher: '张老师',
  price: '¥299',
  originalPrice: '¥599',
  image: 'https://picsum.photos/800/400?random=8',
  rating: 4.8,
  ratingCount: 120,
  studentCount: 1560,
  duration: '36小时',
  level: '初级到中级',
  updateTime: '2023-08-15',
  description: `<p>本课程从Vue.js的基础开始讲起，逐步深入到Vue.js的高级特性和实战应用。通过本课程的学习，你将能够:</p>
<ul>
  <li>掌握Vue.js的核心概念和基本语法</li>
  <li>理解组件化开发思想</li>
  <li>熟练使用Vue.js的各种指令和生命周期钩子</li>
  <li>掌握Vuex状态管理和Vue Router路由管理</li>
  <li>能够独立开发Vue.js应用</li>
</ul>
<p>课程内容丰富，包含大量实战案例，适合对前端开发有一定基础，想学习Vue.js的开发者。</p>`,
  chapters: [
    {
      id: 1,
      title: '第1章：Vue.js基础',
      duration: '6小时',
      videos: [
        { id: 1, title: '1.1 Vue.js介绍和环境搭建', duration: '45分钟', isFree: true },
        { id: 2, title: '1.2 Vue.js基本语法', duration: '60分钟', isFree: true },
        { id: 3, title: '1.3 Vue.js指令', duration: '75分钟', isFree: false },
        { id: 4, title: '1.4 Vue.js计算属性和监听器', duration: '60分钟', isFree: false }
      ]
    },
    {
      id: 2,
      title: '第2章：Vue.js组件化开发',
      duration: '8小时',
      videos: [
        { id: 5, title: '2.1 组件基本概念', duration: '45分钟', isFree: false },
        { id: 6, title: '2.2 组件通信', duration: '75分钟', isFree: false },
        { id: 7, title: '2.3 组件插槽', duration: '60分钟', isFree: false },
        { id: 8, title: '2.4 组件生命周期', duration: '60分钟', isFree: false }
      ]
    },
    {
      id: 3,
      title: '第3章：Vue.js高级特性',
      duration: '10小时',
      videos: [
        { id: 9, title: '3.1 Vuex状态管理', duration: '90分钟', isFree: false },
        { id: 10, title: '3.2 Vue Router路由管理', duration: '90分钟', isFree: false },
        { id: 11, title: '3.3 自定义指令', duration: '60分钟', isFree: false },
        { id: 12, title: '3.4 插件开发', duration: '60分钟', isFree: false }
      ]
    },
    {
      id: 4,
      title: '第4章：Vue.js项目实战',
      duration: '12小时',
      videos: [
        { id: 13, title: '4.1 项目初始化', duration: '60分钟', isFree: false },
        { id: 14, title: '4.2 项目组件设计', duration: '90分钟', isFree: false },
        { id: 15, title: '4.3 项目状态管理', duration: '90分钟', isFree: false },
        { id: 16, title: '4.4 项目部署上线', duration: '60分钟', isFree: false }
      ]
    }
  ],
  teacherInfo: {
    avatar: 'https://picsum.photos/100/100?random=20',
    title: '资深前端工程师 | Vue.js专家',
    introduction: `<p>张老师，10年Web前端开发经验，5年Vue.js开发和教学经验。</p>
<p>曾任职于多家知名互联网公司，负责过多个大型Vue.js项目的架构设计和开发。</p>
<p>精通Vue.js、React、Angular等前端框架，热衷于分享前端开发经验和技术。</p>`
  },
  reviews: [
    {
      id: 1,
      userName: '小明',
      userAvatar: 'https://picsum.photos/50/50?random=21',
      rating: 5,
      time: '2023-08-10',
      content: '课程非常好，老师讲解清晰，内容丰富，实战性强。学完后我成功跳槽到了一家更好的公司，薪资提升了30%。',
      images: []
    },
    {
      id: 2,
      userName: '小红',
      userAvatar: 'https://picsum.photos/50/50?random=22',
      rating: 4,
      time: '2023-08-05',
      content: '课程内容很全面，从基础到高级都有覆盖。不过有些地方讲解有点快，需要反复观看才能理解。',
      images: []
    },
    {
      id: 3,
      userName: '小刚',
      userAvatar: 'https://picsum.photos/50/50?random=23',
      rating: 5,
      time: '2023-07-28',
      content: '老师讲得非常好，实战案例很有代表性，学完后我立刻在工作中应用了所学知识，效率提升了很多。',
      images: ['https://picsum.photos/200/150?random=24', 'https://picsum.photos/200/150?random=25']
    },
    {
      id: 4,
      userName: '小丽',
      userAvatar: '/api/50/50?random=26',
      rating: 3,
      time: '2023-07-20',
      content: '课程内容还可以，但是有些地方讲解不够详细，希望老师能改进一下。',
      images: []
    },
    {
      id: 5,
      userName: '小强',
      userAvatar: '/api/50/50?random=27',
      rating: 5,
      time: '2023-07-15',
      content: '非常棒的课程，内容深入浅出，实战性强，强烈推荐给想学习Vue.js的同学。',
      images: []
    }
  ]
}

// 生命周期钩子
onMounted(() => {
  // 检查id参数是否存在
  if (!id) {
    error.value = '无法加载课程详情：缺少课程ID'
    loading.value = false
    console.error(error.value)
    return
  }

  // 使用id参数从API获取课程数据
  fetchCourseData()
})

// 获取课程数据
const fetchCourseData = async () => {
  loading.value = true
  error.value = ''

  try {
    // 使用真实API调用获取课程数据
    const response = await courseAPI.getCourseDetail(id)
    course.value = response.data

    // 模拟API请求延迟（实际应用中可移除）
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch (err) {
    error.value = '获取课程数据失败：' + (err.message || '未知错误')
    console.error(error.value)
    // 使用默认数据作为备份
    course.value = defaultCourseData
  } finally {
    loading.value = false
  }
}

// 切换章节展开/收起
const toggleChapter = (chapterId) => {
  if (expandedChapters.value.includes(chapterId)) {
    expandedChapters.value = expandedChapters.value.filter(id => id !== chapterId)
  } else {
    expandedChapters.value.push(chapterId)
  }
}

// 处理评价分页变化
const handleReviewCurrentChange = (currentPage) => {
  reviewCurrentPage.value = currentPage
}

// 购买课程
const purchaseCourse = () => {
  // 购买课程逻辑
  router.push({ path: '/payment', query: { courseId: course.value?.id } })
}

// 加入购物车
const addToCart = () => {
  // 加入购物车逻辑
  ElMessage.success('已添加到购物车')
}

// 计算筛选后的评价
const filteredReviews = computed(() => {
  if (!course.value || !course.value.reviews) return []
  let reviews = [...course.value.reviews]

  // 按评价类型筛选
  switch (reviewFilter.value) {
    case 'positive':
      reviews = reviews.filter(review => review.rating >= 4)
      break
    case 'neutral':
      reviews = reviews.filter(review => review.rating === 3)
      break
    case 'negative':
      reviews = reviews.filter(review => review.rating <= 2)
      break
  }

  return reviews
})

// 计算分页后的评价
const paginatedReviews = computed(() => {
  const startIndex = (reviewCurrentPage.value - 1) * reviewPageSize.value
  const endIndex = startIndex + reviewPageSize.value
  return filteredReviews.value.slice(startIndex, endIndex)
})
</script>

<style scoped>
.course-cover-placeholder,
.avatar-placeholder {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
  color: #999;
  font-size: 16px;
}

.course-cover-placeholder {
  width: 100%;
  height: 400px;
}

.avatar-placeholder {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}
.course-detail-container {
  padding: 20px 50px;
}

.loading-container,
.error-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 500px;
}

.error-container {
  width: 100%;
}

.course-content {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.course-header {
  display: flex;
  gap: 30px;
  margin-bottom: 30px;
  flex-wrap: wrap;
}

.course-cover {
  width: 400px;
  height: 250px;
  object-fit: cover;
}

.course-basic-info {
  flex: 1;
  min-width: 300px;
}

.course-basic-info h1 {
  font-size: 24px;
  margin-bottom: 15px;
}

.course-teacher {
  font-size: 16px;
  margin-bottom: 10px;
  color: #666;
}

.course-rating {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
}

.rating-score {
  font-size: 18px;
  font-weight: bold;
  margin-left: 10px;
  color: #f56c6c;
}

.rating-count {
  margin-left: 10px;
  color: #999;
  font-size: 14px;
}

.course-price {
  margin-bottom: 20px;
}

.current-price {
  font-size: 28px;
  font-weight: bold;
  color: #f56c6c;
}

.original-price {
  margin-left: 10px;
  color: #999;
  text-decoration: line-through;
}

.course-actions {
  display: flex;
  gap: 15px;
}

.course-tabs {
  margin-bottom: 30px;
}

.introduction-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.course-description {
  margin-bottom: 30px;
  line-height: 1.8;
}

.course-info-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
}

.info-card {
  flex: 1;
  min-width: 200px;
}

.info-item {
  display: flex;
  justify-content: space-between;
}

.info-label {
  color: #999;
}

.info-value {
  font-weight: bold;
}

.chapters-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.chapter {
  margin-bottom: 15px;
  border-bottom: 1px solid #e5e5e5;
  padding-bottom: 15px;
}

.chapter:last-child {
  border-bottom: none;
}

.chapter-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 5px;
}

.chapter-title:hover {
  background-color: #e9e9e9;
}

.chapter-duration {
  color: #999;
  font-size: 14px;
}

.chapter-videos {
  padding: 10px 20px;
}

.video-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px;
  border-bottom: 1px dashed #e5e5e5;
}

.video-item:last-child {
  border-bottom: none;
}

.video-title {
  display: flex;
  align-items: center;
  gap: 10px;
}

.video-duration {
  color: #999;
  font-size: 14px;
}

.teacher-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  display: flex;
  gap: 30px;
  flex-wrap: wrap;
}

.teacher-avatar {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  object-fit: cover;
}

.teacher-info {
  flex: 1;
  min-width: 300px;
}

.teacher-info h3 {
  font-size: 20px;
  margin-bottom: 10px;
}

.teacher-title {
  color: #666;
  margin-bottom: 15px;
}

.teacher-intro {
  line-height: 1.8;
}

.reviews-content {
  padding: 20px;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.review-filter {
  margin-bottom: 20px;
  width: 200px;
}

.review-item {
  display: flex;
  gap: 15px;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e5e5e5;
}

.review-item:last-child {
  border-bottom: none;
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
}

.review-content {
  flex: 1;
}

.review-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.user-name {
  font-weight: bold;
}

.review-time {
  color: #999;
  font-size: 14px;
}

.review-text {
  margin-bottom: 15px;
  line-height: 1.6;
}

.review-images {
  display: flex;
  gap: 10px;
}

.review-image {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 5px;
}

.pagination-container {
  margin-top: 30px;
  display: flex;
  justify-content: center;
}
</style>