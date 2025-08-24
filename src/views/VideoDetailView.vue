<template>
  <div class="video-detail-container">
    <el-card class="video-card">
      <div class="video-player-container">
        <video
          ref="videoPlayer"
          class="video-player"
          controls
          :poster="videoInfo.coverImageUrl"
        >
          <source :src="videoInfo.videoUrl" type="video/mp4">
          您的浏览器不支持HTML5视频播放。
        </video>
      </div>
      <div class="video-info">
        <h1 class="video-title">{{ videoInfo.title }}</h1>
        <div class="video-meta">
          <span class="video-author" @click="goToUserProfile">{{ videoInfo.authorName }}</span>
          <span class="video-category">{{ videoInfo.category }}</span>
          <span class="video-date">{{ formatDate(videoInfo.createdAt) }}</span>
          <span class="video-views"><i class="el-icon-eye"></i> {{ videoInfo.viewCount }} 次观看</span>
        </div>
        <div class="video-description">
          <p v-html="videoInfo.description"></p>
        </div>
        <div class="video-actions">
          <el-button type="primary" icon="el-icon-thumbs-up" :disabled="isLiked" @click="likeVideo">{{ likeCount }} {{ isLiked ? '已点赞' : '点赞' }}</el-button>
          <el-button icon="el-icon-collection" :disabled="isFavorite" @click="favoriteVideo">{{ isFavorite ? '已收藏' : '收藏' }}</el-button>
          <el-button icon="el-icon-share">分享</el-button>
          <el-button v-if="isAuthor" type="danger" @click="deleteVideo">删除视频</el-button>
        </div>
      </div>
    </el-card>

    <!-- 评论区域 -->
    <div class="comments-section">
      <h2>评论 ({{ comments.length }})</h2>
      <el-form ref="commentForm" :model="commentForm" class="comment-form">
        <el-form-item prop="content">
          <el-input
            v-model="commentForm.content"
            type="textarea"
            placeholder="写下你的评论..."
            :rows="3"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitComment">发表评论</el-button>
        </el-form-item>
      </el-form>

      <div class="comment-list">
        <el-card v-for="comment in comments" :key="comment.id" class="comment-card">
          <div class="comment-header">
            <img :src="comment.userAvatar" alt="用户头像" class="user-avatar">
            <div class="user-info">
              <h4 class="user-name">{{ comment.userName }}</h4>
              <p class="comment-time">{{ formatDate(comment.createdAt) }}</p>
            </div>
          </div>
          <div class="comment-content">{{ comment.content }}</div>
          <div class="comment-actions">
            <span class="like-action" @click="likeComment(comment.id)">
              <i class="el-icon-thumbs-up"></i> {{ comment.likeCount }} {{ comment.isLiked ? '已点赞' : '点赞' }}
            </span>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElConfirm } from 'element-plus';
import { videoAPI, commentAPI } from '@/api/index';

// 路由和导航
const router = useRouter();
const route = useRoute();
const id = route.params.id;

// 状态管理
const videoInfo = reactive({
  id: '',
  title: '',
  description: '',
  videoUrl: '',
  coverImageUrl: '',
  authorId: '',
  authorName: '',
  authorAvatar: '',
  category: '',
  viewCount: 0,
  likeCount: 0,
  favoriteCount: 0,
  isPublic: true,
  createdAt: ''
});

const isLiked = ref(false);
const isFavorite = ref(false);
const isAuthor = ref(false);
const likeCount = ref(0);
const comments = ref([]);
const commentForm = reactive({
  content: ''
});
const currentUser = reactive({
  id: localStorage.getItem('userId') || '',
  avatar: localStorage.getItem('userAvatar') || ''
});
const loading = ref(true);
const error = ref('');

// 生命周期钩子
onMounted(() => {
  // 在实际应用中，这里应该根据id从API获取视频详情
  fetchVideoDetail();
});

// 获取视频详情
const fetchVideoDetail = async () => {
  loading.value = true;
  error.value = '';
  try {
    // 实际应用中，使用API调用
    const response = await videoAPI.getVideoDetail(id);
    Object.assign(videoInfo, response.data);
    likeCount.value = videoInfo.likeCount;
    isAuthor.value = currentUser.id === videoInfo.authorId;
    // 检查用户是否已点赞和收藏
    checkUserInteraction();
    fetchComments();
    // 更新观看次数
    updateViewCount();
  } catch (err) {
    error.value = '获取视频详情失败: ' + (err.message || '未知错误');
    ElMessage.error(error.value);
    // 模拟数据，防止页面空白
    if (process.env.NODE_ENV !== 'production') {
      mockVideoData();
    }
  } finally {
    loading.value = false;
  }
};

