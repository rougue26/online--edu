<template>
  <div class="admin-post-list-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>帖子管理</span>
        </div>
      </template>
      <el-table :data="posts" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="title" label="帖子标题"></el-table-column>
        <el-table-column prop="authorName" label="作者" width="120"></el-table-column>
        <el-table-column prop="viewCount" label="浏览次数" width="100"></el-table-column>
        <el-table-column prop="commentCount" label="评论数" width="100"></el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="180"></el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'published' ? 'success' : 'info'">
              {{ scope.row.status === 'published' ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleEditPost(scope.row)">编辑</el-button>
            <el-button type="danger" size="small" @click="handleDeletePost(scope.row.id)">删除</el-button>
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
import { postAPI } from '../api/index';

export default {
  name: 'AdminPostListView',
  setup() {
    const posts = ref([]);
    const currentPage = ref(1);
    const pageSize = ref(10);
    const total = ref(0);
    const loading = ref(false);

    // 获取帖子列表
    const fetchPosts = async () => {
      try {
        loading.value = true;
        const res = await postAPI.getPostList({
          page: currentPage.value,
          pageSize: pageSize.value
        });
        // 数据验证和处理
        console.log('帖子列表API响应:', res);
        if (res && typeof res === 'object') {
          // 确保posts.value始终是数组
          // 注意：后端返回的是res.data.list而不是res.items
          posts.value = Array.isArray(res.data?.list) ? res.data.list : [];
          // 确保total.value是数字
          total.value = typeof res.data?.total === 'number' ? res.data.total : 0;
        } else {
          console.error('API响应格式不符合预期:', res);
          posts.value = [];
          total.value = 0;
        }
      } catch (error) {
        console.error('获取帖子列表失败:', error);
        posts.value = [];
        total.value = 0;
      } finally {
        loading.value = false;
      }
    };

    // 处理分页大小变化
    const handleSizeChange = (size) => {
      pageSize.value = size;
      fetchPosts();
    };

    // 处理当前页变化
    const handleCurrentChange = (current) => {
      currentPage.value = current;
      fetchPosts();
    };

    const router = useRouter();

    // 添加帖子管理相关路由
    router.addRoute('adminDashboard', {
      path: 'posts/:id/edit',
      name: 'adminEditPost',
      component: () => import('../views/AdminPostEditView.vue'),
      meta: { requiresAuth: true, requiresAdmin: true }
    });

    // 处理编辑帖子
    const handleEditPost = (post) => {
      router.push(`/admin/dashboard/posts/${post.id}/edit`);
    };

    // 处理删除帖子
    const handleDeletePost = (id) => {
      alert(`删除帖子 ID: ${id} 功能待实现`);
    };

    onMounted(() => {
      fetchPosts();
    });

    return {
      posts,
      currentPage,
      pageSize,
      total,
      loading,
      handleSizeChange,
      handleCurrentChange,
      handleEditPost,
      handleDeletePost
    };
  }
};
</script>

<style scoped>
.admin-post-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>