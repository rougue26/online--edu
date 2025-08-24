<template>
  <div class="admin-home-container">
    <el-card class="welcome-card">
      <div class="welcome-content">
        <h1>欢迎使用在线教育系统后台管理</h1>
        <p>这里是系统管理的核心区域，您可以管理课程、用户、视频和社区内容。</p>
      </div>
    </el-card>

    <div class="stats-container">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card class="stat-card" shadow="hover">
            <div class="stat-content">
              <div class="stat-title">总课程数</div>
              <div class="stat-value">{{ courseCount }}</div>
              <div class="stat-change">+12% 较上月</div>
            </div>
            <div class="stat-icon"><Cpu /></div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card" shadow="hover">
            <div class="stat-content">
              <div class="stat-title">总用户数</div>
              <div class="stat-value">{{ userCount }}</div>
              <div class="stat-change">+8% 较上月</div>
            </div>
            <div class="stat-icon"><User /></div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card" shadow="hover">
            <div class="stat-content">
              <div class="stat-title">总视频数</div>
              <div class="stat-value">{{ videoCount }}</div>
              <div class="stat-change">+15% 较上月</div>
            </div>
            <div class="stat-icon"><VideoCamera /></div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="stat-card" shadow="hover">
            <div class="stat-content">
              <div class="stat-title">总帖子数</div>
              <div class="stat-value">{{ postCount }}</div>
              <div class="stat-change">+20% 较上月</div>
            </div>
            <div class="stat-icon"><Message /></div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="charts-container">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-card class="chart-card">
            <template #header>
              <div class="card-header">
                <span>课程分类统计</span>
              </div>
            </template>
            <div class="chart-content">
              <!-- 这里将放置图表组件 -->
              <div class="placeholder">图表区域</div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card class="chart-card">
            <template #header>
              <div class="card-header">
                <span>用户增长趋势</span>
              </div>
            </template>
            <div class="chart-content">
              <!-- 这里将放置图表组件 -->
              <div class="placeholder">图表区域</div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { courseAPI, userAPI, videoAPI, postAPI } from '../api/index';
import { Cpu, User, VideoCamera, Message } from '@element-plus/icons-vue';

export default {
  name: 'AdminHomeView',
  components: {
    Cpu,
    User,
    VideoCamera,
    Message
  },
  setup() {
    const courseCount = ref(0);
    const userCount = ref(0);
    const videoCount = ref(0);
    const postCount = ref(0);
    const loading = ref(true);

    // 获取统计数据
    const fetchStats = async () => {
      try {
        loading.value = true;
        // 获取课程数量
        const coursesRes = await courseAPI.getCourseList({ page: 1, pageSize: 1 });
        courseCount.value = coursesRes.total || 0;

        // 获取用户数量 - 暂时使用模拟数据，因为API未实现
        userCount.value = 100;

        // 获取视频数量
        const videosRes = await videoAPI.getVideoList({ page: 1, pageSize: 1 });
        videoCount.value = videosRes.total || 0;

        // 获取帖子数量
        const postsRes = await postAPI.getPostList({ page: 1, pageSize: 1 });
        postCount.value = postsRes.total || 0;
      } catch (error) {
        console.error('获取统计数据失败:', error);
      } finally {
        loading.value = false;
      }
    };

    onMounted(() => {
      fetchStats();
    });

    return {
      courseCount,
      userCount,
      videoCount,
      postCount,
      loading
    };
  }
};
</script>

<style scoped>
.admin-home-container {
  padding: 20px;
}

.welcome-card {
  margin-bottom: 20px;
}

.welcome-content {
  padding: 30px;
  text-align: center;
}

.stats-container {
  margin-bottom: 20px;
}

.stat-card {
  height: 100%;
  display: flex;
  padding: 20px;
}

.stat-content {
  flex: 1;
}

.stat-title {
  font-size: 14px;
  color: #909399;
  margin-bottom: 10px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 5px;
}

.stat-change {
  font-size: 12px;
  color: #67c23a;
}

.stat-icon {
  font-size: 36px;
  color: #409eff;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 60px;
}

.charts-container {
  margin-bottom: 20px;
}

.chart-card {
  height: 300px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chart-content {
  height: calc(100% - 48px);
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder {
  color: #909399;
  font-size: 14px;
}
</style>