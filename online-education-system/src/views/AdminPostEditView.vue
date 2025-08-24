<template>
  <div class="admin-post-edit-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>编辑帖子</span>
        </div>
      </template>
      <el-form ref="postForm" :model="post" label-width="100px">
        <el-form-item label="帖子标题" prop="title">
          <el-input v-model="post.title" placeholder="请输入帖子标题"></el-input>
        </el-form-item>
        <el-form-item label="帖子内容" prop="content">
          <el-input
            v-model="post.content"
            type="textarea"
            placeholder="请输入帖子内容"
            rows="6"
          ></el-input>
        </el-form-item>
        <el-form-item label="帖子状态">
          <el-select v-model="post.status" placeholder="请选择帖子状态">
            <el-option label="草稿" value="draft"></el-option>
            <el-option label="已发布" value="published"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">保存</el-button>
          <el-button @click="handleCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref, onMounted, reactive } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { postAPI } from '../api/index';

export default {
  name: 'AdminPostEditView',
  setup() {
    const router = useRouter();
    const route = useRoute();
    const postId = route.params.id;
    const post = reactive({
      id: postId,
      title: '',
      content: '',
      status: 'draft'
    });
    const loading = ref(false);

    // 获取帖子详情
    const fetchPostDetail = async () => {
      try {
        loading.value = true;
        const res = await postAPI.getPostDetail(postId);
        Object.assign(post, res);
      } catch (error) {
        console.error('获取帖子详情失败:', error);
      } finally {
        loading.value = false;
      }
    };

    // 提交表单
    const handleSubmit = async () => {
      try {
        loading.value = true;
        await postAPI.updatePost(post.id, post);
        ElMessage.success('保存成功');
        router.push('/admin/dashboard/posts');
      } catch (error) {
        console.error('保存失败:', error);
        ElMessage.error('保存失败');
      } finally {
        loading.value = false;
      }
    };

    // 取消编辑
    const handleCancel = () => {
      router.push('/admin/dashboard/posts');
    };

    onMounted(() => {
      fetchPostDetail();
    });

    return {
      post,
      loading,
      handleSubmit,
      handleCancel
    };
  }
};
</script>

<style scoped>
.admin-post-edit-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>