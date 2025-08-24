<template>
  <div class="admin-user-create-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>新增用户</span>
        </div>
      </template>
      <el-form ref="userForm" :model="user" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="user.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="user.email" type="email" placeholder="请输入邮箱"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="user.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="user.role" placeholder="请选择角色">
            <el-option label="普通用户" value="user"></el-option>
            <el-option label="管理员" value="admin"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">创建</el-button>
          <el-button @click="handleCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { userAPI } from '../api/index';

export default {
  name: 'AdminUserCreateView',
  setup() {
    const router = useRouter();
    const user = reactive({
      username: '',
      email: '',
      password: '',
      role: 'user'
    });
    const loading = ref(false);

    // 提交表单
    const handleSubmit = async () => {
      try {
        loading.value = true;
        await userAPI.createUser(user);
        ElMessage.success('创建成功');
        router.push('/admin/dashboard/users');
      } catch (error) {
        console.error('创建失败:', error);
        ElMessage.error('创建失败');
      } finally {
        loading.value = false;
      }
    };

    // 取消创建
    const handleCancel = () => {
      router.push('/admin/dashboard/users');
    };

    return {
      user,
      loading,
      handleSubmit,
      handleCancel
    };
  }
};
</script>

<style scoped>
.admin-user-create-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>