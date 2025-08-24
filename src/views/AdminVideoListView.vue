<template>
  <div class="admin-video-list-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>视频列表</span>
          <el-button type="primary" size="small" @click="handleUploadVideo">上传视频</el-button>
        </div>
      </template>
      <el-table :data="videos" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column label="封面" width="120">
          <template #default="scope">
            <img :src="scope.row.coverImageUrl" alt="封面" class="cover-image">
          </template>
        </el-table-column>
        <el-table-column prop="title" label="视频标题"></el-table-column>
        <el-table-column prop="authorName" label="作者" width="120"></el-table-column>
        <el-table-column prop="duration" label="时长" width="100">
          <template #default="scope">
            {{ formatDuration(scope.row.duration) }}
          </template>
        </el-table-column>
        <el-table-column prop="viewCount" label="观看次数" width="100"></el-table-column>
        <el-table-column prop="isPublic" label="是否公开" width="100">
          <template #default="scope">
            <el-switch v-model="scope.row.isPublic" @change="handleTogglePublic(scope.row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEditVideo(scope.row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDeleteVideo(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { videoAPI } from '../api/index';

export default {
  name: 'AdminVideoListView',
  setup() {
    const videos = ref([]);
    const currentPage = ref(1);
    const pageSize = ref(10);
    const total = ref(0);
    const loading = ref(false);

    // 获取视频列表
    const fetchVideos = async () => {
      try {
        loading.value = true;
        const res = await videoAPI.getVideoList({
          page: currentPage.value,
          pageSize: pageSize.value
        });
        videos.value = res.items || [];
        total.value = res.total || 0;
      } catch (error) {
        console.error('获取视频列表失败:', error);
      } finally {
        loading.value = false;
      }
    };

    // 格式化时长
    const formatDuration = (seconds) => {
      const minutes = Math.floor(seconds / 60);
      const remainingSeconds = seconds % 60;
      return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`;
    };

    // 处理分页大小变化
    const handleSizeChange = (size) => {
      pageSize.value = size;
      fetchVideos();
    };

    // 处理当前页变化
    const handleCurrentChange = (current) => {
      currentPage.value = current;
      fetchVideos();
    };

    const router = useRouter();

    // 添加视频管理相关路由
    router.addRoute('adminDashboard', {
      path: 'videos/create',
      name: 'adminCreateVideo',
      component: () => import('../views/AdminVideoUploadView.vue'),
      meta: { requiresAuth: true, requiresAdmin: true }
    });

    router.addRoute('adminDashboard', {
      path: 'videos/:id/edit',
      name: 'adminEditVideo',
      component: () => import('../views/AdminVideoUploadView.vue'),
      meta: { requiresAuth: true, requiresAdmin: true }
    });

    // 处理上传视频
    const handleUploadVideo = () => {
      // 跳转到上传视频页面
      router.push('/admin/dashboard/videos/create');
    };

    // 处理编辑视频
    const handleEditVideo = (video) => {
      // 跳转到编辑视频页面
      router.push(`/admin/dashboard/videos/${video.id}/edit`);
    };

    // 处理删除视频
    const handleDeleteVideo = (id) => {
      alert(`删除视频 ID: ${id} 功能待实现`);
    };

    // 处理切换公开状态
    const handleTogglePublic = (video) => {
      alert(`切换视频 ${video.title} 公开状态功能待实现`);
    };

    onMounted(() => {
      fetchVideos();
    });

    return {
      videos,
      currentPage,
      pageSize,
      total,
      loading,
      formatDuration,
      handleSizeChange,
      handleCurrentChange,
      handleUploadVideo,
      handleEditVideo,
      handleDeleteVideo,
      handleTogglePublic
    };
  }
};
</script>

<style scoped>
.admin-video-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.cover-image {
  width: 80px;
  height: 45px;
  object-fit: cover;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>