<template>
  <div class="community-container">
    <div class="community-header">
      <h1>社区讨论</h1>      
      <el-button type="primary" @click="goToCreatePost">发布新帖子</el-button>
      <el-button @click="fetchPosts">刷新帖子</el-button>
    </div>
    <div class="community-content">
      <template v-if="loading">
        <div class="loading-container">
          <el-skeleton active :count="3" /> 
        </div>
      </template>
      <template v-else>
        <div v-if="posts && posts.length > 0">
          <div class="post-list">
            <el-card v-for="post in posts" :key="post.id" class="post-card">
              <template #header>
                <div class="post-header">
                  <img :src="post.authorAvatar" alt="作者头像" class="author-avatar">
                  <div class="author-info">
                    <h3 class="author-name">{{ post.authorName }}</h3>
                    <p class="post-time">{{ formatDate(post.createdAt) }}</p>
                  </div>
                </div>
              </template>
              <div class="post-body">
                <h2 class="post-title" @click="goToPostDetail(post.id)">{{ post.title }}</h2>
                <p class="post-excerpt">{{ post.excerpt }}</p>
              </div>
              <div class="post-footer">
                <div class="post-meta">
                  <span class="post-category">{{ post.category }}</span>
                  <span class="post-comment-count"><i class="el-icon-message"></i> {{ post.commentCount }}</span>
                  <span class="post-view-count"><i class="el-icon-eye"></i> {{ post.viewCount }}</span>
                </div>
              </div>
            </el-card>
          </div>
          <!-- 调试信息区域，通过showDebugInfo属性控制显示 -->
          <div v-if="showDebugInfo" class="debug-info">
          共 {{ posts.length }} 条帖子
          <div v-for="(post, index) in posts" :key="'debug-' + index" class="debug-post-item">
            {{ index + 1 }}. {{ post.title }} (ID: {{ post.id }})
          </div>
        </div>
        </div>
        <div v-else class="no-data">暂无帖子数据</div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
// 控制调试信息显示的变量，默认为false不显示
defineProps({
  showDebugInfo: {
    type: Boolean,
    default: false
  }
});
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router';
import { postAPI } from '@/api/index';
import { ElMessage } from 'element-plus';

const router = useRouter();
const route = useRoute();
const posts = ref([]);
const loading = ref(true);
let routeListener = null;

// 获取帖子列表
const fetchPosts = async () => {
  console.log('开始获取帖子列表');
  try {
    loading.value = true;
    // 检查是否有token
    const token = localStorage.getItem('token');
    console.log('当前token:', token);
    console.log('准备调用API获取帖子列表');
    // 添加默认参数，确保与后端API期望的参数格式匹配
    const data = await postAPI.getPostList({ page: 1, pageSize: 10 });
    // 打印响应数据，以便调试
    console.log('帖子列表API响应数据:', JSON.stringify(data));
    console.log('API返回的数据类型:', typeof data);
    console.log('API返回的code:', data.code);
    console.log('API返回的数据总数:', data.data?.total);
    console.log('API返回的列表长度:', data.data?.list?.length);
    // 详细记录API返回的原始数据
    console.log('API返回的原始数据:', JSON.stringify(data));
    // 简化响应处理，直接使用后端返回的格式
    let postList = [];
    if (data && data.code === 200 && data.data && Array.isArray(data.data.list)) {
      // 匹配后端返回格式: {code: 200, data: {list: [...]}, msg: "success"}
      postList = data.data.list;
      console.log('提取到帖子数据:', postList);
      console.log('帖子数据长度:', postList.length);
      // 详细记录每条原始帖子数据的分类
      postList.forEach((post, index) => {
        console.log(`原始帖子 ${index + 1} (ID: ${post.id}) 分类: ${post.category}`);
      });
    } else {
      console.error('API响应格式不符合预期:', data);
    }
    // 确保posts.value始终是数组
    console.log('赋值前posts.value:', posts.value);
    console.log('准备赋值的postList:', postList);
    posts.value = postList;
    console.log('赋值后posts.value:', posts.value);
    console.log('赋值后posts.value长度:', posts.value.length);
    // 添加去除HTML标签的辅助函数
    const stripHtml = (html) => {
      const div = document.createElement('div');
      div.innerHTML = html;
      return div.textContent || '';
    };

    // 为帖子添加头像和处理摘要，并映射字段名
    posts.value = posts.value.map(post => ({
      ...post,
      authorAvatar: post.avatar || `/photos/100/100?random=${post.id || Math.random()}`,
      authorName: post.user_name || post.username || '匿名用户',
      excerpt: post.content ? stripHtml(post.content).substring(0, 100) + '...' : '暂无内容',
      createdAt: post.created_at || post.createdAt || new Date().toISOString(),
      category: post.category || '未分类',
      commentCount: post.comment_count || post.commentCount || 0,
      viewCount: post.view_count || post.viewCount || 0
    }));
    console.log('映射后的帖子数据:', posts.value);
    // 日志记录成功信息
    console.log('获取帖子列表成功，共', posts.value.length, '条帖子');
    // 详细记录每条帖子的分类信息
    posts.value.forEach((post, index) => {
      console.log(`帖子 ${index + 1} (ID: ${post.id}) 分类: ${post.category}`);
    });
    console.log('处理后的帖子数据:', posts.value);
  } catch (error) {
    console.error('获取帖子列表失败:', error);
    console.error('错误响应数据:', error.response?.data);
    // 提供更具体的错误信息
    const errorMsg = error.response?.data?.msg || error.response?.data?.message || error.message || '未知错误';
    ElMessage.error(`获取帖子列表失败: ${errorMsg}`);
    // 提供默认数据以防API失败
    posts.value = [
      {
        id: 1,
        title: '如何有效地学习编程？',
        excerpt: '分享一些编程学习的经验和技巧，帮助新手快速入门...',
        authorName: '张教授',
        authorAvatar: '/photos/100/100?random=1',
        createdAt: '2023-11-15T10:30:00',
        category: '学习方法',
        commentCount: 24,
        viewCount: 320
      },
      {
        id: 2,
        title: '前端开发的未来趋势是什么？',
        excerpt: '讨论当前前端技术的发展方向和未来可能的趋势...',
        authorName: '李工程师',
        authorAvatar: '/photos/100/100?random=2',
        createdAt: '2023-11-14T15:45:00',
        category: '技术讨论',
        commentCount: 18,
        viewCount: 276
      },
      {
        id: 3,
        title: '如何准备技术面试？',
        excerpt: '分享技术面试的准备经验和常见问题解答...',
        authorName: '王经理',
        authorAvatar: '/photos/100/100?random=3',
        createdAt: '2023-11-12T09:15:00',
        category: '求职经验',
        commentCount: 32,
        viewCount: 412
      }
    ];
  } finally {
    loading.value = false;
  }
};

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString();
};