// 检查用户是否已点赞和收藏
const checkUserInteraction = async () => {
  try {
    // 实际应用中，使用API调用
    const likeResponse = await videoAPI.checkVideoLike(id);
    isLiked.value = likeResponse.data.isLiked;
    const favoriteResponse = await videoAPI.checkVideoFavorite(id);
    isFavorite.value = favoriteResponse.data.isFavorite;
  } catch (err) {
    console.error('检查用户交互状态失败:', err);
  }
};

// 获取评论
const fetchComments = async () => {
  try {
    // 实际应用中，使用API调用
    const response = await commentAPI.getVideoComments(id);
    comments.value = response.data;
  } catch (err) {
    ElMessage.error('获取评论失败: ' + (err.message || '未知错误'));
    // 模拟数据，防止页面空白
    if (process.env.NODE_ENV !== 'production') {
      mockCommentsData();
    }
  }
};

// 更新观看次数
const updateViewCount = async () => {
  try {
    // 实际应用中，使用API调用
    await videoAPI.updateVideoViewCount(id);
  } catch (err) {
    console.error('更新观看次数失败:', err);
  }
};

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString();
};

// 点赞视频
const likeVideo = async () => {
  try {
    // 实际应用中，使用API调用
    await videoAPI.likeVideo(id, !isLiked.value);
    if (isLiked.value) {
      likeCount.value--;
    } else {
      likeCount.value++;
    }
    isLiked.value = !isLiked.value;
  } catch (err) {
    ElMessage.error('点赞失败: ' + (err.message || '未知错误'));
  }
};

// 收藏视频
const favoriteVideo = async () => {
  try {
    // 实际应用中，使用API调用
    await videoAPI.favoriteVideo(id, !isFavorite.value);
    isFavorite.value = !isFavorite.value;
    ElMessage.success(isFavorite.value ? '收藏成功' : '取消收藏成功');
  } catch (err) {
    ElMessage.error('收藏操作失败: ' + (err.message || '未知错误'));
  }
};

// 点赞评论
const likeComment = async (commentId) => {
  const comment = comments.value.find(c => c.id === commentId);
  if (comment) {
    try {
      // 实际应用中，使用API调用
      await commentAPI.likeComment(commentId, !comment.isLiked);
      if (comment.isLiked) {
        comment.likeCount--;
      } else {
        comment.likeCount++;
      }
      comment.isLiked = !comment.isLiked;
    } catch (err) {
      ElMessage.error('点赞评论失败: ' + (err.message || '未知错误'));
    }
  }
};

// 提交评论
const submitComment = async () => {
  if (!commentForm.content.trim()) {
    ElMessage.warning('评论内容不能为空');
    return;
  }

  try {
    // 实际应用中，使用API调用
    const response = await commentAPI.addVideoComment(id, {
      content: commentForm.content
    });
    const newComment = response.data;
    comments.value.unshift(newComment);
    commentForm.content = '';
    ElMessage.success('评论发表成功');
  } catch (err) {
    ElMessage.error('发表评论失败: ' + (err.message || '未知错误'));
  }
};

// 跳转到用户主页
const goToUserProfile = () => {
  router.push({ name: 'userProfile', params: { id: videoInfo.authorId } });
};

// 删除视频
const deleteVideo = () => {
  ElConfirm('确定要删除该视频吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      // 实际应用中，使用API调用
      await videoAPI.deleteVideo(id);
      ElMessage.success('视频删除成功');
      router.push({ name: 'videoList' });
    } catch (err) {
      ElMessage.error('删除视频失败: ' + (err.message || '未知错误'));
    }
  }).catch(() => {
    // 取消删除
  });
};

