<template>
  <div class="video-list-container">
    <el-card class="page-header">
      <div class="header-content">
        <h1>视频中心</h1>
        <el-button type="primary" icon="el-icon-plus" @click="goToUpload">上传视频</el-button>
      </div>
    </el-card>

    <el-card class="filter-section">
      <el-form :model="filterForm" inline>
        <el-form-item label="分类">
          <el-select v-model="filterForm.category" placeholder="全部">
            <el-option label="全部" value=""></el-option>
            <el-option label="编程开发" value="programming"></el-option>
            <el-option label="设计创意" value="design"></el-option>
            <el-option label="职业技能" value="skills"></el-option>
            <el-option label="兴趣爱好" value="hobbies"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序">
          <el-select v-model="filterForm.sort" placeholder="最新">
            <el-option label="最新" value="newest"></el-option>
            <el-option label="最热" value="hottest"></el-option>
            <el-option label="最多播放" value="most_viewed"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="fetchVideos">筛选</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <div class="video-grid">
      <template v-if="loading">
        <el-card v-for="i in 8" :key="i" class="video-card skeleton-card">
          <div class="video-thumbnail-container skeleton-thumbnail"></div>
          <div class="video-info">
            <h3 class="video-title skeleton-line"></h3>
            <div class="video-meta">
              <span class="video-author skeleton-line"></span>
              <span class="video-date skeleton-line"></span>
              <span class="video-views skeleton-line"></span>
            </div>
          </div>
        </el-card>
      </template>

      <template v-else-if="videos.length > 0">
        <el-card v-for="video in videos" :key="video.id" class="video-card" @click="goToDetail(video.id)">
          <div class="video-thumbnail-container">
            <img :src="video.coverImageUrl" alt="视频封面" class="video-thumbnail">
            <div class="video-duration">{{ formatDuration(video.duration) }}</div>
          </div>
          <div class="video-info">
            <h3 class="video-title">{{ video.title }}</h3>
            <div class="video-meta">
              <span class="video-author">{{ video.authorName }}</span>
              <span class="video-date">{{ formatDate(video.createdAt) }}</span>
              <span class="video-views"><i class="el-icon-eye"></i> {{ video.viewCount }}</span>
            </div>
          </div>
        </el-card>
      </template>

      <div v-else class="no-data">
        <el-empty description="暂无视频数据"></el-empty>
      </div>
    </div>

    <div class="pagination-container">
      <el-pagination
        v-model:current-page="currentPage"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalVideos"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { videoAPI } from '@/api/index'

// 状态定义
const videos = ref([])
const filterForm = ref({
  category: '',
  sort: 'newest'
})
const currentPage = ref(1)
const pageSize = ref(10)
const totalVideos = ref(0)
const loading = ref(false)
const error = ref('')
const router = useRouter()

// 默认视频数据（API失败时使用）
const defaultVideos = [
  {
    id: '1',
    title: 'Vue.js基础教程 - 组件化开发',
    coverImageUrl: 'https://picsum.photos/800/450?random=10',
    duration: 360,
    authorName: '张教授',
    createdAt: '2023-11-20T14:30:00',
    viewCount: 1250
  },
  {
    id: '2',
    title: 'React Hooks实战指南',
    coverImageUrl: 'https://picsum.photos/800/450?random=11',
    duration: 420,
    authorName: '李工程师',
    createdAt: '2023-11-19T09:15:00',
    viewCount: 980
  },
  {
    id: '3',
    title: 'Python数据分析入门',
    coverImageUrl: 'https://picsum.photos/800/450?random=12',
    duration: 540,
    authorName: '王老师',
    createdAt: '2023-11-18T16:45:00',
    viewCount: 1560
  },
  {
    id: '4',
    title: 'UI设计原理与实践',
    coverImageUrl: 'https://picsum.photos/800/450?random=13',
    duration: 480,
    authorName: '陈设计师',
    createdAt: '2023-11-17T10:20:00',
    viewCount: 870
  }
]

// 获取视频列表
const fetchVideos = async () => {
  loading.value = true
  error.value = ''

  try {
    const response = await videoAPI.getVideoList({
      categoryID: filterForm.value.category,
      sort: filterForm.value.sort,
      page: currentPage.value,
      pageSize: pageSize.value
    })

    // 确保response.data是对象
    const data = response.data || {};
    // 确保items始终是数组
    videos.value = Array.isArray(data.items) ? data.items : [];
    totalVideos.value = data.total || 0;
  } catch (err) {
    error.value = '获取视频列表失败'
    ElMessage.error(error.value + ': ' + (err.message || ''))
    // 使用默认数据
    videos.value = defaultVideos
    totalVideos.value = 24
  } finally {
    loading.value = false
  }
}

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString()
}

// 格式化时长
const formatDuration = (seconds) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`
}

// 跳转到视频详情页
const goToDetail = (videoId) => {
  router.push({ name: 'videoDetail', params: { id: videoId } })
}

// 跳转到上传页面
const goToUpload = () => {
  router.push({ name: 'videoUpload' })
}

// 处理分页大小变化
const handleSizeChange = (size) => {
  pageSize.value = size
  fetchVideos()
}

// 处理当前页变化
const handleCurrentChange = (current) => {
  currentPage.value = current
  fetchVideos()
}

// 组件挂载时获取数据
onMounted(() => {
  fetchVideos()
})
</script>

<style scoped>
.video-list-container {
  padding: 20px;
}

.page-header {
  margin-bottom: 20px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.filter-section {
  margin-bottom: 20px;
}

.video-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.video-card {
  cursor: pointer;
  transition: transform 0.2s;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.video-card:hover {
  transform: translateY(-5px);
}

.video-thumbnail-container {
  position: relative;
  width: 100%;
  padding-top: 56.25%; /* 16:9 宽高比 */
  overflow: hidden;
  margin-bottom: 15px;
}

.video-thumbnail {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.video-card:hover .video-thumbnail {
  transform: scale(1.05);
}

.video-duration {
  position: absolute;
  bottom: 10px;
  right: 10px;
  background-color: rgba(0, 0, 0, 0.7);
  color: #fff;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 14px;
}

.video-info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.video-title {
  font-size: 16px;
  margin-bottom: 10px;
  color: #333;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.video-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  color: #666;
  font-size: 14px;
  margin-top: auto;
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 30px;
}

/* 骨架屏样式 */
.skeleton-card {
  cursor: default;
}

.skeleton-thumbnail {
  background-color: #f0f0f0;
  animation: pulse 1.5s infinite;
}

.skeleton-line {
  height: 16px;
  background-color: #f0f0f0;
  margin-bottom: 8px;
  border-radius: 4px;
  animation: pulse 1.5s infinite;
}

.video-title.skeleton-line {
  width: 80%;
  height: 20px;
}

.video-author.skeleton-line {
  width: 60%;
}

.video-date.skeleton-line {
  width: 50%;
}

.video-views.skeleton-line {
  width: 40%;
}

@keyframes pulse {
  0%, 100% {
    opacity: 0.5;
  }
  50% {
    opacity: 1;
  }
}

/* 无数据样式 */
.no-data {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 300px;
  width: 100%;
}
</style>