// 导航到帖子详情
const goToPostDetail = (postId) => {
  router.push({ name: 'postDetail', params: { id: postId } });
};

// 导航到创建帖子页面
const goToCreatePost = () => {
  router.push({ name: 'createPost' });
};

// 监听路由变化，确保每次进入页面都刷新数据
const setupRouteListener = () => {
  routeListener = route.meta.refresh ? route.meta.refresh : router.afterEach((to) => {
    if (to.name === 'community') {
      fetchPosts();
    }
  });
};

onMounted(() => {
  fetchPosts();
  setupRouteListener();
});

onBeforeUnmount(() => {
  if (routeListener && typeof routeListener === 'function') {
    routeListener();
  }
});
</script>

<style scoped>
.community-container {
  padding: 20px;
}

.community-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.post-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.post-card {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.post-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.author-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.debug-info {
  margin-top: 20px;
  padding: 10px;
  background-color: #f0f0f0;
  border-radius: 4px;
  text-align: left;
}
.debug-post-item {
  margin-top: 5px;
  padding: 5px;
  background-color: #e0e0e0;
  border-radius: 3px;
}

.no-data {
  text-align: center;
  padding: 50px 0;
  color: #999;
}

.author-info {
  display: flex;
  flex-direction: column;
}

.author-name {
  margin: 0;
  font-size: 16px;
  font-weight: bold;
}

.post-time {
  margin: 0;
  font-size: 12px;
  color: #999;
}

.post-title {
  margin: 10px 0;
  font-size: 18px;
  cursor: pointer;
}

.post-title:hover {
  color: #409eff;
}

.post-excerpt {
  margin: 0 0 10px 0;
  color: #666;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-footer {
  margin-top: auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.post-meta {
  display: flex;
  gap: 15px;
}

.post-category {
  background-color: #f0f9eb;
  color: #67c23a;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.post-comment-count, .post-view-count {
  font-size: 12px;
  color: #999;
}

.author-info {
  flex: 1;
}

.author-name {
  font-size: 16px;
  font-weight: bold;
  margin: 0;
}

.post-time {
  font-size: 12px;
  color: #999;
  margin: 0;
}

.post-title {
  font-size: 18px;
  margin: 10px 0;
  cursor: pointer;
  color: #333;
}

.post-title:hover {
  color: #409eff;
}

.post-excerpt {
  color: #666;
  margin-bottom: 15px;
  flex: 1;
}

.post-footer {
  margin-top: auto;
}

.post-meta {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
  color: #999;
}

.post-category {
  background-color: #f0f9eb;
  color: #67c23a;
  padding: 2px 8px;
  border-radius: 4px;
}

/* 加载状态样式 */
.loading-container {
  display: flex;
  justify-content: center;
  padding: 30px 0;
}

/* 无数据状态样式 */
.no-data {
  text-align: center;
  padding: 50px 0;
  color: #999;
  font-size: 16px;
  background-color: #f9f9f9;
  border-radius: 8px;
}
</style>