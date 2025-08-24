<template>
  <div class="admin-comment-list-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>评论管理</span>
        </div>
      </template>
      <el-table :data="comments" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="content" label="评论内容">
          <template #default="scope">
            <div class="comment-content" :title="scope.row.content">{{ scope.row.content }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="authorName" label="评论者" width="120"></el-table-column>
        <el-table-column prop="videoTitle" label="视频标题" width="180">
          <template #default="scope">
            <div class="video-title" :title="scope.row.videoTitle">{{ scope.row.videoTitle }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="评论时间" width="180"></el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="danger" size="small" @click="handleDeleteComment(scope.row.id)">删除</el-button>
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

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { commentAPI } from '@/api/index';

// 状态管理
const comments = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const loading = ref(false);

// 获取评论列表
const fetchComments = async () => {
  try {
    loading.value = true;
    const res = await commentAPI.getCommentList({
      page: currentPage.value,
      pageSize: pageSize.value
    });
    // 数据验证和处理
    console.log('评论列表API响应:', res);
    if (res && typeof res === 'object') {
      // 确保comments.value始终是数组
      comments.value = Array.isArray(res.data?.list) ? res.data.list : [];
      // 确保total.value是数字
      total.value = typeof res.data?.total === 'number' ? res.data.total : 0;
    } else {
      console.error('API响应格式不符合预期:', res);
      comments.value = [];
      total.value = 0;
    }
  } catch (error) {
    console.error('获取评论列表失败:', error);
    ElMessage.error('获取评论列表失败: ' + (error.message || '未知错误'));
  } finally {
    loading.value = false;
  }
};

// 处理分页大小变化
const handleSizeChange = (size) => {
  pageSize.value = size;
  fetchComments();
};

// 处理当前页变化
const handleCurrentChange = (current) => {
  currentPage.value = current;
  fetchComments();
};

// 处理删除评论
const handleDeleteComment = async (id) => {
  try {
    loading.value = true;
    await commentAPI.deleteComment(id);
    ElMessage.success('删除评论成功');
    // 重新获取评论列表
    fetchComments();
  } catch (error) {
    console.error('删除评论失败:', error);
    ElMessage.error('删除评论失败: ' + (error.message || '未知错误'));
  } finally {
    loading.value = false;
  }
};

// 初始加载
onMounted(() => {
  fetchComments();
});
</script>

<style scoped>
.admin-comment-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.comment-content,
.video-title {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>