// 模拟视频数据（开发环境）
const mockVideoData = () => {
  Object.assign(videoInfo, {
    id: id,
    title: 'Vue.js基础教程 - 组件化开发',
    description: '<p>本视频介绍了Vue.js中的组件化开发概念和实践方法，包括组件的创建、注册、通信等核心知识点。</p><p>通过本视频学习，你将能够掌握Vue.js组件化开发的精髓，提高前端开发效率。</p>',
    videoUrl: '/videos/video123/mp4/720/big_buck_bunny_720p_1mb.mp4',
    coverImageUrl: 'https://picsum.photos/800/450?random=10',
    authorId: '1',
    authorName: '张教授',
    authorAvatar: 'https://picsum.photos/100/100?random=1',
    category: '编程开发',
    viewCount: 1250,
    likeCount: 89,
    favoriteCount: 42,
    isPublic: true,
    createdAt: '2023-11-20T14:30:00'
  });
  likeCount.value = videoInfo.likeCount;
  isAuthor.value = currentUser.id === videoInfo.authorId;
};

// 模拟评论数据（开发环境）
const mockCommentsData = () => {
  comments.value = [
    {
      id: 1,
      userName: '小李',
      userAvatar: 'https://picsum.photos/100/100?random=4',
      content: '讲解非常清晰，受益良多！',
      createdAt: '2023-11-20T15:20:00',
      likeCount: 12,
      isLiked: false
    },
    {
      id: 2,
      userName: '小王',
      userAvatar: 'https://picsum.photos/100/100?random=5',
      content: '请问老师，组件通信还有其他方式吗？',
      createdAt: '2023-11-20T16:45:00',
      likeCount: 5,
      isLiked: false
    }
  ];
};
</script>

<style scoped>
.video-detail-container {
  padding: 20px;
}

.video-card {
  margin-bottom: 20px;
}

.video-player-container {
  position: relative;
  width: 100%;
  padding-top: 56.25%; /* 16:9 宽高比 */
  margin-bottom: 20px;
}

.video-player {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #000;
}

.video-info {
  padding: 0 20px 20px;
}

.video-title {
  font-size: 24px;
  margin-bottom: 15px;
  color: #333;
}

.video-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  margin-bottom: 15px;
  color: #666;
  font-size: 14px;
}

.video-author {
  cursor: pointer;
  color: #409eff;
}

.video-author:hover {
  color: #409eff;
}

.video-category {
  background-color: #f0f9eb;
  color: #67c23a;
  padding: 2px 8px;
  border-radius: 4px;
}

.video-description {
  margin-bottom: 20px;
  line-height: 1.6;
}

.video-actions {
  display: flex;
  gap: 10px;
}

.comments-section {
  margin-top: 30px;
}

.comment-form {
  margin-bottom: 20px;
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.comment-card {
  padding: 15px;
}

.comment-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.user-info {
  flex: 1;
}

.user-name {
  font-weight: bold;
  margin-bottom: 3px;
}

.comment-time {
  font-size: 12px;
  color: #999;
}

.comment-content {
  margin-bottom: 10px;
  line-height: 1.6;
}

.comment-actions {
  color: #666;
  font-size: 14px;
}

.like-action {
  cursor: pointer;
}

.like-action:hover {
  color: #409eff;
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
  margin-bottom: 10px;
  border-radius: 4px;
  animation: pulse 1.5s infinite;
}

.video-title.skeleton-line {
  width: 60%;
  height: 28px;
}

.video-author.skeleton-line {
  width: 30%;
}

.video-category.skeleton-line {
  width: 25%;
}

.video-date.skeleton-line {
  width: 35%;
}

.video-views.skeleton-line {
  width: 40%;
}

.video-description .skeleton-line {
  width: 90%;
  height: 18px;
  margin-bottom: 8px;
}

.video-description .skeleton-line:nth-child(2) {
  width: 70%;
}

.video-actions.skeleton-line {
  width: 50%;
  height: 36px;
}

.comments-section .skeleton-line:nth-child(1) {
  width: 20%;
  height: 24px;
}

.comment-form.skeleton-line {
  width: 100%;
  height: 120px;
}

.skeleton-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f0f0f0;
  margin-right: 10px;
  animation: pulse 1.5s infinite;
}

.user-name.skeleton-line {
  width: 40%;
  height: 18px;
}

.comment-time.skeleton-line {
  width: 30%;
  height: 14px;
  margin-bottom: 5px;
}

.comment-content.skeleton-line {
  width: 90%;
  height: 16px;
}

.comment-actions.skeleton-line {
  width: 30%;
  height: 16px;
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
  height: 200px;
  width: 100%;
}

/* 错误容器样式 */
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 0;
  gap: 20px;
}

@media (max-width: 768px) {
  .video-meta {
    flex-direction: column;
    gap: 5px;
  }

  .video-actions {
    flex-wrap: wrap;
  }
}
</style>