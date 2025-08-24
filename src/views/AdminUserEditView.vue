<template>
  <div class="admin-user-edit-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>编辑用户</span>
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
          <el-input v-model="user.password" type="password" placeholder="请输入密码(可选)"></el-input>
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="user.role" placeholder="请选择角色">
            <el-option label="普通用户" value="user"></el-option>
            <el-option label="管理员" value="admin"></el-option>
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
import { userAPI } from '../api/index';

export default {
  name: 'AdminUserEditView',
  setup() {
    const router = useRouter();
    const route = useRoute();
    const userId = route.params.id;
    const user = reactive({
      id: userId,
      username: '',
      email: '',
      password: '',
      role: 'user'
    });
    const loading = ref(false);

    // 获取用户详情
    const fetchUserDetail = async () => {
      try {
        loading.value = true;
        const res = await userAPI.getUserDetail(userId);
        // 不设置密码字段，让用户选择性更新
        user.username = res.username;
        user.email = res.email;
        user.role = res.role;
      } catch (error) {
        console.error('获取用户详情失败:', error);
      } finally {
        loading.value = false;
      }
    };

    // 提交表单
    const handleSubmit = async () => {
      try {
        loading.value = true;
        // 只提交非空字段
        const updateData = {
          username: user.username,
          email: user.email,
          role: user.role
        };
        // 如果密码不为空，则包含密码
        if (user.password) {
          updateData.password = user.password;
        }
        await userAPI.updateUser(user.id, updateData);
        ElMessage.success('保存成功');
        router.push('/admin/dashboard/users');
      } catch (error) {
        console.error('保存失败:', error);
        ElMessage.error('保存失败');
      } finally {
        loading.value = false;
      }
    };

    // 取消编辑
    const handleCancel = () => {
      router.push('/admin/dashboard/users');
    };

    onMounted(() => {
      fetchUserDetail();
    });

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
.admin-user-edit-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>